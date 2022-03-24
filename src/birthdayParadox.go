package main

import (
    "log"
    "runtime"
    "time"
)

func birthdayParadox(nBits, maxColls int) []pair {
    s := make(map[uint64][]byte)
    c := make([]pair, maxColls)
    
    nColls := 0
    
    for nColls < maxColls {
        x := genBytes()
        hx := shaXX(x, nBits).ToUint64()
        
        if y, found := s[hx]; found {
            c[nColls] = makePair(x, y)
            nColls += 1
        }
        
        s[hx] = x
    }
    
    return c
}

func birthdayParadoxWrapper() [][]pair {
    const low = 15
    const high = 21
    const maxColls = 100
    
    sc := make([][]pair, high - low)
    
    for nBits := low; nBits < high; nBits++ {
        start := time.Now()
        sc[nBits - low] = birthdayParadox(nBits, maxColls)
        elapsed := time.Since(start)
        
        var m runtime.MemStats
        runtime.ReadMemStats(&m)
        
        log.Printf("[SHA%d] alloc: %v KiB", nBits, m.Alloc >> 10)
        log.Printf("[SHA%d] time: %s", nBits, elapsed)
    }
    
    return sc
}