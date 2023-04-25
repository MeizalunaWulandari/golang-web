package golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateActionIf(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/if.gohtml"))
	t.ExecuteTemplate(writer, "if.gohtml", Page{
		Title: "Template Action If",
		Name:  "Luna",
	})
}

func TestTemplateActionIf(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000", nil)
	recorder := httptest.NewRecorder()
	TemplateActionIf(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionOperator(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/comperator.gohtml"))
	t.ExecuteTemplate(writer, "comperator.gohtml", map[string]interface{}{
		"Title":      "Template Action Operator",
		"FinalValue": 60,
	})
}

func TestTemplateActionOperator(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000", nil)
	recorder := httptest.NewRecorder()
	TemplateActionOperator(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
