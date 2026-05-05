from typing import Any, Generic, List, Optional, Type, TypeVar
from sqlalchemy.orm import Session

T = TypeVar("T")

class crudRepository(Generic[T]):
    def __init__(self, model: Type[T]):
        self.model = model
    
    def get(self, db: Session, id: Any) -> Optional[T]:
        return db.query(self.model).filter(self.model.id == id).first()

    def create(self, db: Session, *, obj_in_data: dict) -> T:
        obj_in_data.pop("id", None)
        db_obj = self.model(**obj_in_data)
        db.add(db_obj)
        db.commit()
        db.refresh(db_obj)
        return db_obj

    def update(self, db: Session, *, db_obj: T, obj_in_data: dict) -> T:
        for field, value in obj_in_data.items():
            if hasattr(db_obj, field) and field != "id":
                setattr(db_obj, field, value)
        db.add(db_obj)
        db.commit()
        db.refresh(db_obj)
        return db_obj    

    def delete(self, db: Session, *, id: int) -> bool:
        rows_deleted = db.query(self.model).filter(self.model.id == id).delete()
        db.commit()
        return rows_deleted > 0