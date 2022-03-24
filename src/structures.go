package main

import (
    "bytes"
    "encoding/hex"
    
    "github.com/tunabay/go-bitarray"
)

type bitArray = *bitarray.BitArray

type node struct {
    k, z []byte
    j    int
}

func makeNode(k, z []byte, j int) node {
    return node {
        k: k,
        j: j,
        z: z,
    }
}

type pair struct {
    x, y []byte
}

func nilPair() pair {
    return pair{}
}

func makePair(x, y []byte) pair {
    return pair {
        x: x,
        y: y,
    }
}

func (p *pair) isEqual() bool {
    return bytes.Equal(p.x, p.y)
}

func (p *pair) toString() string {
    return hex.EncodeToString(p.x) + "\t" + hex.EncodeToString(p.y)
}