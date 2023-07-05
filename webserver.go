package main
import (
	"fmt"
	"log"
	"net/http"
)
func main(){
	//Define the route handler function
	handler := func(w http.ResponseWriter,r *http.Request ){
		fmt.Fprintf(w,"hello world")
	}

	//create a new server mux
	mux := http.NewServeMux()
	mux.HandleFunc("/",handler)
	server := &http.Server{
		Addr:":8080",
		Handler:mux,
	}
	log.Println("Server listening on port 8080...")
	log.Fatal(server.ListenAndServe())
}