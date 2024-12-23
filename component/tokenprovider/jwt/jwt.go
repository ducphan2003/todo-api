package jwt

import (
	"encoding/base64"
	"time"
	"todo-api/component/tokenprovider"

	"github.com/golang-jwt/jwt/v4"
)

type jwtProvider struct {
	secret string
	public string
}

func NewTokenJWTProvider(secret string, public string) *jwtProvider {
	return &jwtProvider{
		secret: secret,
		public: public,
	}
}

type myClaims struct {
	Payload tokenprovider.TokenPayload `json:"payload"`
	jwt.StandardClaims
}

func (j *jwtProvider) Generate(data tokenprovider.TokenPayload,
	expiry int,
	inputKey string,
) (*tokenprovider.Token, error) {
	decodedPrivateKey, err := base64.StdEncoding.DecodeString(inputKey)
	if err != nil {
		return nil, err
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(decodedPrivateKey)
	if err != nil {
		return nil, err
	}

	t := jwt.NewWithClaims(jwt.SigningMethodRS256, myClaims{
		data,
		jwt.StandardClaims{
			ExpiresAt: time.Now().UTC().Add(time.Second * time.Duration(expiry)).Unix(),
			IssuedAt:  time.Now().UTC().Unix(),
		},
	})

	//token, err := t.SignedString([]byte(j.secret))
	token, err := t.SignedString(key)

	if err != nil {
		return nil, err
	}

	return &tokenprovider.Token{
		Token:     token,
		Expiry:    expiry,
		CreatedAt: time.Now().UTC(),
	}, nil
}

func (j *jwtProvider) Validate(myToken string, publicKey string) (*tokenprovider.TokenPayload, error) {
	decodedPublicKey, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return nil, err
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(decodedPublicKey)
	if err != nil {
		return nil, err
	}

	res, err := jwt.ParseWithClaims(myToken, &myClaims{}, func(token *jwt.Token) (interface{}, error) {
		//return []byte(j.secret), nil
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, err
		}
		return key, nil
	})

	if err != nil {
		return nil, tokenprovider.ErrInvalidToken
	}
	if !res.Valid {
		return nil, tokenprovider.ErrInvalidToken
	}
	claims, ok := res.Claims.(*myClaims)
	if !ok {
		return nil, tokenprovider.ErrInvalidToken
	}
	return &claims.Payload, nil
}

func (j *jwtProvider) String() string {
	return "JWT implements Provider"
}
