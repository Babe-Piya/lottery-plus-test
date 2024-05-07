package main

import (
	"fmt"

	"lottery-plus/route"
)

func main() {
	app := route.Route()

	fmt.Println(app.Listen(":8080"))
}
