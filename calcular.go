package main

import (
	"math"


)

var calcs = map[string]func(x float64, y float64) float64{

	"sum": sum,
	"+"  : sum,
	"sub": sub,
	"-"  : sub,
	"mul": mult,
	"*"  : mult,
	"div": divi,
	"/"  : divi,
	"pow": pot,
	"^"  : pot,
	"rot": raiz,
	"&"  : raiz,
}

func sum(x float64, y float64) float64{
	return x + y

}

func sub(x float64, y float64) float64{
	return x - y
}

func mult(x float64, y float64) float64{
	return x * y
}

func divi(x float64, y float64) float64{
	return x / y
}

func pot(x float64, y float64) float64{
	return math.Pow(x,y)
}

func raiz(x float64, y float64) float64{
	return math.Pow(y,1/x)
}

