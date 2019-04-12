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
// can return empty array
func readProductFile(file *os.File) ([]product) {
    // read input file
    scanner := bufio.NewScanner(file)
    products := []product{}
    for scanner.Scan() {
        lineItem := scanner.Text()
        itemElements := strings.Split(lineItem, ",")  // assumed column delimiter is a comma
        if len(itemElements) == 1 {
            fmt.Println("only 1 item")
            break
        }
        // trim leading and trailing whitespace
        itemElements[0] = strings.TrimSpace(itemElements[0])
        itemElements[1] = strings.TrimSpace(itemElements[1])

        productPrice, err := strconv.Atoi(itemElements[1])
        check(err)
        product := product{itemElements[0], productPrice}
        products = append(products, product)
    }
    return products
}

func inputArgs() (cardBalance int, products []product) {
    // handle argument input
    if len(os.Args) < 3 {
        err := fmt.Errorf("Missing arguments. find-pair <input file> <gift card balance>")
        fmt.Println(err.Error())
        return
    }

    fileName := os.Args[1]

    // check is file exists and is not a directory
    fileInfo, err := os.Stat(fileName)
    if os.IsNotExist(err) {
        fmt.Println("File", fileName, "does not exist")
        return cardBalance, products
    } else if fileInfo.IsDir() {
        fmt.Println("File", fileName, "is a directory")
        return cardBalance, products
    }

    file, err := os.Open(fileName)
    check(err)
    defer file.Close()

    cardBalance, err = strconv.Atoi(os.Args[2])
    check(err)

    // read input file
    products = readProductFile(file)
    return cardBalance, products
}

func main() {
    cardBalance, products := inputArgs()

    // setup shrinking window
    top := len(products)-1
    bottom := 0

    // check if enough products
    if !(top > bottom) {
        //not enough products
        fmt.Println("Not Possible")
        return
    }

    // memory for best pair algorithm
    // initial pair is invalid by size
    var bestProductTop = len(products)
    var bestProductBottom = 0
    bestDIfference := cardBalance

    // loop on shrinking window
    for top != bottom {
        // lookup prices
        topPrice := products[top].price
        currentPrice := products[bottom].price
        attemptedPrice := products[bottom+1].price

        // summed prices and difference
        pairPrice := topPrice + currentPrice
        peekPrice := topPrice + attemptedPrice
        currentDifference := cardBalance - pairPrice

        // track best product pair
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

    // print final output
    if bestProductTop ==  len(products) {
        //no product pairs
        fmt.Println("Not Possible")
    } else {
        if bottom == len(products) {
            // maximum pair
            bestProductBottom--
        }
        fmt.Println(products[bestProductTop].name, products[bestProductTop].price, ",",
                    products[bestProductBottom].name, products[bestProductBottom].price)
    }
}
