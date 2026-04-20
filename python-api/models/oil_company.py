from dataclasses import dataclass, field
from sqlalchemy import Column, Integer, String, ForeignKey
from models.industrial_companies import IndustrialCompanies

@dataclass
class OilCompany(IndustrialCompanies):
    __tablename__ = "oil_company"

    id = Column(Integer, ForeignKey("industrial_companies.id"), primary_key=True)
    name = Column(String, nullable=False)
    annualTurnover = Column(Integer)
    company_type = Column(String)
    volume = Column(Integer)
    workingObjects = Column(Integer)