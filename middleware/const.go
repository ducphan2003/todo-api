package middleware

type ORDER_ENUM int

const (
	ADMIN ORDER_ENUM = iota
	OWNER
	REQUESTER
)
