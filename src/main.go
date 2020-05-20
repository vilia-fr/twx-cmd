package main

import (
    "flag"
    "fmt"
    "./api"
)

func parseParams() (operationPtr, urlPtr, appKeyPtr, usernamePtr, passwordPtr *string, isVerbose *bool) {
    operationPtr = flag.String("op", "get-version", "Operation")
    urlPtr = flag.String("url", "http://localhost:8080/Thingworx", "ThingWorx base URL")
    appKeyPtr = flag.String("appKey", "", "ThingWorx AppKey")
    usernamePtr = flag.String("username", "Administrator", "ThingWorx user")
    passwordPtr = flag.String("password", "", "ThingWorx password")
    isVerbose = flag.Bool("verbose", false, "Verbose output mode")
    flag.Parse()
    return
}

func main() {
    operationPtr, urlPtr, appKeyPtr, usernamePtr, passwordPtr, isVerbose := parseParams()
    twx := api.Api { *usernamePtr, *passwordPtr, *appKeyPtr, *urlPtr, *isVerbose }

	switch *operationPtr {
	case "get-version":
		if *isVerbose {
			fmt.Println("Getting platform version...")
		}
		fmt.Println(twx.GetVersion())
	default:
		fmt.Println("ERROR: Unrecognized operation:", *operationPtr)
	}    
}
