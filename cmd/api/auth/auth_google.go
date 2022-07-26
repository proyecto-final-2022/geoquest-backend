package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"
)

type GoogleParameters struct {
	GoogleJWT *string
}

// ValidateGoogleJWT
func getGooglePublicKey(keyID string) (string, error) {
	resp, err := http.Get("https://www.googleapis.com/oauth2/v1/certs")
	if err != nil {
		return "", err
	}
	dat, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	myResp := map[string]string{}
	err = json.Unmarshal(dat, &myResp)
	if err != nil {
		return "", err
	}
	key, ok := myResp[keyID]
	if !ok {
		return "", errors.New("key not found")
	}
	return key, nil
}

func ValidateGoogleJWT(tokenString string) (domain.GoogleClaims, error) {
	claimsStruct := domain.GoogleClaims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		&claimsStruct,
		func(token *jwt.Token) (interface{}, error) {
			pem, err := getGooglePublicKey(fmt.Sprintf("%s", token.Header["kid"]))
			if err != nil {
				return nil, err
			}
			key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(pem))
			if err != nil {
				return nil, err
			}
			return key, nil
		},
	)
	if err != nil {
		return domain.GoogleClaims{}, err
	}

	claims, ok := token.Claims.(*domain.GoogleClaims)
	if !ok {
		return domain.GoogleClaims{}, errors.New("Invalid Google JWT")
	}

	if claims.Issuer != "accounts.google.com" && claims.Issuer != "https://accounts.google.com" {
		return domain.GoogleClaims{}, errors.New("iss is invalid")
	}

	if claims.Audience != "YOUR_CLIENT_ID_HERE" {
		return domain.GoogleClaims{}, errors.New("aud is invalid")
	}

	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return domain.GoogleClaims{}, errors.New("JWT is expired")
	}

	return *claims, nil
}
