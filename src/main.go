package main

import (
    "fmt"
    "path/filepath"
)

func main() {
    fmt.Println("Birthday paradox attack:")
    
    sc := birthdayParadoxWrapper()
    filePath, _ := filepath.Abs("../data/birthdayParadox.txt")
    writeBytes(filePath, sc)
    
    fmt.Println("Pollard rho algorithm:")
    
    sc = rhoPollardWrapper()
    filePath, _ = filepath.Abs("../data/rhoPollard.txt")
    writeBytes(filePath, sc)
}