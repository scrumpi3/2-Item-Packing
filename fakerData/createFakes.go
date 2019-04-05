package main

import (
    "fmt"
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
    err := fmt.Errorf("Missing arguments. createFakes <# of rows> <output file>")
        fmt.Println(err.Error())
        return
    }

    file, err := os.Create(os.Args[2])
    check(err)
    defer file.Close()

    numberOfLines, err := strconv.Atoi(os.Args[1])
    check(err)

    for numberOfLines > 0 {
        fakeItem := fake.Product() + "," + fake.Digits() + "\n"
        _, err := file.WriteString(fakeItem)
        check(err)
        fmt.Println(fakeItem)
        numberOfLines--
    }
}
