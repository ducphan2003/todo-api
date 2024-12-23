package common

const CurrentUser = "user"
const ConversationID = "conversation_id"

type Requester interface {
	GetUserId() uint
	GetEmail() string
	GetCode() string
}
type TransactionContextKey struct {
}
