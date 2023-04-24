package golang_web

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"fmt"
	"io"
)

func RequestHeader(writer http.ResponseWriter, request *http.Request){
	contentType := request.Header.Get("content-type")
	fmt.Fprint(writer, contentType)
}

func TestRequestHeader(t *testing.T){
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8000/", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	
	RequestHeader(recorder, request)

	response := recorder.Result()
	body,_ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func ResponseHeader(writer http.ResponseWriter, request *http.Request){
	writer.Header().Add("X-Powered-By", "Programmer Zaman Now")
	fmt.Fprint(writer, "OK")
}

func TestResponseHeader(t *testing.T){
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8000/", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	
	ResponseHeader(recorder, request)

	response := recorder.Result()
	body,_ := io.ReadAll(response.Body)

	fmt.Println(string(body))
	fmt.Println(response.Header.Get("X-Powered-By"))
}