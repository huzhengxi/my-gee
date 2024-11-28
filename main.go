package main

import (
	"fmt"
	"gee"
	"html/template"
	"log"
	"net/http"
	"time"
)

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := gee.Default()
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHtmlGlob("templates/*")
	r.Static("/assets", "./static")

	stu1 := &student{Name: "Geekytu", Age: 20}
	stu2 := &student{Name: "jack", Age: 22}
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})

	r.GET("/students", func(c *gee.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", gee.H{
			"title":  "gee",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	r.GET("/date", func(c *gee.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", gee.H{
			"title": "gee",
			"now":   time.Date(2020, 11, 10, 0, 0, 0, 0, time.UTC),
		})
	})

	r.GET("/panic", func(c *gee.Context) {
		name := []string{"geektutu"}
		c.String(http.StatusOK, name[100])
	})

	v1 := r.Group("/v1")
	{

		v1.GET("/hello", func(c *gee.Context) {
			c.String(http.StatusOK, "hello %s", c.Query("name"))
		})
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *gee.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})

		v2.POST("/login", func(c *gee.Context) {
			c.JSON(http.StatusOK, gee.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})

		})
	}

	log.Println("server run :9999")
	r.Run(":9999")
}
