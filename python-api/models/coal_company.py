from sqlalchemy import Column, String, BigInteger, Integer, ForeignKey
from sqlalchemy.orm import relationship
from industrial_companies import IndustrialCompanies, Base

class CoalCompany(IndustrialCompanies):
    __tablename__ = 'coal_company'

    # Первичный ключ также является внешним ключом к родительской таблице
    id = Column(Integer, ForeignKey('industrial_companies.id'), primary_key=True)
    
    coal_volume = Column(BigInteger, nullable=True)
    mine_count = Column(BigInteger, nullable=True)
    coal_action = Column(String, nullable=True)

    __mapper_args__ = {
        'polymorphic_identity': 'CoalCompany',
    }

    def __init__(self, company_name: str, annual_turnover: int = None, 
        coal_volume: int = None, mine_count: int = None):
        super().__init__(company_name, annual_turnover)
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
        if (self.coal_volume is not None and self.coal_volume > 0 
                and self.annual_turnover is not None):
            counter = self.annual_turnover // self.coal_volume
            print(f"Стоимость за единицу: {counter}")
            self.coal_action = f"Стоимость за единицу: {counter}"

    def add_coal_mines(self, count: int):
        if count is not None and count > 0:
            current_mines = self.mine_count if self.mine_count is not None else 0
            self.mine_count = current_mines + count
            print(f"Добавлено шахт: {count}")
            self.coal_action = f"Добавлено шахт: {count}"

    def stop_expanding(self):
        self.coal_action = "STOPPED"
        print("Производство остановлено")
        # Дублируем логику setCoalAction из Java
        self.coal_action = "Производство остановлено"

    def __repr__(self):
        return super().__repr__()