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
var executable string = "./fakerData"
var defaultFile string = "fakeProducts.txt"
var defaultNumProducts int= 0
var defaultSeed int64 = 7

// TempFileName generates a temporary filename for use in testing or whatever
// from https://stackoverflow.com/questions/28005865/golang-generate-unique-filename-with-extension
func TempFileName(prefix, suffix string) string {
    randBytes := make([]byte, 16)
    rand.Read(randBytes)
    return prefix+hex.EncodeToString(randBytes)+suffix
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

// Testing creation of a specific file containing a specific number of products
func TestCreateFileSpecificLengthAndName(t *testing.T) {
    tempFileName := TempFileName("", ".txt")
    defer os.Remove(tempFileName)
    productCount := rand.Intn(100)

    cmd := exec.Command(executable, "-rows="+strconv.Itoa(productCount),
                        "-seed=5", "-file="+tempFileName)
    err := cmd.Run()

    dat, err := ioutil.ReadFile(tempFileName)
    check(err)

    if strings.Count(string(dat), "\n") != productCount {
        t.Errorf("Specific execution of %s did not produce file with %d products",
                executable, productCount)
    }
}

// Testing the random number seed function
func TestCreateFileSpecificLengthNameAndSeed(t *testing.T) {
    tempFileName := TempFileName("", ".txt")
    defer os.Remove(tempFileName)
    productCount := 1
    productString := "Yodoo GPS Input Amplifier,81"

    cmd := exec.Command(executable, "-rows="+strconv.Itoa(productCount),
                        "-seed=1", "-file="+tempFileName)
    err := cmd.Run()

    dat, err := ioutil.ReadFile(tempFileName)
    check(err)

    if strings.Count(string(dat), "\n") != productCount {
        t.Errorf("Specific execution of %s did not produce file with %d products",
                executable, productCount)
    }

    if !strings.Contains(string(dat), productString){
        t.Errorf("Seeded execution of %s did not produce product %s",
                executable, productString)
    }
}
