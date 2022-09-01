# twilio-azure-functions-notification
Serverless function that sends a text message when called with a payload from a GitHub Webhook

# Quickstart

1) Clone this repository
2) Open project in VSCode
3) Create a Function App in Azure (windows or linux based)
4) Create a binary of the local project. If the Function App in Azure is Windows based, run `GOOS=windows GOARCH=amd64 go build handler.go` command. If the Function App in Azure is Linux based then run `go build handler.go`
5) Ensure the `host.json` "defaultExecutablePath" reflects the correct binary
6) Using the Azure Functions App extention in VSCode, push the local function to the remote function you have created