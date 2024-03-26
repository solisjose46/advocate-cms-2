package main

import (
    "fmt"
    // "advocate-cms-2/internal/dockerman"
    "advocate-cms-2/internal/dao"
)

// func testDockermanPreview() {
//     err := dockerman.StartAdvocatePreview()
//     if err != nil {
//         fmt.Println("Error with advocate preview", err)
//         return
//     }
// }

func testDbInit() {
    err := dao.DatabaseInit()

    if err != nil {
        fmt.Println("Error with database init", err)
        return
    }
}

func main() {

    testDbInit()

	fmt.Println("Advocate cms 2")
}
