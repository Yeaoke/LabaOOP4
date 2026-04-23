from crudv.repository.crudRepository import crudRepository
from models.industrial_companies import IndustrialCompanies
from typing import List
from sqlalchemy import select
from sqlalchemy.orm import Session
from crudv.repository.crudRepository import T

class IndustrialCompaniesRepository(crudRepository[IndustrialCompanies]):
    def __init__(self):
        super().__init__(IndustrialCompanies)


    def get_all(self, db: Session, skip: int = 0, limit: int = 100) -> List[T]:
        return db.query(self.model).offset(skip).limit(limit).all()  

CRUD = IndustrialCompaniesRepository()