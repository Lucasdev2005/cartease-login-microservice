package HttpClientError

const (
	BadRequest                   = 400
	Unauthorized                 = 401
	PaymentRequired              = 402
	Forbidden                    = 403
	NotFound                     = 404
	MethodNotAllowed             = 405
	NotAcceptable                = 406
	ProxyAuthRequired            = 407
	RequestTimeout               = 408
	Conflict                     = 409
	Gone                         = 410
	LengthRequired               = 411
	PreconditionFailed           = 412
	RequestEntityTooLarge        = 413
	RequestURITooLong            = 414
	UnsupportedMediaType         = 415
	RequestedRangeNotSatisfiable = 416
	ExpectationFailed            = 417
	Teapot                       = 418
	UnprocessableEntity          = 422
	Locked                       = 423
	FailedDependency             = 424
	UpgradeRequired              = 426
	PreconditionRequired         = 428
	TooManyRequests              = 429
	RequestHeaderFieldsTooLarge  = 431
	UnavailableForLegalReasons   = 451
)

var statusNames = map[int]string{
	BadRequest:                   "Bad Request",
	Unauthorized:                 "Unauthorized",
	PaymentRequired:              "Payment Required",
	Forbidden:                    "Forbidden",
	NotFound:                     "Not Found",
	MethodNotAllowed:             "Method Not Allowed",
	NotAcceptable:                "Not Acceptable",
	ProxyAuthRequired:            "Proxy Authentication Required",
	RequestTimeout:               "Request Timeout",
	Conflict:                     "Conflict",
	Gone:                         "Gone",
	LengthRequired:               "Length Required",
	PreconditionFailed:           "Precondition Failed",
	RequestEntityTooLarge:        "Request Entity Too Large",
	RequestURITooLong:            "Request URI Too Long",
	UnsupportedMediaType:         "Unsupported Media Type",
	RequestedRangeNotSatisfiable: "Requested Range Not Satisfiable",
	ExpectationFailed:            "Expectation Failed",
	Teapot:                       "I'm a teapot",
	UnprocessableEntity:          "Unprocessable Entity",
	Locked:                       "Locked",
	FailedDependency:             "Failed Dependency",
	UpgradeRequired:              "Upgrade Required",
	PreconditionRequired:         "Precondition Required",
	TooManyRequests:              "Too Many Requests",
	RequestHeaderFieldsTooLarge:  "Request Header Fields Too Large",
	UnavailableForLegalReasons:   "Unavailable For Legal Reasons",
}

func GetStatusName(code int) string {
	return statusNames[code]
}
