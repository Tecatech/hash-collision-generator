package main

import (
    "bytes"
    "crypto/rand"
    "crypto/sha256"
    "io"
    "math/big"
    "math/bits"
    "os"
    "strconv"
    
    "github.com/tunabay/go-bitarray"
    "golang.org/x/exp/slices"
)

func deepEqual(p1, p2 pair) bool {
    switch {
    case bytes.Equal(p1.x, p2.x) && bytes.Equal(p1.y, p2.y):
        return true
    case bytes.Equal(p1.x, p2.y) && bytes.Equal(p1.y, p2.x):
        return true
    default:
        return false
    }
}

func genBytes() []byte {
    const minLen = 15
    const maxLen = 90
    
    len, _ := rand.Int(rand.Reader, big.NewInt(maxLen))
    bs := make([]byte, len.Int64() + minLen)
    
    if _, err := io.ReadFull(rand.Reader, bs); err != nil {
        panic(err)
    }
    
    return bs
}

func shaXX(bs []byte, nBits int) bitArray {
    sum := sha256.Sum256(bs)
    ba := bitarray.NewFromBytes(sum[:], 0, nBits)
    
    return ba
}

func padZeros(ba bitArray) []byte {
    const zero = 0
    
    bs, _ := ba.Bytes()
    bs = append(bs, zero)
    
    return bs
}

func leadingZeros(bs []byte) int {
    const bitLen = 8
    
    nb := 0
    
    for i := 0; i < len(bs); i++ {
        nz := bits.LeadingZeros8(bs[i])
        nb += nz
        
        if nz != bitLen {
            return nb
        }
    }
    
    return nb
}

func searchNode(s []node, key []byte) int {
    return slices.IndexFunc(s, func(n node) bool { return bytes.Equal(n.k, key) })
}

func searchColl(c []pair, coll pair) int {
    return slices.IndexFunc(c, func(p pair) bool { return deepEqual(p, coll) })
}

func writeBytes(file string, sc [][]pair) {
    const low = 15
    
    fo, err := os.Create(file)
    if err != nil {
        panic(err)
    }
    
    defer func() {
        if err := fo.Close(); err != nil {
            panic(err)
        }
    }()
    
    for ind, c := range sc {
        if _, err := fo.WriteString("SHA" + strconv.Itoa(ind + low) + "\n"); err != nil {
            panic(err)
        }
        
        for _, coll := range c {
            if _, err := fo.WriteString(coll.toString() + "\n"); err != nil {
                panic(err)
            }
        }
    }
}