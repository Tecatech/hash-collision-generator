package main

import (
    "log"
    "math"
    "runtime"
    "time"
)

func rhoPollard(s []node, y0 []byte, nBits, q int) pair {
    yi := y0
    iter := 0
    
    coll := nilPair()
    
    for coll.isEqual() {
        for leadingZeros(yi) != q {
            yi = padZeros(shaXX(yi, nBits))
            iter += 1
        }
        
        if ind := searchNode(s, yi); ind != -1 {
            n := s[ind]
            
            y := y0
            z := n.z
            
            d := int(math.Abs(float64(iter - n.j)))
            
            if iter > n.j {
                for i := 0; i < d; i++ {
                    y = padZeros(shaXX(y, nBits))
                }
            } else {
                for i := 0; i < d; i++ {
                    z = padZeros(shaXX(z, nBits))
                }
            }
            
            for shaXX(y, nBits).ToUint64() != shaXX(z, nBits).ToUint64() {
                y = padZeros(shaXX(y, nBits))
                z = padZeros(shaXX(z, nBits))
            }
            
            coll = makePair(y, z)
        }
        
        s = append(s, makeNode(yi, y0, iter))
    }
    
    return coll
}

func rhoPollardWrapper() [][]pair {
    const nThreads = 10
    const low = 15
    const high = 21
    const maxColls = 100
    
    sc := make([][]pair, high - low)
    
    for nBits := low; nBits < high; nBits++ {
        s := []node{}
        q := int(float64(nBits / 2) - math.Log2(nThreads))
        
        nColls := 0
        iter := 0
        
        elapsed := time.Duration(0)
        
        for nColls < maxColls {
            y0 := genBytes()
            
            start := time.Now()
            coll := rhoPollard(s, y0, nBits, q)
            elapsed += time.Since(start)
            
            if !coll.isEqual() {
                if searchColl(sc[nBits - low], coll) == -1 {
                    sc[nBits - low] = append(sc[nBits - low], coll)
                    nColls += 1
                }
            }
            
            iter += 1
        }
        
        var m runtime.MemStats
        runtime.ReadMemStats(&m)
        
        log.Printf("[SHA%d] alloc: %v KiB", nBits, m.Alloc >> 10)
        log.Printf("[SHA%d] time: %s", nBits, elapsed)
    }
    
    return sc
}