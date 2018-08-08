package jwt

import (
    "fmt"
)

var Sign string

type Jwt struct {
}

func init() {
    Sign = "daheige"
}

func (j *Jwt) Test(str string) {
    fmt.Println(Sign)
    fmt.Println(str)
}
