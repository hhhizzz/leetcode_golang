package _468

import (
    "fmt"
    "testing"
)

func TestValidIPAddress(t *testing.T) {
    address := validIPAddress("20EE:Fb8:85a3:0:0:8A2E:0370:7334")
    fmt.Println(address)
}
