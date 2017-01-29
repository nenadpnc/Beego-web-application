package controllers

import (
	"github.com/astaxie/beego"
)

type BankingCOntroller struct {
	beego.Controller
}

func (c *BankingCOntroller) URLMapping() {
	c.Mapping("ShowAccounts", c.ShowAccounts)
}

// @router /banking [get]
func (c *BankingCOntroller) ShowAccounts() {
	c.TplName = "banking.tpl"
}
