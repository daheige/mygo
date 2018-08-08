package main

import (
    "fmt"
    "projects/web/jwt"
)

func main() {
    j := &jwt.Jwt{}
    jwt.Sign = "ssf33"
    j.Test("sss")
    fmt.Println("fefe")
}
