from dataclasses import dataclass, field
from sqlalchemy import Column, Integer, String, Text, Float
from models.coal_company import CoalCompany
from models.oil_company import OilCompany
from database import Base

@dataclass
class IndustrialCompanies(Base):
    __tablename__ = "industrial_companies"

    id = Column(Integer, primary_key=True, index=True)
    name = Column(String, nullable=False)
    annualTurnover = Column(Integer)
    company_type = Column(String)
    volume = Column(Integer)
    workingObjects = Column(Integer),

    def __init__(self, name: str = "", annualTurnover: int = 0, company_type: str = "", volume: int = 0, workingObjects: int = 0):
        

        self.name = name,
        self.annualTurnover = annualTurnover,
        self.company_type = company_type,
        self.volume = volume,
        self.workingObjects = workingObjects

    