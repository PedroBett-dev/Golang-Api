package config

func (c *Config) GetApiUrl() string {
	return c.ApiConfig.url
}

func (c *Config) GetApiPort() string {
	return c.ApiConfig.port
}
