package main

import (
    "encoding/hex"
    "io/ioutil"
    "math/rand"
    "os"
    "os/exec"
    "strings"
    "strconv"
    "testing"
)

// default executable and input values
const executable string = "./fakerData"
const defaultFile string = "fakeProducts.txt"
const defaultNumProducts int = 0
const defaultSeed int64 = 7

// TempFileName generates a temporary filename for use in testing or whatever
// from https://stackoverflow.com/questions/28005865/golang-generate-unique-filename-with-extension
func TempFileName(prefix, suffix string) string {
    randBytes := make([]byte, 16)
    rand.Read(randBytes)
    return prefix+hex.EncodeToString(randBytes)+suffix
}

// Helper function
// Checks if the actual number of products matches the expected
func checkNumProducts(products []byte, numProducts int, t *testing.T) {
    if strings.Count(string(products), "\n") != numProducts {
        t.Errorf("Execution of %s did not produce a file with %d products",
                    executable, numProducts)
    }
}

// Helper function
// Checks if a named product exists in product list
func checkProductExists(products []byte, productString string , t *testing.T) {
    if !strings.Contains(string(products), productString){
        t.Errorf("Seeded execution of %s did not produce product %s",
                    executable, productString)
    }
}


// Testing default execution
func TestDefaultCreate(t *testing.T) {
    cmd := exec.Command(executable)
    err := cmd.Run()

    dat, err := ioutil.ReadFile(defaultFile)
    check(err)

    if len(string(dat)) != 0 {
        t.Errorf("Default execution of %s did not produce empty file", executable)
    }
}

// Testing creation of a user named file containing a specific number of products
func TestCreateFileSpecificLengthAndName(t *testing.T) {
    tempFileName := TempFileName("test", ".txt")
    defer os.Remove(tempFileName)
    productCount := rand.Intn(1000)

    cmd := exec.Command(executable, "-rows="+strconv.Itoa(productCount),
                        "-file="+tempFileName)
    err := cmd.Run()

    dat, err := ioutil.ReadFile(tempFileName)
    check(err)

    checkNumProducts(dat, productCount, t)
}

// Testing the random number seed function
// Generated product must all ways be the same
func TestCreateFileSpecificLengthNameAndSeed(t *testing.T) {
    tempFileName := TempFileName("test", ".txt")
    defer os.Remove(tempFileName)
    productCount := 1
    productString := "Yodoo GPS Input Amplifier,81"

    cmd := exec.Command(executable, "-rows="+strconv.Itoa(productCount),
                        "-seed=1", "-file="+tempFileName)
    err := cmd.Run()

    dat, err := ioutil.ReadFile(tempFileName)
    check(err)

    checkNumProducts(dat, productCount, t)

    checkProductExists(dat, productString, t)
}
