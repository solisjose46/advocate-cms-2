package main

import (
    "fmt"
    "advocate-cms-2/internal/dockerman"
)

func main() {
    err := dockerman.StartAdvocatePreview()

    if err != nil {
        fmt.Println("Error with advocate preview", err)
        return
    }

	fmt.Println("Good test")
}
