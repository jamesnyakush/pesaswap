package api

type Env string


const (
	SANDBOX = iota
	PRODUCTION
)

type Service struct {
	consumerKey string
	appKey      string
	Env         int
}

func New(consumerKey, appKey string, env int) (Service, error) {
	return Service{consumerKey, appKey, env}, nil
}

func (s Service) baseURL() string {
	if s.Env == PRODUCTION {
		return "https://www.pesaswap.com/"
	}
	return "https://devpesaswap.azurewebsites.net/"
}
