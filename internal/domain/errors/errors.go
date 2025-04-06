package errors

type AppErrors struct {
	Code            ErrorCode
	Errors          []string
	OriginFunc      string
	Cause           string
	FriendlyMessage string
}

func (de AppErrors) Error() string {
	return string(de.Code) + ": " + de.FriendlyMessage
}
