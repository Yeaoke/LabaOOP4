from typing import Any, Generic, List, Optional, Type, TypeVar
from sqlalchemy.orm import Session
from models.industrial_companies import IndustrialCompanies

T = TypeVar("T", bound=True)

class crudRepository(Generic[T]):
    def __init__(self, model: Type[T]):
        self.model = model
   
    def get(self, db: Session, id: Any) -> Optional[T]:
        return db.query(self.model).filter(self.model.id == id).first()

    def create(self, db: Session, *, obj_in_data: dict) -> T:
        db_obj = self.model(**obj_in_data)
        db.add(db_obj)
        db.commit()
        db.refresh(db_obj)
        return db_obj

    def update(self, db: Session, *, db_obj: T, obj_in_data: dict) -> T:
        for field in obj_in_data:
            setattr(db_obj, field, obj_in_data[field])
        db.add(db_obj)
        db.commit()
        db.refresh(db_obj)
        return db_obj    

    def delete(self, db: Session, *, id: int) -> Optional[T]:
        db.query(self.model).filter(IndustrialCompanies.id == id).delete()
        db.commit()
    