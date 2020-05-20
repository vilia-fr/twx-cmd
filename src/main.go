package main

import (
    "flag"
    "fmt"
    "./api"
)

func parseParams() (urlPtr, appKeyPtr, usernamePtr, passwordPtr *string) {
    urlPtr = flag.String("url", "http://localhost:8080/Thingworx", "ThingWorx base URL")
    appKeyPtr = flag.String("appKey", "", "ThingWorx AppKey")
    usernamePtr = flag.String("username", "Administrator", "ThingWorx user")
    passwordPtr = flag.String("password", "", "ThingWorx password")
    flag.Parse()
    return
}

func main() {
    urlPtr, appKeyPtr, usernamePtr, passwordPtr := parseParams()

    fmt.Println("Parameters:")
    fmt.Println("- URL:", *urlPtr)
    fmt.Println("- AppKey:", *appKeyPtr)
    fmt.Println("- User:", *usernamePtr)
    fmt.Println("- Password:", *passwordPtr)
    
    twx := api.Api { *usernamePtr, *passwordPtr, *appKeyPtr, *urlPtr }
    
    fmt.Println("- World:", twx.NewApi())
    
}
