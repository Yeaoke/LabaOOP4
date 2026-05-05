from sqlalchemy import Column, Integer, BigInteger, String, ForeignKey
from sqlalchemy.orm import relationship
from models.industrial_companies import IndustrialCompanies, Base


class CoalCompany(IndustrialCompanies):
    __tablename__ = 'coal_company'

    id = Column(
        Integer, 
        ForeignKey('industrial_companies.id', ondelete='CASCADE'), 
        primary_key=True
    )
    
    coal_volume = Column(BigInteger, nullable=True)
    mine_count = Column(BigInteger, nullable=True) 
    coal_action = Column(String(255), nullable=True)

    __mapper_args__ = {
        'polymorphic_identity': 'CoalCompany',
    }

    def __init__(self, company_name: str, annual_turnover: int = None, 
                 coal_volume: int = None, mine_count: int = None, 
                 holding_companies_id: int = None):
        super().__init__(
            company_name, 
            annual_turnover, 
            company_type='CoalCompany',
            holding_companies_id=holding_companies_id
        )
        self.coal_volume = coal_volume
        self.mine_count = mine_count
        self.coal_action = None

    def show_info(self):
        print("=== Угольная компания ===")
        print(f"ID: {self.id}")
        print(f"Название: {self.company_name}")
        print(f"Оборот: {self.annual_turnover}")
        print(f"Объём угля: {self.coal_volume}")
        print(f"Шахт: {self.mine_count}")
        if self.coal_action:
            print(f"Действие: {self.coal_action}")

    def cost_of_time_production(self):
        if self.coal_volume and self.coal_volume > 0 and self.annual_turnover:
            counter = self.annual_turnover // self.coal_volume
            self.coal_action = f"Стоимость за единицу: {counter}"
            return counter
        return None

    def add_coal_mines(self, count: int):
        if count and count > 0:
            current = self.mine_count or 0
            self.mine_count = current + count
            self.coal_action = f"Добавлено шахт: {count}"

    def stop_expanding(self):
        self.coal_action = "Производство остановлено"

    def to_dict(self):
        data = super().to_dict()
        data.update({
            "coal_volume": self.coal_volume,
            "mineCount": self.mine_count,
            "coal_action": self.coal_action,
        })
        return data