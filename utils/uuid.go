package utils

import (
    "os"
    "fmt"
    "log"
)

var Random *os.File

func InitRandom() {
    f, err := os.Open("/dev/urandom")
    if err != nil {
            log.Fatal(err)
    }
    Random = f
}

func UUID() string {
    InitRandom()
    b := make([]byte, 16)
    Random.Read(b)
    b[6] = (b[6] & 0x0f) | 0x40
    b[8] = (b[8] & 0x3f) | 0x80
    return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}