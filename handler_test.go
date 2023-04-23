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

 func TestServeMux(t *testing.T){
 	mux := http.NewServeMux()
 	
 	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request){
 		fmt.Fprint(writer, "Hello World")
 	})
 	mux.HandleFunc("/hi", func(writer http.ResponseWriter, request *http.Request){
 		fmt.Fprint(writer, "hi")
 	})
	mux.HandleFunc("/images/", func(writer http.ResponseWriter, request *http.Request){
 		fmt.Fprint(writer, "Images")
 	})
 	mux.HandleFunc("/images/thumbnails/", func(writer http.ResponseWriter, request *http.Request){
 		fmt.Fprint(writer, "Thumbnails")
 	})

 	server := http.Server{
	 	Addr: "localhost:8000",
	 	Handler: mux,
	 }
	 err := server.ListenAndServe()
	 if err != nil {
	 	panic(err)
	 }
 }