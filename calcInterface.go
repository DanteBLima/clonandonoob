package main

import (
	
	)

type calculation interface{

	Calcular()func(float64, float64)float64

}

func (a Arith) Calcula()func(x float64, y float64) float64{
		return calcs[a.Op]
}



