package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"web/src/config"
	"web/src/cookies"
	"web/src/router"
	"web/src/utils"

	"github.com/gorilla/securecookie"
)

func init() {
	hasKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
	blockKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
	fmt.Println(hasKey)
	fmt.Println(blockKey)
}

func main() {
	config.Load()
	cookies.Configure()

	r := router.Generate()
	utils.LoadTemplate()

	fmt.Printf("Running web application on port %d", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
