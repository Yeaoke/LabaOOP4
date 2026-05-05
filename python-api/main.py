from fastapi import FastAPI, Request, Depends, HTTPException
from fastapi.responses import JSONResponse
from sqlalchemy.orm import Session
from database import get_db
from crudv.crud import CRUD

app = FastAPI()

# 🔥 Base.metadata.create_all() уже вызван в database.py — здесь не нужен!

@app.get("/")
async def home(db: Session = Depends(get_db)):
    companies = CRUD.get_all(db)
    return [c.to_dict() for c in companies]

@app.post("/add")
async def create(request: Request, db: Session = Depends(get_db)):
    data = await request.json()
    if not data.get("id") or data["id"] == "00000000-0000-0000-0000-000000000000":
        data.pop("id", None)
    result = CRUD.create(db, obj_in_data=data)
    return result.to_dict()

@app.post("/edit/{company_id}")
async def edit(company_id: int, request: Request, db: Session = Depends(get_db)):
    data = await request.json()
    db_obj = CRUD.get(db, id=company_id)
    if not db_obj:
        raise HTTPException(status_code=404, detail="Not found")
    result = CRUD.update(db, db_obj=db_obj, obj_in_data=data)
    return result.to_dict()

@app.get("/info/{company_id}")
async def info(company_id: int, db: Session = Depends(get_db)):
    result = CRUD.get(db, id=company_id)
    if not result:
        raise HTTPException(status_code=404, detail="Not found")
    return result.to_dict()

@app.delete("/delete/{company_id}")
async def delete(company_id: int, db: Session = Depends(get_db)):
    success = CRUD.delete(db, id=company_id)
    if not success:
        raise HTTPException(status_code=404, detail="Not found")
    return {"message": "Deleted", "id": company_id}