package main

import (
	"net/http"
	//"fmt"	
	"time"
	//"strings"
	"log"
	)
	

var ValidPaths = []string{"/result"}

func isPathV(path string) bool{

	for i := 0; i<len(ValidPaths); i++{
		if (path == ValidPaths[i]){
			return true
		}
		
	}
	return false
}		
	
	

type calculator int

//Handler
func (c *calculator) ServeHTTP(w http.ResponseWriter, r *http.Request){



	if (!isPathV(r.URL.Path)){
	
		w.WriteHeader(418)
		w.Write([]byte("{\"code\" : 418, \"message\": invalid path}"))
	}	
		
}
	
var calc calculator
	
func NewServer(c calculator) *http.Server{	
	
	return &http.Server{
		Addr: ":8080",
		Handler: &calc,
		ReadTimeout: 1 * time.Second,
		WriteTimeout: 1* time.Second,
	
	}	
	
	}
	
	
	
	
	
	
func main(){
	
	log.Fatal(NewServer(1).ListenAndServe())

}
