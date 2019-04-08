package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

type product struct {
    name string
    price int
}

// reads file from string and returns array of product
// can return nil
func readProductFile(file *os.File) ([]product) {
    // read input file
    scanner := bufio.NewScanner(file)
    products := []product{}
    for scanner.Scan() {
        lineItem := scanner.Text()
        itemElements := strings.Split(lineItem, ",")
        if len(itemElements) == 1 {
            fmt.Println("only 1 item")
            break
        }
        productPrice, err := strconv.Atoi(itemElements[1])
        check(err)
        product := product{itemElements[0], productPrice}
        products = append(products, product)
    }
    return products
}


func main() {
    // handle argument input
    if len(os.Args) < 3 {
    err := fmt.Errorf("Missing arguments. find-pair <input file> <gift card balance>")
        fmt.Println(err.Error())
        return
    }
    file, err := os.Open(os.Args[1])
    check(err)
    defer file.Close()
    cardBalance, err := strconv.Atoi(os.Args[2])
    check(err)

    // read input file
    products := readProductFile(file)

    // setup shrinking window
    top := len(products)-1
    bottom := 0

    // check if enough products
    if !(top > bottom) {
        //not enough products
        fmt.Println("Not Possible")
        return
    }

    // algorithm best pair memory
    // initial pair is invalid by size
    var bestProductTop = len(products)
    var bestProductBottom = 0
    bestDIfference := cardBalance

    // loop on shrinking window
    for top != bottom {
        topPrice := products[top].price
        currentPrice := products[bottom].price
        attemptedPrice := products[bottom+1].price

        pairPrice := topPrice + currentPrice
        peekPrice := topPrice + attemptedPrice
        currentDifference := cardBalance - pairPrice

        if currentDifference < bestDIfference && currentDifference >= 0 {
            bestProductTop = top
            bestProductBottom = bottom
            bestDIfference = cardBalance - pairPrice
        }

        // shrinking window algorithm
        if peekPrice > cardBalance {
            top--
        } else {
            bottom++
        }
    }

    // final output
    if bestProductTop ==  len(products) {
        //no product pairs
        fmt.Println("Not Possible")
    } else {
        if bottom == len(products) {
            bestProductBottom--
        }
        fmt.Println(products[bestProductTop].name, products[bestProductTop].price, ",",
                    products[bestProductBottom].name, products[bestProductBottom].price)
    }
}
