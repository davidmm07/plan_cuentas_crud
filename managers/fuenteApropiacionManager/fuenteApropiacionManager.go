package fuenteapropiacionmanager

import (
	"log"

	"github.com/astaxie/beego/orm"
	"github.com/udistrital/plan_cuentas_crud/models"
)

// RegistrarMultipleManager registra multiples registros en la tabla fuente_financiamiento_apropiacion y devuelve un arreglo de
// enteros correspondientes a todos los id de los registros realizados. En caso de que algún registro falle, se hace un rollback a la bd
func RegistrarMultipleManager(fuentesApropiacion []*models.FuenteFinanciamientoApropiacion) (idRegistrados []int64) {
	o := orm.NewOrm()

	if err := o.Begin(); err == nil {
		for _, fuenteApropiacion := range fuentesApropiacion {
			id, err := o.Insert(fuenteApropiacion)
			if err != nil {
				o.Rollback()
				log.Panicln(err.Error())
			}
			idRegistrados = append(idRegistrados, id)
		}
	}

	o.Commit()
	return
}
