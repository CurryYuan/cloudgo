package service

import (
	"time"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/binding"
)

type Post struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func NewServer(port string) { 
	// default setting
	m := martini.Classic()
	// use dir public
	m.Use(martini.Static("public"))
	m.Use(render.Renderer())
	// routing
	m.Get("/", func(r render.Render) {
		t1 := time.Now().UTC().Format(time.UnixDate)
		newmap := map[string]interface{}{"Name": "Zhi Hao","formatTime":t1}

        r.HTML(200, "index", newmap)
	})

	m.Post("/", binding.Bind(Post{}), func(post Post, r render.Render) {
		p:=Post{Username: post.Username, Password: post.Password}
		newmap := map[string]interface{}{"metatitle": "created post", "post": p}
        r.HTML(200, "form", newmap)
	})
	//run
	m.RunOnAddr(":"+port)
}

// template: form:3:18: executing "form" at <.post.username>: username is an unexported field of struct type interface {}
