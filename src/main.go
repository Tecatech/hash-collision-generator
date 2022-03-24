package main

import (
    "fmt"
)

func main() {
    fmt.Println("Birthday paradox attack:")
    sc := birthdayParadoxWrapper()
    writeBytes("birthdayParadox.txt", sc)
    
    fmt.Println("Pollard rho algorithm:")
    sc = rhoPollardWrapper()
    writeBytes("rhoPollard.txt", sc)
}