package main

import (
    "fmt"
    "advocate-cms-2/internal/dockerman"
)


func main() {

    var err = dockerman.CreateAdvocateDocker() 

    if  err != nil {
        fmt.Println("Error creating docker image")
        fmt.Println(err)
        return
    }

	fmt.Println("Docker started")
}
