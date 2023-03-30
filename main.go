package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/joho/godotenv"
)

const webport = 8080

type Config struct{}

func main(){
	
	app := Config{}
	fmt.Printf("Starting Store Backend on port %d \n", webport)

	serve := &http.Server{
		Addr: fmt.Sprintf(":%d", webport),
		Handler: app.StartRouter(),
	}

	err := serve.ListenAndServe()
	if err != nil{
		log.Panic(err)
	}
	
	// fmt.Println("test", reflect.TypeOf(router))
}


// Initializing the environment variables to run. 
func init() {
    // Get the absolute path to the directory where the executable is located
    exeDir, err := filepath.Abs("./")
    if err != nil {
        log.Fatal(err)
    }

    // Load the .env file from the directory
    err = godotenv.Load(filepath.Join(exeDir, ".env"))
    if err != nil {
        log.Fatal("Error loading .env file\n","exeDir: ", exeDir)
    }
}
