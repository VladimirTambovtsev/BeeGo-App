package controllers

import "github.com/astaxie/beego"

type BooksController struct {
	beego.Controller
}

func (c *BooksController) URLMapping() {
	c.Mapping("ShowBooks", c.ShowBooks)
}

// @router /books [get]
func (c *BooksController) ShowBooks() {
	c.TplName = "books.tpl"
}
