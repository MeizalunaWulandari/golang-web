package golang_web

import(
	"net/http"
	"testing"
	_"embed"
	"fmt"
)

func ServeFile(writer http.ResponseWriter, request *http.Request){
	if request.URL.Query().Get("name") != "" {
		http.ServeFile(writer, request, "./resources/ok.html")
	}else{
		http.ServeFile(writer, request, "./resources/notfound.html")
	}
}

func TestServeFileServer(t *testing.T){
	server := http.Server{
		Addr: "localhost:8000",
		Handler: http.HandlerFunc(ServeFile),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources/ok.html
var resourceOk string

//go:embed resources/notfound.html
var resourceNotFound string

func ServeFileEmbed(writer http.ResponseWriter, request *http.Request){
	if request.URL.Query().Get("name") != "" {
		fmt.Fprint(writer, resourceOk)
	}else{
		fmt.Fprint(writer, resourceNotFound)
	}
}

func TestServeFileServerEmbed(t *testing.T){
	server := http.Server{
		Addr: "localhost:8000",
		Handler: http.HandlerFunc(ServeFileEmbed),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
