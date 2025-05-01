package apperrors

type AppErrors struct {
	Code            ErrorCode
	Errors          []string
	Map             map[string]any
	OriginFunc      string
	Cause           string
	FriendlyMessage string
}

func (de AppErrors) Error() string {
	return string(de.Code) + ": " + de.FriendlyMessage
}
