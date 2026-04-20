from fastapi import Depends, FastAPI
from fastapi.responses import RedirectResponse
from crud import crud
from sqlalchemy import Session
from database import get_db
from models.industrial_companies import IndustrialCompanies_id

app = FastAPI()

@app.get("/")
async def home(db: Session = Depends(get_db)):
    crud.get_all(db)
    return 

@app.get("/add")
async def add():
    return crud.create(), RedirectResponse(url="/", status_code=303)

@app.get("/edit/{company_id}")
async def edit(company_id: int, db: Session = Depends(get_db)):
    crud.update(db)
    return RedirectResponse(url="/", status_code=303), {"message": "Company was updated successfully", "id": company_id}

@app.get("/info")
async def info(company_id: int, db: Session = Depends(get_db)):
    crud.get(db)
    return RedirectResponse(url="/", status_code=303), {"message": "Info was got successfully", "id": company_id}

@app.get("/delete/{company_id}")
async def delete(company_id: int, db: Session = Depends(get_db)):
    crud.delete(db)
    return RedirectResponse(url="/", status_code=303), {"message": "Company was deleted successfully", "id": company_id}