package lago

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"

	jwt "github.com/golang-jwt/jwt/v5"
)

func (c *Client) GetWebhookPublicKey(ctx context.Context) (*rsa.PublicKey, error) {
	u := c.url("webhooks/public_key", nil)
	result, err := get[string](ctx, c, u)
	if err != nil {
		return nil, err
	}
	// Decode the base64-encoded key
	bytesResult, err := base64.StdEncoding.DecodeString(*result)
	if err != nil {
		return nil, fmt.Errorf("failed to base64 decode the public key: %v", err)
	}

	// Parse the PEM block
	block, _ := pem.Decode(bytesResult)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("failed to parse the public key: %v", err)
	}

	// Parse the DER-encoded public key
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse the public key: %v", err)
	}

	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("failed to cast public key to RSA public key")
	}

	return rsaPublicKey, nil
}

func (c *Client) parseSignature(ctx context.Context, signature string) (*jwt.Token, error) {
	publicKey, err := c.GetWebhookPublicKey(ctx)
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(signature, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (c *Client) ValidateSignature(ctx context.Context, signature string) (bool, error) {
	if token, err := c.parseSignature(ctx, signature); err == nil && token.Valid {
		return true, nil
	} else {
		return false, err
	}
}

func (c *Client) ValidateBody(ctx context.Context, signature string, body string) (bool, error) {
	if token, err := c.parseSignature(ctx, signature); err == nil && token.Valid {
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return false, errors.New("failed to cast claims to jwt.MapClaims")
		}

		return claims["data"] == body, nil
	} else {
		return false, err
	}
}
