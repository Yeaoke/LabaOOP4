# main.py - исправленные эндпоинты
from fastapi import FastAPI, Request, Depends
from fastapi.responses import JSONResponse
from sqlalchemy.orm import Session
from database import get_db
from crudv.crud import CRUD
import json

app = FastAPI(root_path="/api")

@app.get("/")
async def home(db: Session = Depends(get_db)):
    companies = CRUD.get_all(db)
    return [c.__dict__ for c in companies]  

@app.post("/add")
async def create(request: Request, db: Session = Depends(get_db)):
    data = await request.json()
    if data.get("id") == "00000000-0000-0000-0000-000000000000":
        data.pop("id", None)
    result = CRUD.create(db, obj_in_data=data)
    return result.__dict__

@app.post("/edit/{company_id}")
async def edit(company_id: int, request: Request, db: Session = Depends(get_db)):
    data = await request.json()
    db_obj = CRUD.get(db, id=company_id)
    if not db_obj:
        return JSONResponse(status_code=404, content={"error": "Not found"})
    result = CRUD.update(db, db_obj=db_obj, obj_in_data=data)
    return result.__dict__

@app.get("/info/{company_id}")
async def info(company_id: int, db: Session = Depends(get_db)):
    result = CRUD.get(db, id=company_id)
    if not result:
        return JSONResponse(status_code=404, content={"error": "Not found"})
    return result.__dict__

@app.delete("/delete/{company_id}")
async def delete(company_id: int, db: Session = Depends(get_db)):
    CRUD.delete(db, id=company_id)
    return {"message": "Deleted", "id": company_id}