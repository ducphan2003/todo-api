package tokenprovider

type Provider interface {
	Generate(data TokenPayload, expiry int, inputKey string) (*Token, error)
	Validate(token string, publicKey string) (*TokenPayload, error)
}
