package main

import (
    "fmt"
    "advocate-cms-2/internal/dockerman"
)

func main() {
    err := dockerman.StartAdvocate()

    if err != nil {
        fmt.Println("Error:")
        fmt.Println(err)
        return
    }

	fmt.Println("Good test")
}
