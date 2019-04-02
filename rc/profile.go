package rc

import (
	"fmt"
)

type PivnetProfile struct {
	Name              string `yaml:"name"`
	APIToken          string `yaml:"api_token"`
	Host              string `yaml:"host"`
	AccessToken       string `yaml:"access_token"`
	AccessTokenExpiry int64  `yaml:"access_token_expiry"`
}

func (p *PivnetProfile) Validate() error {
	if p.Name == "" {
		return fmt.Errorf("Name is empty")
	}

	if p.APIToken == "" {
		return fmt.Errorf("API token is empty")
	}

	if p.Host == "" {
		return fmt.Errorf("Host is empty")
	}

	return nil
}
