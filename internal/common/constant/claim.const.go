package constant

const (
	IDKey          CtxKey = "id"
	PhoneNumberKey CtxKey = "phone_number"
)

var (
	UserClaimKeys = []CtxKey{IDKey, PhoneNumberKey}
)
