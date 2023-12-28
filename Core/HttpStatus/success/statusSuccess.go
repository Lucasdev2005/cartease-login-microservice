package HttpstatusSuccess

const (
	OK               = 200
	Created          = 201
	Accepted         = 202
	NonAuthoritative = 203
	NoContent        = 204
	ResetContent     = 205
	PartialContent   = 206
	MultiStatus      = 207
	AlreadyReported  = 208
	IMUsed           = 226
)

var statusNames = map[int]string{
	OK:               "OK",
	Created:          "Created",
	Accepted:         "Accepted",
	NonAuthoritative: "Non-Authoritative Information",
	NoContent:        "No Content",
	ResetContent:     "Reset Content",
	PartialContent:   "Partial Content",
	MultiStatus:      "Multi-Status",
	AlreadyReported:  "Already Reported",
	IMUsed:           "IM Used",
}

func GetStatusName(code int) string {
	return statusNames[code]
}
