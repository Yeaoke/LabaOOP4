from sqlalchemy import Column, Integer, String
from sqlalchemy.orm import relationship 
from models.industrial_companies import Base


class HoldingCompanies(Base):
    __tablename__ = 'holding_companies'

    id = Column(Integer, primary_key=True, autoincrement=True)
    holding_name = Column('holding_name', String(100), nullable=False, unique=True)

    industrial_companies = relationship(
        "IndustrialCompanies", 
        back_populates="holding_company",
        foreign_keys="IndustrialCompanies.holding_companies_id"
    )

    def __init__(self, holding_name: str):
        self.holding_name = holding_name

    def add_company(self, company):
        self.industrial_companies.append(company)

    def __repr__(self):
        return f"Holding(id={self.id}, name='{self.holding_name}')"