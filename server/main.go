package main

import (
	"fmt"
	"github.com-NagarjunMallesh/NagarjunMallesh/golang-react-rodo.git/router"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	fmt.Println("starting the server on port 9000!!!")

	log.Fatal(http.ListenAndServe(":9000", r))

}
