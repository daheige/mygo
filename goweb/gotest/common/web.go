package common

import (
	"encoding/json"
	"net/http"
)

func Routes() {
	http.HandleFunc("/sendjson", SendJson)
}

func SendJson(res http.ResponseWriter, req *http.Request) {
	u := struct {
		Name string `json:"name"`
	}{
		Name: "heige大黑哥",
	}

	res.Header().Set("Content-Type", "applicaion/json")
	res.WriteHeader(http.StatusOK)
	//创建json编码器，然后把u编码为json数据返回
	// json.NewEncoder(res).Encode(u)

	//另一种方式返回json
	json_data, _ := json.Marshal(&u)
	res.Write(json_data)
}
