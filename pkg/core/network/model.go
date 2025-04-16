package network

type HttpResponse struct {
	StatusCode int
	Body       []byte
	Error      error
}

type HttpRequest struct {
	Url  string
	Body []byte
}
