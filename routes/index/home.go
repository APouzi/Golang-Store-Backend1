package index

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type IndexResponse struct{
	Message string `json:"message"`
	StoreOpen bool `json:"storeOpen"`
}

func HomePage(router *httprouter.Router){
	router.GET("/",Index)
}

func Index(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	w.Header().Set("Content-Type","application/json")
	response := IndexResponse{Message:"This is a message", StoreOpen: true}
	
	err := json.NewEncoder(w).Encode(response)

	if err != nil {
		fmt.Println("fail")
	}
	
}