package controllers

import (
	"encoding/json"
	"fmt"
	"hgaction/data"
	"log"
	"net/http"
)

func Handler(res http.ResponseWriter, req *http.Request) {
	user := &data.UserInfo{
		Id:   1,
		Name: "heige",
	}
	s, err := json.Marshal(user)
	fmt.Println(s)
	if err != nil {
		log.Fatalf("error is %s", err)
		fmt.Fprintf(res, "err request!")
		return
	}

	fmt.Fprintf(res, "res is %s", string(s))
}
