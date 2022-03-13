package main

import "bookstore/src"

func main() {
	app := src.App{}
	app.Configure()
	app.Run()
}
