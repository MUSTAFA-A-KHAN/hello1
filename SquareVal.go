package main
 
import "fmt"

func SquareValgo(){
	a:=4
	squareVal(a)
	fmt.Println(a)
	squarevariableVAl(&a)
	fmt.Println(a)
}

func squareVal(v int){
	v*=v
	fmt.Println(&v , v)
}


func squarevariableVAl(v *int){
	*v *= *v
	fmt.Println(&v , *v)
}