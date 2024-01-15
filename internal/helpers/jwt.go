package helpers

import (
	jwtv4 "github.com/golang-jwt/jwt/v4"
)

// const secret = "asdafdsggdqwerdsfffasdasxsd"

// Create the JWT key used to create the signature
var jwtKey = []byte("asdafdsggdqwerdsfffasdasxsd")

// Claims a struct that will be encoded to a JWT.
// We add jwt.RegisteredClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	UserID int64 `json:"user_id"`
	RoleID int64 `json:"role_id"`
	jwtv4.RegisteredClaims
}

//func (claims Claims) Valid() error {
//	var now = time.Now().UTC().Unix()
//	if claims.VerifyExpiresAt(now, true) {
//		return nil
//	}
//	return fmt.Errorf("invalid token")
//}

// JwtSign - sign with key.
func JwtSign(claims *Claims) (string, error) {
	var tokenString, err = jwtv4.NewWithClaims(jwtv4.SigningMethodHS256, claims).SignedString(jwtKey)

	return tokenString, err
}

// JwtParse - parse and check key.
func JwtParse(tokenString string) (*jwtv4.Token, error) {
	return jwtv4.ParseWithClaims(tokenString, &Claims{}, jwtKeyFunc)
}

func jwtKeyFunc(token *jwtv4.Token) (interface{}, error) {
	return jwtKey, nil
}
