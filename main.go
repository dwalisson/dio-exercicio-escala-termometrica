package main

import "fmt"

const ebulicaoK = 373.2
const ebulicaoC = ebulicaoK - 273

func main()  {
	
	fmt.Printf("Ponto de ebulição da água em Kelvin = %g. ", ebulicaoK)
	fmt.Printf("Ponto de ebulição da água em Celsius = %g", ebulicaoC)

}