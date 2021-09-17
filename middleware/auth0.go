package middleware

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gin-gonic/gin"
	cache "github.com/patrickmn/go-cache"
)

type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

var Cache = cache.New(12*time.Hour, 12*time.Hour)

func getPemCert(token *jwt.Token) (string, error) {
	cert := ""
	var jwks = Jwks{}

	if cached, found := Cache.Get("JWKS"); found {
		jwks = cached.(Jwks)
	} else {
		resp, err := http.Get(os.Getenv("ISS") + ".well-known/jwks.json")
		if err != nil {
			return cert, err
		}
		defer resp.Body.Close()

		err = json.NewDecoder(resp.Body).Decode(&jwks)
		if err != nil {
			return cert, err
		}

		Cache.Set("JWKS", jwks, cache.DefaultExpiration)
	}

	for k := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		err := errors.New("unable to find appropriate public key")
		return cert, err
	}

	return cert, nil
}

var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		aud := os.Getenv("AUD")
		// convert audience in the JWT token to []interface{} if multiple audiences
		convAud, ok := token.Claims.(jwt.MapClaims)["aud"].([]interface{})
		if !ok {
			// convert audience in the JWT token to string if only 1 audience
			strAud, ok := token.Claims.(jwt.MapClaims)["aud"].(string)
			// return error if can't convert to string
			if !ok {
				return token, errors.New("invalid audience")
			}
			// return error if audience doesn't match
			if strAud != aud {
				return token, errors.New("invalid audience")
			}
		} else {
			for _, v := range convAud {
				if v == aud {
					break
				} else {
					return token, errors.New("invalid audience")
				}
			}
		}

		iss := os.Getenv("ISS")
		checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, true)
		if !checkIss {
			return token, errors.New("invalid issuer")
		}

		cert, err := getPemCert(token)
		if err != nil {
			return token, err
		}

		result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
		return result, nil
	},
	SigningMethod: jwt.SigningMethodRS256,
})

func Auth0() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the client secret key
		err := jwtMiddleware.CheckJWT(c.Writer, c.Request)
		if err != nil {
			// Token not found
			c.AbortWithStatus(http.StatusUnauthorized)
			// c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.Next()
	}
}
