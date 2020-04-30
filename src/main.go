package main

import (
    "flag"
    "fmt"
)

func main() {

    urlPtr := flag.String("url", "http://localhost:8080/Thingworx", "ThingWorx base URL")
    appKeyPtr := flag.String("appKey", "", "ThingWorx AppKey")
    usernamePtr := flag.String("username", "Administrator", "ThingWorx user")
    passwordPtr := flag.String("password", "", "ThingWorx password")

    flag.Parse()

    fmt.Println("url:", *urlPtr)
    fmt.Println("appKeyPtr:", *appKeyPtr)
    fmt.Println("usernamePtr:", *usernamePtr)
    fmt.Println("passwordPtr:", *passwordPtr)
    fmt.Println("tail:", flag.Args())
}