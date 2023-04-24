package golang_web
import (
	"fmt"
	"testing"
	"net/http/httptest"
	"net/http"
	"strings"
	"io"
)

func FormPost(writer http.ResponseWriter, request *http.Request){
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}

	firstName := request.PostForm.Get("first_name")
	lastName := request.PostForm.Get("last_name")

	fmt.Fprintf(writer, "Hello %s %s", firstName, lastName)
}

func TestFormPost(t *testing.T){
	requestBody := strings.NewReader("first_name=Meizaluna&last_name=Wulandari")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8000", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()
	FormPost(recorder, request)

	response := recorder.Result()
	body,_ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}