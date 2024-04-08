package main

import (
	"net/http"
	"fmt"	
	)
	
	
		
func Calc (w http.ResponseWriter, r *http.Request){

	fmt.Fprintln (w,"Welcome calc!")
}
	

	
	
	
func main(){

	http.HandleFunc("/calc",Calc)	
	http.ListenAndServe(":8090",nil)


}
