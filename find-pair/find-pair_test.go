package main

import (
    "log"
    "os/exec"
    "strconv"
    "strings"
    "testing"
)


func runCommand(fileName string, cardBalance int) (string){
    const executable string = "./find-pair"
    cmd := exec.Command(executable, fileName, strconv.Itoa(cardBalance))
    out, err := cmd.CombinedOutput()
    if err != nil {
        log.Fatalf("cmd.Run() failed with %s\n", err)
    }
    return string(out)
}

func notPossible(output string) (bool){
    return strings.Contains(output, "Not Possible")
}

//File does not exist
func TestNonexistentFile(t *testing.T) {
    const nonFile = "nonexistentFile.xyz"
    output := runCommand(nonFile, 1000)
    if !strings.Contains(output, "does not exist") || !notPossible(output) {
        t.Errorf("Incorrect outptut for nonexistent file")
    }
}

//File is directory
func TestFileIsDirectory(t *testing.T) {
    const directoryFile = "testData"
    output := runCommand(directoryFile, 1000)
    if !strings.Contains(output, "is a directory") || !notPossible(output) {
        t.Errorf("Incorrect outptut for directory file")
    }
}

//No products
func TestEmptyFile(t *testing.T) {
    const zeroProducts = "testData/zeroProducts.txt"
    output := runCommand(zeroProducts, 1000)
    if !notPossible(output) {
        t.Errorf("Incorrect outptut for empty file")
    }
}

//Not enough products
func TestNotEnoughProducts(t *testing.T) {
    const oneProduct = "testData/oneProduct.txt"
    output := runCommand(oneProduct, 1000)
    if !notPossible(output) {
        t.Errorf("Incorrect outptut for file with one product")
    }
}

//No iteration
func TestAlgorithmWithNoIteration(t *testing.T) {
    const twoProducts  = "testData/twoProducts.txt"
    firstOutput := runCommand(twoProducts, 0)
    if !notPossible(firstOutput) {
        t.Errorf("Card balance set to zero, but a solution was produced")
    }

    secondOutput := runCommand(twoProducts, 1000)
    if notPossible(secondOutput) {
        t.Errorf("Card balance set to 1k, but a solution was not produced")
    }
}

//Large file; find neighboring pairs
func TestFindNeighboringPairs(t *testing.T) {
    const tenKProducts  = "testData/tenKProducts.txt"
    firstOutput := runCommand(tenKProducts, 0)
    if !notPossible(firstOutput) {
        t.Errorf("Card balance set to zero, but a solution was produced")
    }

    secondOutput := runCommand(tenKProducts, 40000)
    thirdOutput  := runCommand(tenKProducts, 39999)
    fourthOutput := runCommand(tenKProducts, 39998)

    if notPossible(secondOutput) || notPossible(thirdOutput) || notPossible(fourthOutput) {
        t.Errorf("Card balance is high, but a solution was not produced")
    }

    if secondOutput == thirdOutput || thirdOutput == fourthOutput {
        t.Errorf("Resulting outputs should all be different")
    }
}

//Odd formatting; eat white space
func TestFileWithWhitespace(t *testing.T) {
    const tenKProducts  = "testData/tenKProductsWithWhiteSpace.txt"
    const tenKProductsWithWhiteSpace  = "testData/tenKProductsWithWhiteSpace.txt"

    firstOutput := runCommand(tenKProducts, 40000)
    secondOutput := runCommand(tenKProductsWithWhiteSpace, 40000)
    if firstOutput != secondOutput {
        t.Errorf("Resulting outputs should be the same")
    }
}

