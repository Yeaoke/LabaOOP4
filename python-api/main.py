from fastapi import FastAPI, Request, Depends, HTTPException, Query
from fastapi.middleware.cors import CORSMiddleware
from sqlalchemy.orm import Session
from sqlalchemy import select
from database import get_db, IndustrialCompanies, HoldingCompanies
from crudv.crud import CRUD
from typing import Optional
import uuid

app = FastAPI(title="Industrial Companies API")

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_methods=["*"],
    allow_headers=["*"],
)

def company_to_dict(company) -> dict:
    """Безопасное преобразование компании в dict"""
    data = company.to_dict() if hasattr(company, 'to_dict') else {
        'id': str(company.id),
        'company_name': company.company_name,
        'company_type': company.company_type,
        'annual_turnover': company.annual_turnover,
    }
    if hasattr(company, 'coal_action') and company.coal_action:
        data['action_result'] = company.coal_action
    elif hasattr(company, 'oil_action') and company.oil_action:
        data['action_result'] = company.oil_action
    return data

@app.get("/")
async def get_all_companies(db: Session = Depends(get_db)):
    companies = CRUD.get_all(db)
    return [company_to_dict(c) for c in companies]


@app.get("/holding")
async def get_companies_by_holding(
    holding_name: str = Query(..., description="Название холдинга"),
    db: Session = Depends(get_db)
):
    """Вернуть список компаний по названию холдинга"""
    stmt = select(IndustrialCompanies).join(
        HoldingCompanies, 
        IndustrialCompanies.holding_companies_id == HoldingCompanies.id
    ).where(
        HoldingCompanies.holding_name == holding_name
    )
    
    companies = db.execute(stmt).scalars().all()
    return [company_to_dict(c) for c in companies]


@app.post("/add")
async def create(request: Request, db: Session = Depends(get_db)):
    data = await request.json()
    
    raw_id = data.get("id")
    if not raw_id or str(raw_id) in ("00000000-0000-0000-0000-000000000000", "null", ""):
        data.pop("id", None)
    else:
        try:
            data["id"] = str(uuid.UUID(str(raw_id)))
        except ValueError:
            raise HTTPException(status_code=400, detail="Invalid UUID format")
    
    try:
        result = CRUD.create(db, obj_in_data=data)
    except ValueError as e:
        raise HTTPException(status_code=400, detail=str(e))
    except Exception as e:
        raise HTTPException(status_code=409, detail=f"Database error: {str(e)}")
    
    return company_to_dict(result)


@app.post("/edit/{company_id}")
async def edit(company_id: str, request: Request, db: Session = Depends(get_db)):
    data = await request.json()
    db_obj = CRUD.get(db, id=company_id)
    if not db_obj:
        raise HTTPException(status_code=404, detail="Company not found")
    
    result = CRUD.update(db, db_obj=db_obj, obj_in_data=data)
    return company_to_dict(result)


@app.get("/info/{company_id}")
async def info(company_id: str, db: Session = Depends(get_db)):
    result = CRUD.get(db, id=company_id)
    if not result:
        raise HTTPException(status_code=404, detail="Company not found")
    return company_to_dict(result)


@app.delete("/delete/{company_id}")
async def delete(company_id: str, db: Session = Depends(get_db)):
    success = CRUD.delete(db, id=company_id)
    if not success:
        raise HTTPException(status_code=404, detail="Company not found")
    return {"message": "Deleted", "id": company_id}

@app.post("/action/{company_id}/calculate")
async def calculate(company_id: str, db: Session = Depends(get_db)):
    company = CRUD.get(db, id=company_id)
    if not company:
        raise HTTPException(status_code=404, detail="Company not found")
    
    result = company.cost_of_time_production()
    db.add(company)
    db.commit()
    db.refresh(company)
    
    return {**company_to_dict(company), "action_result": result or ""}


@app.post("/action/{company_id}/add-mines")
async def add_mines(company_id: str, request: Request, db: Session = Depends(get_db)):
    from models.coal_company import CoalCompany
    
    data = await request.json()
    count = data.get("count", 0)
    
    company = CRUD.get(db, id=company_id)
    if not company:
        raise HTTPException(status_code=404, detail="Company not found")
    if not isinstance(company, CoalCompany):
        raise HTTPException(status_code=400, detail="Only for coal companies")
    
    company.add_coal_mines(int(count))
    db.add(company)
    db.commit()
    db.refresh(company)
    
    return {**company_to_dict(company), "action_result": company.coal_action or ""}


@app.post("/action/{company_id}/stop-expanding")
async def stop_expanding(company_id: str, db: Session = Depends(get_db)):
    from models.coal_company import CoalCompany
    
    company = CRUD.get(db, id=company_id)
    if not company:
        raise HTTPException(status_code=404, detail="Company not found")
    if not isinstance(company, CoalCompany):
        raise HTTPException(status_code=400, detail="Only for coal companies")
    
    company.stop_expanding()
    db.add(company)
    db.commit()
    db.refresh(company)
    
    return {**company_to_dict(company), "action_result": company.coal_action or ""}


@app.post("/action/{company_id}/add-wells")
async def add_wells(company_id: str, request: Request, db: Session = Depends(get_db)):
    from models.oil_company import OilCompany
    
    data = await request.json()
    count = data.get("count", 0)
    
    company = CRUD.get(db, id=company_id)
    if not company:
        raise HTTPException(status_code=404, detail="Company not found")
    if not isinstance(company, OilCompany):
        raise HTTPException(status_code=400, detail="Only for oil companies")
    
    company.add_oil_wells(int(count))
    db.add(company)
    db.commit()
    db.refresh(company)
    
    return {**company_to_dict(company), "action_result": company.oil_action or ""}


@app.post("/action/{company_id}/check-resources")
async def check_resources(company_id: str, db: Session = Depends(get_db)):
    from models.oil_company import OilCompany
    
    company = CRUD.get(db, id=company_id)
    if not company:
        raise HTTPException(status_code=404, detail="Company not found")
    if not isinstance(company, OilCompany):
        raise HTTPException(status_code=400, detail="Only for oil companies")
    
    company.check_resources()
    db.add(company)
    db.commit()
    db.refresh(company)
    
    return {**company_to_dict(company), "action_result": company.oil_action or ""}