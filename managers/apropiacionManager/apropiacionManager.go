package apropiacionmanager

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	appmessagemanager "github.com/udistrital/plan_cuentas_crud/managers/appMessageManager"
)

// funcion que lista las distintas vigencias para las que existen apropiaciones
func VigenciaApropiacion() (listavigencias []int, err error) {
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("DISTINCT vigencia").
		From("" + beego.AppConfig.String("PGschemas") + ".apropiacion")

	sql := qb.String()
	o := orm.NewOrm()
	o.Raw(sql).QueryRows(&listavigencias)

	if len(listavigencias) == 0 {
		return nil, err
	}
	return listavigencias, nil
}

//AprobarPresupuesto... Aprobacion de presupuesto (cambio de estado).
func AprobarPresupuesto(UnidadEjecutora int, Vigencia int) {
	o := orm.NewOrm()
	o.Begin()
	qb, _ := orm.NewQueryBuilder("mysql")
	qb2, _ := orm.NewQueryBuilder("mysql")
	qb2.Select("id").From(beego.AppConfig.String("PGschemas") + ".rubro").Where("unidad_ejecutora = ?")
	qb.Update(beego.AppConfig.String("PGschemas") + ".apropiacion").Set("estado = ?").Where("vigencia = ? AND rubro in (" + qb2.String() + ")")
	_, err := o.Raw(qb.String(), 2, Vigencia, UnidadEjecutora).Exec()
	if err != nil {
		o.Rollback()
		panic(appmessagemanager.AprobarPresupuestoErrorMessage())
	}
	o.Commit()
}
