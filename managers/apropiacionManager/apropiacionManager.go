package apropiacionmanager

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func VigenciaApropiacion() (ml []int, err error) {
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("DISTINCT vigencia").
		From("" + beego.AppConfig.String("PGschemas") + ".apropiacion")

	sql := qb.String()
	o := orm.NewOrm()
	o.Raw(sql).QueryRows(&ml)

	if len(ml) == 0 {
		return nil, err
	}
	return ml, nil
}

//AprobarPresupuesto... Aprobacion de presupuesto (cambio de estado).
func AprobarPresupuesto(UnidadEjecutora int, Vigencia int) (err error) {
	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")
	qb2, _ := orm.NewQueryBuilder("mysql")
	qb2.Select("id").From(beego.AppConfig.String("PGschemas") + ".rubro").Where("unidad_ejecutora = ?")
	qb.Update(beego.AppConfig.String("PGschemas") + ".apropiacion").Set("estado = ?").Where("vigencia = ? AND rubro in (" + qb2.String() + ")")
	_, err = o.Raw(qb.String(), 2, Vigencia, UnidadEjecutora).Exec()
	return
}
