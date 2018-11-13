package service

import "github.com/go-martini/martini"

func NewServer(port string) { 
	// default setting
	m := martini.Classic()
	// use dir public
	m.Use(martini.Static("public"))
	// routing
	m.Get("/", func() string {
		return "index.html"
	})
	//run
	m.RunOnAddr(":"+port)
}