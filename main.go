package main

import "log"

func main(){

	log.Fatal(NewServer(1).ListenAndServe())
}
