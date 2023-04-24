package golang_web

import(
	"fmt"
	"io"
	"testing"
	"net/http"
	"net/http/httptest"
)

func SetCookie(writer http.ResponseWriter, request *http.Request){
	cookie := new(http.Cookie)
	cookie.Name = "X-PZN-Name"
	cookie.Value = request.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(writer, cookie)
	fmt.Fprint(writer, "Success create cookie")
}

func GetCookie(writer http.ResponseWriter, request *http.Request){
	cookie, err := request.Cookie("X-PZN-Name")
	if err != nil {
		fmt.Fprint(writer, "no cookie")
	}else{
		name := cookie.Value
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestCookie(t *testing.T){
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr: "localhost:8000",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	} 
}

func TestSetCookie(t *testing.T){
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000/?name=Luna", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)

	cookies := recorder.Result().Cookies()

	for _, cookies := range cookies {
		fmt.Printf("Cookie %s %s \n", cookies.Name, cookies.Value)
	}
}

func TestGetCookie(t *testing.T){
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000/", nil)
	cookie := new(http.Cookie)
	cookie.Name = "X-PZN-Name"
	cookie.Value = "Luna"
	request.AddCookie(cookie)
	recorder := httptest.NewRecorder()

	GetCookie(recorder, request)

	body,_ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}