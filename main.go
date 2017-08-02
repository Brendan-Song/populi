package main

import (
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {

}

// main creates and starts a Server listening.
func main() {
	// const address = "localhost:8080"
	// // read credentials from environment variables if available
	// config := &user.Config{
	// 	FacebookClientID:     os.Getenv("FACEBOOK_CLIENT_ID"),
	// 	FacebookClientSecret: os.Getenv("FACEBOOK_CLIENT_SECRET"),
	// }
	// // allow consumer credential flags to override config fields
	// clientID := flag.String("client-id", "", "Facebook Client ID")
	// clientSecret := flag.String("client-secret", "", "Facebook Client Secret")
	// flag.Parse()
	// if *clientID != "" {
	// 	config.FacebookClientID = *clientID
	// }
	// if *clientSecret != "" {
	// 	config.FacebookClientSecret = *clientSecret
	// }
	// if config.FacebookClientID == "" {
	// 	log.Fatal("Missing Facebook Client ID")
	// }
	// if config.FacebookClientSecret == "" {
	// 	log.Fatal("Missing Facebook Client Secret")
	// }

	router := NewRouter()

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
