package brawlhalla

type brawlhallaService struct {
	url    string
	apiKey string
}

func NewBrawlhallaService(apiKey string) *brawlhallaService {
	brawlhallaService := &brawlhallaService{
		url:    "https://api.brawlhalla.com",
		apiKey: apiKey,
	}
	return brawlhallaService
}
