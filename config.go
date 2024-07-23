package certificate_center_sdk

type Config struct {
	BaseURL string
	ApiKey  string
}

func NewConfig(baseURL string, apiKey string) *Config {
	return &Config{
		BaseURL: baseURL,
		ApiKey:  apiKey,
	}
}
