package lago

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"net/http"

	jwt "github.com/golang-jwt/jwt/v5"
)

func (c *Client) GetWebhookPublicKey(ctx context.Context) (*rsa.PublicKey, *Error) {
	u := c.url("webhooks/public_key", nil)
	result, err := get[string](ctx, c, u)
	if err != nil {
		return nil, err
	}
	// Decode the base64-encoded key
	bytesResult, decodeErr := base64.StdEncoding.DecodeString(*result)
	if decodeErr != nil {
		return nil, &Error{
			Err:            decodeErr,
			HTTPStatusCode: http.StatusInternalServerError,
			Message:        "cannot decode the key",
		}
	}

	// Parse the PEM block
	block, _ := pem.Decode(bytesResult)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, &Error{
			Err:            errors.New("Failed to decode PEM block containing public key"),
			HTTPStatusCode: http.StatusInternalServerError,
			Message:        "Failed to decode PEM block containing public key",
		}
	}

	// Parse the DER-encoded public key
	publicKey, parseErr := x509.ParsePKIXPublicKey(block.Bytes)
	if parseErr != nil {
		return nil, &Error{
			Err:            parseErr,
			HTTPStatusCode: http.StatusInternalServerError,
			Message:        "Failed to to parse the public key",
		}
	}

	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return nil, &Error{
			Err:            errors.New("Unexpected type of public key"),
			HTTPStatusCode: http.StatusInternalServerError,
			Message:        "Unexpected type of public key",
		}
	}

	return rsaPublicKey, nil
}

func (c *Client) parseSignature(ctx context.Context, signature string) (*jwt.Token, *Error) {
	publicKey, err := c.GetWebhookPublicKey(ctx)
	if err != nil {
		return nil, err
	}

	token, parseErr := jwt.Parse(signature, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})
	if parseErr != nil {
		return nil, &Error{
			Err:            parseErr,
			HTTPStatusCode: http.StatusInternalServerError,
			Message:        "cannot parse token",
		}
	}

	return token, nil
}

func (c *Client) ValidateSignature(ctx context.Context, signature string) (bool, *Error) {
	if token, err := c.parseSignature(ctx, signature); err == nil && token.Valid {
		return true, nil
	} else {
		return false, err
	}
}

func (c *Client) ValidateBody(ctx context.Context, signature string, body string) (bool, *Error) {
	if token, err := c.parseSignature(ctx, signature); err == nil && token.Valid {
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return false, &Error{
				Err:            errors.New("error casting claims"),
				HTTPStatusCode: http.StatusInternalServerError,
				Message:        "cannot parse token",
			}
		}

		return claims["data"] == body, nil
	} else {
		return false, err
	}
}
