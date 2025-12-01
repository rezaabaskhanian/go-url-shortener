package param

type UrlRequest struct {
	Original  string `json:"original"`
	ShortCode string `json:"short_code"`
	ExpireAt  string `json:"expire_at"`
}
