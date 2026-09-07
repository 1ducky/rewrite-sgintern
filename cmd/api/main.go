package main

import (
	"RewriteProject/internal/utils"
	"log"
)

func main() {
	var pass string = "asdasdasd"
	hashedPass, _ := utils.HashString(pass)
	log.Println(hashedPass)
	log.Println(utils.CheckHashString(hashedPass, pass))
}
