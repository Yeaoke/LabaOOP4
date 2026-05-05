from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker

# Импортируем ВСЕ модели для регистрации в SQLAlchemy
from models.holding_companies import HoldingCompanies
from models.industrial_companies import Base, IndustrialCompanies
from models.coal_company import CoalCompany
from models.oil_company import OilCompany

DATABASE_URL = "postgresql://postgres:MaL8033038@localhost:5432/OOP2FFF"
engine = create_engine(DATABASE_URL)

SessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine)

Base.metadata.create_all(engine)

async def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()