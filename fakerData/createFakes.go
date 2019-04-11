package main

import (
    "flag"
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
    rowsPtr := flag.Int("rows", 0, "number of fake products")
    filePtr := flag.String("file", "fakeProducts.txt", "name of output file")
    seedPtr := flag.Int64("seed", 7, "random number seed")
    flag.Parse()

    file, err := os.Create(*filePtr)
    check(err)
    defer file.Close()

    numberOfLines := *rowsPtr
    seed := *seedPtr
    fake.Seed(seed)
    rand.Seed(seed)

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
