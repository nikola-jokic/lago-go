package lago

import (
	"errors"
	"fmt"
	"net/url"
)

// Config is a struct that holds the configuration for the client.
type Config struct {
	BaseURL          string
	BaseIngestURL    string
	APIKey           string
	UseIngestService bool
	Debug            bool
	Client           HTTPClient
}

func (c *Config) Validate() error {
	if err := validateBaseURL(c.BaseURL); err != nil {
		return fmt.Errorf("BaseURL validation error: %v", err)
	}
	if err := validateBaseURL(c.BaseIngestURL); err != nil {
		return fmt.Errorf("BaseIngestURL validation error: %v", err)
	}
	if c.Client == nil {
		return errors.New("Client is nil")
	}
	if c.APIKey == "" {
		return errors.New("APIKey is empty")
	}
	return nil
}

func validateBaseURL(s string) error {
	u, err := url.Parse(s)
	if err != nil {
		return err
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return errors.New("invalid scheme")
	}
	if u.Host == "" {
		return errors.New("missing host")
	}
	if u.Path != "" {
		return errors.New("path must be empty")
	}
	if len(u.Query()) > 0 {
		return errors.New("query must be empty")
	}
	return nil
}
