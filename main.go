package main

import (
	"github.com/minzhang2110/smart-home/pkg/controllers"
	"log"
)

func main() {
	h := controllers.NewHandler()
	log.Error(controllers.New(h).Start(":" + os.Getenv("PORT")))
}
