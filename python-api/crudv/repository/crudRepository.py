from typing import Any, Generic, List, Optional, Type, TypeVar
from sqlalchemy.orm import Session
import uuid

T = TypeVar("T")


class crudRepository(Generic[T]):
    def __init__(self, model: Type[T]):
        self.model = model

    def get(self, db: Session, id: Any) -> Optional[T]:
        # Принимаем и строку, и UUID-объект
        if isinstance(id, str):
            try:
                id = uuid.UUID(id)
            except ValueError:
                return None
        return db.query(self.model).filter(self.model.id == id).first()

    def create(self, db: Session, *, obj_in_data: dict) -> T:
        from models.coal_company import CoalCompany
        from models.oil_company import OilCompany
        from models.holding_companies import HoldingCompanies

        company_type = obj_in_data.get("company_type", "")

        # Нормализация типа: принимаем оба формата
        if company_type in ("coal", "CoalCompany"):
            model_cls = CoalCompany
        elif company_type in ("oil", "OilCompany"):
            model_cls = OilCompany
        else:
            raise ValueError(f"Unknown company_type: {company_type!r}")

        # Обработка холдинга по имени
        holding_id = None
        holding_name = obj_in_data.get("holding_name")
        if holding_name and holding_name.strip():
            holding = db.query(HoldingCompanies).filter_by(holding_name=holding_name.strip()).first()
            if not holding:
                holding = HoldingCompanies(holding_name=holding_name.strip())
                db.add(holding)
                db.flush()
            holding_id = holding.id

        # Создание объекта нужного класса
        if model_cls == CoalCompany:
            obj = CoalCompany(
                company_name=obj_in_data.get("company_name"),
                annual_turnover=obj_in_data.get("annual_turnover"),
                coal_volume=obj_in_data.get("coal_volume"),
                mine_count=obj_in_data.get("mineCount"),
                holding_companies_id=holding_id,
            )
            if obj_in_data.get("coal_action"):
                obj.coal_action = obj_in_data["coal_action"]
        else:
            obj = OilCompany(
                company_name=obj_in_data.get("company_name"),
                annual_turnover=obj_in_data.get("annual_turnover"),
                oil_volume=obj_in_data.get("oil_volume"),
                hole_count=obj_in_data.get("holeCount"),
                holding_companies_id=holding_id,
            )
            if obj_in_data.get("oil_action"):
                obj.oil_action = obj_in_data["oil_action"]

        # Если Go передал UUID — используем его, иначе генерируем новый
        raw_id = obj_in_data.get("id")
        if raw_id:
            try:
                parsed_id = uuid.UUID(str(raw_id))
                # Проверяем, что это не нулевой UUID
                if parsed_id != uuid.UUID(int=0):
                    obj.id = parsed_id
            except (ValueError, AttributeError):
                pass

        db.add(obj)
        db.commit()
        db.refresh(obj)
        return obj

    def update(self, db: Session, *, db_obj: T, obj_in_data: dict) -> T:
        from models.holding_companies import HoldingCompanies

        holding_name = obj_in_data.get("holding_name")
        if holding_name and holding_name.strip():
            holding = db.query(HoldingCompanies).filter_by(holding_name=holding_name.strip()).first()
            if not holding:
                holding = HoldingCompanies(holding_name=holding_name.strip())
                db.add(holding)
                db.flush()
            obj_in_data = {**obj_in_data, "holding_companies_id": holding.id}

        key_mapping = {
            "mineCount": "mine_count",
            "holeCount": "hole_count",
        }

        skip_fields = {"id", "company_type", "holding_name"}

        for field, value in obj_in_data.items():
            if field in skip_fields:
                continue
            attr = key_mapping.get(field, field)
            if hasattr(db_obj, attr):
                setattr(db_obj, attr, value)

        db.add(db_obj)
        db.commit()
        db.refresh(db_obj)
        return db_obj

    def delete(self, db: Session, *, id: Any) -> bool:
        if isinstance(id, str):
            try:
                id = uuid.UUID(id)
            except ValueError:
                return False
        rows_deleted = db.query(self.model).filter(self.model.id == id).delete()
        db.commit()
        return rows_deleted > 0

    def get_all(self, db: Session, skip: int = 0, limit: int = 100) -> List[T]:
        return db.query(self.model).offset(skip).limit(limit).all()