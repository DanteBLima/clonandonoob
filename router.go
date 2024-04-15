package main

import (
	
	"net/http"
	"strconv"
	"strings"
	"fmt"
	"time"
	//"log"
	//"net/url"
	"encoding/json"
	
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

if r.Method != http.MethodGet{
	http.Error(w,"code: 404,  error:  not found", http.StatusNotFound)
	return 

}

if (!isPathV(r.URL.Path)) {
		w.WriteHeader(404)
		w.Write([]byte("{\"code\": 404, \"error\": not found}"))
        return
	}


		

query := r.URL.Query().Get("op")




divide := strings.Fields(query)


if len(divide) != 3{
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, "Invalid format. Try: number operator number")
	return

}


x1, err1:= strconv.ParseFloat(divide[0],64)
y1, err2:= strconv.ParseFloat(divide[2],64)

if err1 != nil || err2 != nil{
	http.Error(w,"Error, invalid number type",http.StatusBadRequest)
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

res := Jsonn{

	Result: result,
	Operation: divide,

}

w.WriteHeader(http.StatusOK)

encode := json.NewEncoder(w)
encode.Encode(res)



}

func NewServer(c Calculator) *http.Server{

	return &http.Server{
	
	Addr: ":8080",
	Handler: c,
	ReadTimeout: 1 * time.Second,
	WriteTimeout: 1 * time.Second,
	
	
	
	}



}













