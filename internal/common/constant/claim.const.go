package constant

const (
	IDKey       CtxKey = "id"
	UsernameKey CtxKey = "username"
)

var (
	UserClaimKeys = []CtxKey{IDKey, UsernameKey}
)
