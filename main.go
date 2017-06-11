package main

import (
	"os"
	"strconv"

	_ "github.com/mark-greene/go-blackjack/routers"

	"github.com/astaxie/beego"
)

func main() {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err == nil {
		beego.BConfig.Listen.HTTPPort = port
	}
	beego.Run()
}
