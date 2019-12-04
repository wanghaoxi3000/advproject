package main

import "advancedproject/server"

func main() {
	r := server.SetupRouter()
	r.Run(":3000")
}
