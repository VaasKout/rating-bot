package network

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"rating-bot/pkg/logger"

	"net/http"
)

type NetworkImpl struct {
	Logger *logger.Logger
}

func New(log *logger.Logger) *NetworkImpl {
	return &NetworkImpl{
		Logger: log,
	}
}

func (net *NetworkImpl) MakeGetRequest(url string) *HttpResponse {
	response, err := http.Get(url)
	if err != nil {
		net.Logger.Log.Error("Error due get request: " + err.Error())
		return &HttpResponse{Error: err}
	}

	return net.handleBody(response)
}

func (net *NetworkImpl) MakePostRequest(httpRequest *HttpRequest) *HttpResponse {
	response, err := http.Post(httpRequest.Url, "application/json", bytes.NewBuffer(httpRequest.Body))
	if err != nil {
		net.Logger.Log.Error("Error due post request: " + err.Error())
		return &HttpResponse{Error: err}
	}

	return net.handleBody(response)
}

func (net *NetworkImpl) handleBody(response *http.Response) *HttpResponse {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			net.Logger.Log.Error("Error on body closure: %s", err.Error())
			return
		}
	}(response.Body)

	body, err := io.ReadAll(response.Body)

	if err != nil {
		res := &HttpResponse{StatusCode: response.StatusCode, Error: err}
		net.Logger.Log.Error(fmt.Sprintf("Body read error: %s, response: %v", err.Error(), res))
		fmt.Println(fmt.Sprintf("BODY READ ERROR: %s", err.Error()))
		return res
	}

	if response.StatusCode != 200 {
		var responseError = errors.New(response.Status)
		res := &HttpResponse{StatusCode: response.StatusCode, Error: responseError, Body: body}
		return res
	}

	return &HttpResponse{StatusCode: response.StatusCode, Body: body}
}
