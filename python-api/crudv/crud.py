from repository.crudRepository import crudRepository
from models.industrial_companies import IndustrialCompanies
from typing import List
from sqlalchemy import Session, select
from repository.crudRepository import T

class IndustrialCompaniesRepository(crudRepository[IndustrialCompanies]):
    def __init__(self, model):
        super().__init__(model)


    def get_all(self, db: Session, skip: int = 0, limit: int = 100) -> List[T]:
        return db.query(self.model).offset(skip).limit(limit).all()  

CRUD = IndustrialCompaniesRepository()