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

//DeleteRubroRelation Delete relation rama by id
func DeleteRubroRelation(id int) {
	o := orm.NewOrm()
	o.Begin()
	v := Rama{Id: id}
	_, err:= o.Read(&v)
	// ascertain id exists in the database
	if err != nil {
		o.Rollback()
		panic(appmessagemanager.DeleteErrorMessage())
	}
	_, err:= o.Delete(&Rama{Id: id})
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
// DeleteRubro deletes Rubro by Id and returns error if
// the record to be deleted doesn't exist
func DeleteRubro(id int) {
	o := orm.NewOrm()
	v := Rubro{Id: id}
	var apropiaciones []int
	var rama []int
	// ascertain id exists in the database
	o.Begin()
	_, err := o.Read(&v)
	if err != nil {
		qb, _ := orm.NewQueryBuilder("mysql")
		qb.Select("id").
			From("" + beego.AppConfig.String("PGschemas") + ".apropiacion").
			Where("rubro=?")
	}
	_, err := o.Raw(qb.String(), id).QueryRows(&apropiaciones)
	if err != nil {
		fmt.Println("Error consulta apropiacion por rubro")
		o.Rollback()
		panic(appmessagemanager.DeleteErrorMessage())
	}		
	if len(apropiaciones) == 0 {
		qb, _ = orm.NewQueryBuilder("mysql")
		qb.Select("id").
			From("" + beego.AppConfig.String("PGschemas") + ".rama").
			//Where("rubro_padre=?").
			Where("rubro_hijo=?")
		_, err:= o.Raw(qb.String(), id).QueryRows(&rama)
		if err != nil {
			fmt.Println("erro en tr")
			o.Rollback()
		}
		for _, idx := range rama {
			_, err := o.Delete(&Rama{Id: idx})
			if err != nil {
				fmt.Println("Error en tr ")
				o.Rollback()
				panic(appmessagemanager.DeleteErrorMessage())
			}
		}
	}
	num, err := o.Delete(&Rubro{Id: id})
	if err != nil {
		fmt.Println("Error 2 ", err)
		o.Rollback()
		panic(appmessagemanager.DeleteErrorMessage())
	} 
	fmt.Println("Number of records deleted in database:", num)
	o.Commit()
}
