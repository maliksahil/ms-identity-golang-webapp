# Azure AD + Golang Web App Sample

This sample shows how to use a Go web app to sign into using AAD.

If you are looking for client credentials flow, managed identities or device code flow, you should check out Azure SDK at https://docs.microsoft.com/en-us/azure/developer/go/azure-sdk-authorization

## Instructions

1. Register a new app in Azure AD
2. Add the "web" platform with redirect uri "http://localhost:3000/callback"
3. Generate a client secret
4. Rename .env.sample to .env and update values

*To run* 
1. Run `go build ./...` to build and install the Go dependencies.
2. Run `go run main.go server.go` to start the app 
3. Point your browser to [http://localhost:3000/](http://localhost:3000/).
4. Login using your credentials.
