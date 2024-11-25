package pokeapi

import (
	"net/http"
	"time"

	"github.com/yuetteyo/pokedex-cli-go/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2"

// Создаем своего клиента, так как нужен определенный timeout
type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
