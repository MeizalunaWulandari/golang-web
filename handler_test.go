 package golang_web

 import(
 	"net/http"
 	"testing"
 	"fmt"
 )

 func TestHandler(t *testing.T){
 	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request){
	 	fmt.Fprint(writer, "Hello World")
	 }

	 server := http.Server{
	 	Addr: "localhost:8000",
	 	Handler: handler,
	 }
	 server.ListenAndServe()
 }