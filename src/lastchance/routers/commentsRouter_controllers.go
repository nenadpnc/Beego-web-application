package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["lastchance/controllers:BankingCOntroller"] = append(beego.GlobalControllerRouter["lastchance/controllers:BankingCOntroller"],
		beego.ControllerComments{
			Method: "ShowAccounts",
			Router: `/banking`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
