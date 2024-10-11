package lago

import (
	"net/http"
	"testing"

	"github.com/google/uuid"
)

func TestConfig_Validate(t *testing.T) {
	tt := map[string]struct {
		c       *Config
		wantErr bool
	}{
		"valid": {
			c: &Config{
				BaseURL: "https://example.com",
				APIKey:  uuid.NewString(),
				Client:  &http.Client{},
			},
			wantErr: false,
		},
		"invalid BaseURL with path": {
			c: &Config{
				BaseURL: "https://example.com/api/v1",
				APIKey:  uuid.NewString(),
				Client:  &http.Client{},
			},
			wantErr: true,
		},
		"invalid BaseURL with query": {
			c: &Config{
				BaseURL: "https://example.com?query=1",
				APIKey:  uuid.NewString(),
				Client:  &http.Client{},
			},
			wantErr: true,
		},
		"invalid BaseURL with invalid scheme": {
			c: &Config{
				BaseURL: "ftp://example.com",
				APIKey:  uuid.NewString(),
				Client:  &http.Client{},
			},
			wantErr: true,
		},
		"empty APIKey": {
			c: &Config{
				BaseURL: "https://example.com",
				APIKey:  "",
				Client:  &http.Client{},
			},
			wantErr: true,
		},
		"nil Client": {
			c: &Config{
				BaseURL: "https://example.com",
				APIKey:  uuid.NewString(),
				Client:  nil,
			},
			wantErr: true,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			gotErr := tc.c.Validate()
			if (gotErr != nil) != tc.wantErr {
				t.Errorf("Validate() = %v, wantErr %v", gotErr, tc.wantErr)
			}
		})
	}
}
