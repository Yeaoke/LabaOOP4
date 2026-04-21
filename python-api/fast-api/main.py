from fastapi import Depends, FastAPI
from crudv import CRUD
from sqlalchemy import Session

from database import get_db
import os

app = FastAPI()

backendURL = os.getenv("backendPython", "http://localhost:8081")

@app.get("/")
async def home(db: Session = Depends(get_db)):
    return CRUD.get_all(db) 

@app.post("/add")
async def create(db: Session = Depends(get_db)):
    return CRUD.create(db)

@app.post("/edit/{company_id}")
async def edit(company_id: int, db: Session = Depends(get_db)):
    CRUD.update(db)
    return {"message": "Company was updated successfully", "id": company_id}

@app.get("/info")
async def info(company_id: int, db: Session = Depends(get_db)):
    CRUD.get(db)
    return {"message": "Info was got successfully", "id": company_id}

@app.delete("/delete/{company_id}")
async def delete(company_id: int, db: Session = Depends(get_db)):
    CRUD.delete(db)
    return {"message": "Company was deleted successfully", "id": company_id}
