package golang_web

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"html/template"
	"io"
	"fmt"
)

func SimpleTemplate(writer http.ResponseWriter, request *http.Request){
	templateText := `<html><body>{{.}}</body></html>`
	// t, err := template.New("SIMPLE").Parse(templateText)
	// if err != nil {
	// 	panic(err)
	// }
	t := template.Must(template.New("SIMPLE").Parse(templateText))
	t.ExecuteTemplate(writer, "SIMPLE", "Hello HTML Template")
}

func TestSimpleTemplate(t *testing.T){
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000", nil)
	recorder := httptest.NewRecorder()
	SimpleTemplate(recorder, request)
	body,_ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}