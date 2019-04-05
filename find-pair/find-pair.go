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

func main() {
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
    scanner := bufio.NewScanner(file)
    products := []product{}
    for scanner.Scan() {
        lineItem := scanner.Text()
        check(err)
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

// setup shrinking window
    top := len(products)-1
    bottom := 0
    topPrice := products[top].price
    currentPrice := 0
    attemptedPrice := products[bottom].price

    for top != bottom {
        pairPrice := topPrice + currentPrice
        peekPrice := topPrice + attemptedPrice

        if pairPrice <= peekPrice && peekPrice <= cardBalance {
            bottom++
            currentPrice = products[bottom-1].price
            attemptedPrice = products[bottom].price
        } else if pairPrice > cardBalance {
            top--
            topPrice = products[top].price
        } else {
            if bottom < 0 {
                bottom++
                currentPrice = products[bottom-1].price
                attemptedPrice = products[bottom].price
                continue
            }
            fmt.Println("total price="+strconv.Itoa(pairPrice)+ " within " + strconv.Itoa(cardBalance-pairPrice) +" cents")

            fmt.Println(products[top].name, products[top].price, ",", products[bottom].name, products[bottom].price)
            return
        }
    }
    fmt.Println("Not Possible")
}
