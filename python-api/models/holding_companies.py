from sqlalchemy import Column, Integer, String
from sqlalchemy.orm import relationship, declarative_base
from industrial_companies import Base, IndustrialCompanies

class HoldingCompanies(Base):
    __tablename__ = 'holding_companies'

    id = Column(Integer, primary_key=True, autoincrement=True)
    holding_name = Column('holding_name', String(100), nullable=False, unique=True)

    industrial_companies = relationship(
        "IndustrialCompanies", 
        back_populates="holding_company", 
        cascade="all, delete-orphan",
        lazy="joined" 
    )

    def __init__(self, holding_name: str):
        self.holding_name = holding_name
        self.industrial_companies = []

    def add_company(self, company: IndustrialCompanies):
        self.industrial_companies.append(company)
        company.holding_company = self

    def __repr__(self):
        return f"Holding(id={self.id}, name='{self.holding_name}', companies={len(self.industrial_companies)})"