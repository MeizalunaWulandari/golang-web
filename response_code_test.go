package golang_web

import(
	"net/http"
	"net/http/httptest"
	"testing"
	"fmt"
	"io"
)

func ResponseCode(writer http.ResponseWriter, request *http.Request){
	name := request.URL.Query().Get("name")

	if name == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "name is emty")
	} else if name == "Andini"{
		writer.WriteHeader(404)
		fmt.Fprint(writer, "name is not found")
	}else{
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestResponseCodeInvalid(t *testing.T){
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8000", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body,_ := io.ReadAll(response.Body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))
}

func TestResponseCodeInvalid2(t *testing.T){
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8000?name=Andini", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body,_ := io.ReadAll(response.Body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))
}

func TestResponseCodeValid(t *testing.T){
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8000?name=Luna", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body,_ := io.ReadAll(response.Body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))
}