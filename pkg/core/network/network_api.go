package network

type NetworkApi interface {
	MakeGetRequest(url string) *HttpResponse
	MakePostRequest(httpRequest *HttpRequest) *HttpResponse
}
