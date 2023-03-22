package routes

import (
	"net/http"

	"github.com/APouzi/golang-shop/routes/index"
	"github.com/julienschmidt/httprouter"
)

func StartRouter(){
	router := httprouter.New()
	
	//index
	index.HomePage(router)
	

	http.ListenAndServe(":8000", router)
}