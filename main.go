package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"challenge/controller"

	"github.com/didip/tollbooth/v8"
)

func main() {
	args := os.Args[1]

	// TODO: Add context
	message := `"Message": {
        "Status": "Request Failed",
        "Body":   "The API is at capacity, try again later.",
    }`
	jsonMessage, _ := json.Marshal(message)

	lmt := tollbooth.NewLimiter(1, nil)

	lmt.SetMessageContentType("application/json")
	lmt.SetMessage(string(jsonMessage))

	router := controller.NewRouter(args)

	http.Handle("/", tollbooth.HTTPMiddleware(lmt)(router))

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", router)
}
