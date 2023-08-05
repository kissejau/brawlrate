package brawlhalla

type BrawlhallaService interface{}

type brawlhallaService struct {
	url    string
	apiKey string
}

func NewBrawlhallaService(apiKey string) BrawlhallaService {
	brawlhallaService := &brawlhallaService{
		url:    "https://api.brawlhalla.com",
		apiKey: apiKey,
	}
	return brawlhallaService
}
