package main

import (
	"errors"
	
	
	)

type calculation interface{

	Calcular()func(float64, float64) float64	
	
}


	
func (a Arith) Calcular()func(x float64, y float64) (float64, error){


	if result, ok := calcs[a.Op]; ok{
		return func(x float64, y float64)(float64, error){
			return result(x,y), nil
	}
	}

return func(x float64, y float64)(float64, error){
	return 0, errors.New("Invalid operand request!")
	
}


}
	

