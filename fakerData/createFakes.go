package main

import (
    "fmt"
    "math/rand"
    "os"
    "strconv"

    "github.com/icrowley/fake"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    if len(os.Args) < 3 {
    err := fmt.Errorf("Missing arguments. fakerData <# of rows> <output file>")
        fmt.Println(err.Error())
        return
    }

    file, err := os.Create(os.Args[2])
    check(err)
    defer file.Close()

    numberOfLines, err := strconv.Atoi(os.Args[1])
    check(err)

    var price = rand.Intn(100)
    for numberOfLines > 0 {
        fakeItem := fake.Product() + "," + strconv.Itoa(price) + "\n"
        _, err := file.WriteString(fakeItem)
        check(err)
        fmt.Println(fakeItem)
        price += rand.Intn(100)
        numberOfLines--
    }
}
