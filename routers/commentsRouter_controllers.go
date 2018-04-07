package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["beegoApp/controllers:BooksController"] = append(beego.GlobalControllerRouter["beegoApp/controllers:BooksController"],
		beego.ControllerComments{
			Method: "ShowBooks",
			Router: `/books`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
