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

// DeleteRubro deletes Rubro by Id and returns error if
// the record to be deleted doesn't exist
func DeleteRubro(id int) (err error) {
	o := orm.NewOrm()
	v := Rubro{Id: id}
	var apropiaciones []int
	var rubrorubro []int
	// ascertain id exists in the database
	o.Begin()
	if err = o.Read(&v); err == nil {
		var num int64
		qb, _ := orm.NewQueryBuilder("mysql")
		qb.Select("id").
			From("" + beego.AppConfig.String("PGschemas") + ".apropiacion").
			Where("rubro=?")
		if _, err = o.Raw(qb.String(), id).QueryRows(&apropiaciones); err == nil {

			if len(apropiaciones) == 0 {
				qb, _ = orm.NewQueryBuilder("mysql")
				qb.Select("id").
					From("" + beego.AppConfig.String("PGschemas") + ".rubro_rubro").
					//Where("rubro_padre=?").
					Where("rubro_hijo=?")
				if _, err = o.Raw(qb.String(), id).QueryRows(&rubrorubro); err == nil {
					for _, idx := range rubrorubro {
						if _, err = o.Delete(&RubroRubro{Id: idx}); err == nil {

						} else {
							o.Rollback()
							err = errors.New("erro en tr")
							return
						}
					}
				}
			} else {
				o.Rollback()
				err = errors.New("erro en tr")
				return
			}

		} else {
			fmt.Println("Error 1 ", err)
			o.Rollback()
			return
		}
		if num, err = o.Delete(&Rubro{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
			o.Commit()
		} else {
			fmt.Println("Error 2 ", err)

			o.Rollback()
			return
		}
	}
	return
}
