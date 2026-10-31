package main

import "github.com/PedroBett-dev/Golang-Api/src/api"

func main() {
	if err := api.Run(); err != nil {
		panic(err)
	}
}
