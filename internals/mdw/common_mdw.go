package mdw

// GetCommonMdw - returns common middleware that is used across handlers
func GetCommonMdw() []Middleware {
	return []Middleware{
		PanicRecover,
		SecureHeaders,
		SetCommonHeaders,
		LogRequest,
	}
}
