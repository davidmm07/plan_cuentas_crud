package rubromanager

import (
	"github.com/astaxie/beego/orm"
	appmessagemanager "github.com/udistrital/plan_cuentas_crud/managers/appMessageManager"
	"github.com/udistrital/plan_cuentas_crud/models"
)

func RubroRelationRegistrator(idParent int, Rubro *models.Rubro) {
	o := orm.NewOrm()
	o.Begin()
	_, err := o.Insert(Rubro)
	if err != nil {
		o.Rollback()
		panic(appmessagemanager.InsertErrorMessage())
	}
	relation := models.Rama{}
	relation.RubroHijo = Rubro
	relation.RubroPadre = &models.Rubro{}
	relation.RubroPadre.Id = idParent
	_, err = o.Insert(&relation)
	if err != nil {
		o.Rollback()
		panic(appmessagemanager.InsertErrorMessage())
	}
	o.Commit()

}
