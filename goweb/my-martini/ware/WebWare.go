package ware

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

//模板渲染配置
func TplConfig(m *martini.ClassicMartini) {
	m.Use(render.Renderer(render.Options{
		Directory:  "views",
		Extensions: []string{".html", ".tmpl"},
	}))
}
