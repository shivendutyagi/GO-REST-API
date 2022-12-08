package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/shivendutyagi/newapi/router"
)

func main() {
	fmt.Println("Connection")
	r := router.Router()
	fmt.Println("server")
	log.Fatal(http.ListenAndServe(":4001", r))
	fmt.Println("listening at port")

}
