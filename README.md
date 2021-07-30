# Blockchain-go

## Steps

### 1 - Go Module system
* We look at the go modules system and build dthe basic framework for a Blockchain  

### 2 - Proof of work
* We look at proof of work and how to implement consensus in a golang blockchain

### 3 - BadgerDB 
* We add persistence via BadgerDB and a Command Line interface to our blockchain application. 
* Commands 
    - go run main.go print
    - go run main.go add -block "first block"
    - go run main.go add -block "Send 10 BTC to Jona"
---

## Imports used
* "bytes"
* "crypto/sha256"
* "encoding/binary"
* "fmt"
* "log"
* "math"
* "math/big"
