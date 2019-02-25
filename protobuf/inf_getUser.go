package main

import (
	"fmt"
	"mygo/protobuf/library"
)

func main() {
	greeterService := library.Greeter{
		Address: "localhost:50051",
	}
	for i := 0; i < 10; i++ {
		userInfo := greeterService.GetUserInfo(1, 2)
		fmt.Println("reply data: ", userInfo)
	}
}

/**
 * go run inf_getUser.go
 * reply data:  uid:1 job:"golang" name:"daheige" age:28
reply data:  uid:1 job:"golang" name:"daheige" age:28
reply data:  uid:1 job:"golang" name:"daheige" age:28
reply data:  uid:1 job:"golang" name:"daheige" age:28
reply data:  uid:1 job:"golang" name:"daheige" age:28
reply data:  uid:1 job:"golang" name:"daheige" age:28
reply data:  uid:1 job:"golang" name:"daheige" age:28
reply data:  uid:1 job:"golang" name:"daheige" age:28
reply data:  uid:1 job:"golang" name:"daheige" age:28
reply data:  uid:1 job:"golang" name:"daheige" age:28
*/
