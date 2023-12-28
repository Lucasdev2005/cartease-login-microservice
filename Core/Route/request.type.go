package Route

type Request struct {
	Body          interface{}
	GetParam      func(key string) string
	GetQueryParam func(key string) string
	GetHeader     func(key string) string
}
