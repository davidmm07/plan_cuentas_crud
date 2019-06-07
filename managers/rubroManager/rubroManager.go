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

//DeleteRubroRelation Delete relation rubro_rubro by id
func DeleteRubroRelation(id int) {
	o := orm.NewOrm()
	o.Begin()
	v := RubroRubro{Id: id}
	_, err:= o.Read(&v)
	// ascertain id exists in the database
	if err != nil {
		o.Rollback()
		panic(appmessagemanager.DeleteErrorMessage())
	}
	_, err:= o.Delete(&RubroRubro{Id: id})
	if err != nil {
		o.Rollback()
		panic(appmessagemanager.DeleteErrorMessage())
	} 
	_, err =: o.Delete(v.RubroHijo)
	if err != nil {
		o.Rollback()
		panic(appmessagemanager.DeleteErrorMessage())
	}
	o.Commit()

	return
}