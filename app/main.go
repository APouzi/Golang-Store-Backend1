package main

import (
	"fmt"

	"github.com/APouzi/golang-shop/routes"
)


func main(){
	routes.StartRouter()
	fmt.Println("Hit")
	// fmt.Println("test", reflect.TypeOf(router))
}