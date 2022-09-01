package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

type Req struct {
	Action string `json:"action"`
}

func sendMessage(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	jsonData := []byte(reqBody)

	var req Req
	err2 := json.Unmarshal(jsonData, &req)
	if err2 != nil {
		log.Println(err)
	}

	fmt.Println("Action: " + req.Action)

	if req.Action == "review_requested" {
		twilioNumber := os.Getenv("TWILIO_NUMBER")
		recipientNumber := os.Getenv("RECIPIENT_NUMBER")

		accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
		authToken := os.Getenv("TWILIO_AUTH_TOKEN")

		client := twilio.NewRestClientWithParams(twilio.ClientParams{
			Username: accountSid,
			Password: authToken,
		})

		params := &openapi.CreateMessageParams{}
		params.SetTo(recipientNumber)
		params.SetFrom(twilioNumber)
		params.SetBody("You have a new GitHub notification! You have been requested to review a Pull Request.")

		resp, err := client.Api.CreateMessage(params)
		if err != nil {
			fmt.Println(err.Error())
			err = nil
			fmt.Fprint(w, "Message not sent")
		} else {
			fmt.Println("Message Sid: " + *resp.Sid)
			fmt.Fprint(w, "Message sent successfully")
		}
	} else {
		fmt.Fprint(w, "Function invoked but no message was sent")
	}
}

func main() {

	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}

	http.HandleFunc("/api/twilionotification", sendMessage)
	log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
