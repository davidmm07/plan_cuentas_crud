// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/plan_cuentas_crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/rubro",
			beego.NSInclude(
				&controllers.RubroController{},
			),
		),
		beego.NSNamespace("/rama",
			beego.NSInclude(
				&controllers.RamaController{},
			),
		),
		beego.NSNamespace("/apropiacion",
			beego.NSInclude(
				&controllers.ApropiacionController{},
			),
		),
		beego.NSNamespace("/estado_apropiacion",
			beego.NSInclude(
				&controllers.EstadoApropiacionController{},
			),
		),
		beego.NSNamespace("/date",
			beego.NSInclude(
				&controllers.Date{},
			),
		),

		beego.NSNamespace("/tipo_fuente_financiamiento",
			beego.NSInclude(
				&controllers.TipoFuenteFinanciamientoController{},
			),
		),

		beego.NSNamespace("/fuente_financiamiento",
			beego.NSInclude(
				&controllers.FuenteFinanciamientoController{},
			),
		),

		beego.NSNamespace("/fuente_financiamiento_apropiacion",
			beego.NSInclude(
				&controllers.FuenteFinanciamientoApropiacionController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
