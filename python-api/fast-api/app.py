from fastapi import Depends, FastAPI
from crud import crud
from sqlalchemy import Session
from database import get_db

app = FastAPI()

@app.get("/")
async def home(db: Session = Depends(get_db)):
    return crud.get_all(db) 

@app.post("/add")
async def create(db: Session = Depends(get_db)):
    return crud.create(db)

@app.post("/edit/{company_id}")
async def edit(company_id: int, db: Session = Depends(get_db)):
    crud.update(db)
    return {"message": "Company was updated successfully", "id": company_id}

@app.get("/info")
async def info(company_id: int, db: Session = Depends(get_db)):
    crud.get(db)
    return {"message": "Info was got successfully", "id": company_id}

@app.delete("/delete/{company_id}")
async def delete(company_id: int, db: Session = Depends(get_db)):
    crud.delete(db)
    return {"message": "Company was deleted successfully", "id": company_id}