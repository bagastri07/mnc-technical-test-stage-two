package constant

import "time"

const (
	EnvDevelopment = "development"
	EnvStaging     = "staging"

	DB CtxKey = "database"

	SuccessMessage = "SUCCESS"

	AsynqDefaultRetention = 5 * time.Minute
	AsynqDefaultMaxRetry  = 3
)
