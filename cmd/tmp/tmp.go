package main

import (
	"github.com/tetsuzawa/ddd_tutorial"
	"log"
)

func main() {
	userRepository := ddd_tutorial.NewUserRepository("root:password@tcp($(DOCKER_DNS):3306)/$(DBNAME)?parseTime=true&time_zone=%27Asia%2FTokyo%27&loc=Local")
	program := ddd_tutorial.NewProgram(userRepository)
	if err := program.CreateUser("tetsuzawa"); err != nil {
		log.Println(err)
	}
}
