package cache

import (
	"LabaOOP4/go-server/models"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/hashicorp/golang-lru/v2/expirable"
)

var cache = expirable.NewLRU[uuid.UUID, models.IndustrialCompanies](100, nil, time.Minute*10)

func Caching(id uuid.UUID, model models.IndustrialCompanies) {
	cache.Add(id, model)

	if company, ok := cache.Get(id); ok {
		log.Println(company)
	}
}
