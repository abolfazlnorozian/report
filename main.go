package main

import (
	"report/app"
	"report/helper/env"
)

func init() {
	env.Load()
}
func main() {
	app.RunServer()

}
