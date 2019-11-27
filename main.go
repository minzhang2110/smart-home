package main

import (
	"log"
	"os"

	"github.com/minzhang2110/smart-home/pkg/controllers"
)

func main() {
	h := controllers.NewHandler()
	log.Fatal((controllers.NewRouter(h).Start(":" + os.Getenv("PORT"))))
}
