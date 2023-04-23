package golang_web

import (
	"net/http"
	"net/http/httptest"
	"fmt"
	"testing"
	"io"
	"strings"

)

func SayHello(writer http.ResponseWriter, request *http.Request){
	name := request.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(writer, "Hello")
	}else {
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestQueryParameter(t *testing.T){
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000/?name=Luna", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body,_ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func MultipleQueryParameter(writer http.ResponseWriter, request *http.Request){
	firstname := request.URL.Query().Get("firstname")
	lastname := request.URL.Query().Get("lastname")

	fmt.Fprintf(writer, "Hello %s %s", firstname, lastname)
}

func TestMultipleQueryParameter(t *testing.T){
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000/?firstname=Meizaluna&lastname=Wulandari", nil)
	recorder := httptest.NewRecorder()

	MultipleQueryParameter(recorder, request)

	response := recorder.Result()
	body,_ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func MultipleParameterValues(writer http.ResponseWriter, request *http.Request){
	query := request.URL.Query()
	names := query["name"]
	fmt.Fprint(writer, strings.Join(names, " "))
}


func TestMultipleParameterValues(t *testing.T){
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000/?name=Meizaluna&name=Wulandari", nil)
	recorder := httptest.NewRecorder()

	MultipleParameterValues(recorder, request)

	response := recorder.Result()
	body,_ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}