package cache

import (
	dto "LabaOOP4/go-server/dto"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/hashicorp/golang-lru/v2/expirable"
)

var cache = expirable.NewLRU[uuid.UUID, dto.IndustrialCompanies](100, nil, time.Minute*10)

func CacheAdd(id uuid.UUID, model dto.IndustrialCompanies) {
	cache.Add(id, model)
	log.Printf("Cache ADD: %s - %s", id, model.CompanyName)
}

func CacheRemove(id uuid.UUID) {
	cache.Remove(id)
	log.Printf("Cache REMOVE: %s", id)
}

func CacheGet(id uuid.UUID) (dto.IndustrialCompanies, bool) {
	return cache.Get(id)
}

func CacheClean() {
	cache.Purge()
	log.Println("Cache CLEANED")
}
