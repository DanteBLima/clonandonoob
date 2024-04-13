package main

import (
	
	"net/http"
	"strconv"
	"strings"
	"fmt"
	"time"
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







type Calculator int

func (Calculator) ServeHTTP(w http.ResponseWriter, r *http.Request){
		

query := r.URL.Query().Get("op")
divide := strings.Fields(query)

if len(divide) != 3{

	fmt.Fprintf(w, "Invalid format. :number operator number")
	return

}

x1, err1:= strconv.ParseFloat(divide[0],64)
y1, err2:= strconv.ParseFloat(divide[2],64)

if err1 != nil || err2 != nil{
	http.Error(w, err1.Error(),http.StatusBadRequest)
	return
}

a := Arith{
	
	x: x1,
	Op: divide[1],
	y: y1,

}

calcFunc := a.Calcular()
result, err := calcFunc(x1,y1)
if err != nil{

	http.Error(w, err.Error(),http.StatusBadRequest)

}
fmt.Fprintf(w,"Result: %.2f",result)



}

func NewServer(c Calculator) *http.Server{

	return &http.Server{
	
	Addr: ":8080",
	Handler: c,
	ReadTimeout: 1 * time.Second,
	WriteTimeout: 1 * time.Second,
	
	
	
	}



}













