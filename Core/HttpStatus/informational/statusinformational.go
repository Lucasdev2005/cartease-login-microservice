package HttpStatusinformational

// Enumeration of HTTP informational status codes
const (
	Continue           = 100
	SwitchingProtocols = 101
	Processing         = 102
)

// Custom names for HTTP informational status codes
var statusNames = map[int]string{
	Continue:           "Continue",
	SwitchingProtocols: "Switching Protocols",
	Processing:         "Processing",
}

// GetStatusName returns the custom name of an HTTP status code
func GetStatusName(code int) string {
	return statusNames[code]
}
