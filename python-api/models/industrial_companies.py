from sqlalchemy import Column, Integer, String, BigInteger, ForeignKey
from sqlalchemy.orm import declarative_base, relationship
from abc import ABC, abstractmethod

Base = declarative_base()

class IndustrialCompanies(Base):
    __tablename__ = 'industrial_companies'

    id = Column(Integer, primary_key=True)
    company_name = Column('name', String(50), nullable=False)
    annual_turnover = Column('turnover', BigInteger, nullable=True)
    type = Column('type', String(50), nullable=False)
    
    holding_companies_id = Column(Integer, ForeignKey('holding_companies.id'), nullable=True)
    holding_company = relationship("HoldingCompanies", back_populates="industrial_companies")

    __mapper_args__ = {
        'polymorphic_identity': 'IndustrialCompanies',
        'polymorphic_on': type
    }

    def __init__(self, company_name: str, annual_turnover: int = None):
        self.company_name = company_name
        self.annual_turnover = annual_turnover

    @abstractmethod
    def cost_of_time_production(self):
        pass

    @abstractmethod
    def show_info(self):
        pass

    def __repr__(self):
        return f"{self.id} {self.company_name} {self.type}"