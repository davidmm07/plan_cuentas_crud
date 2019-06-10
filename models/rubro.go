package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
	"github.com/udistrital/utils_oas/formatdata"
	"github.com/udistrital/utils_oas/optimize"
)

type Rubro struct {
	Id              int    `orm:"auto;column(id);pk"`
	Organizacion    int    `orm:"column(organizacion)"`
	Codigo          string `orm:"column(codigo)"`
	Descripcion     string `orm:"column(descripcion);null"`
	UnidadEjecutora int
	Nombre          string `orm:"column(nombre);null"`
}

func (t *Rubro) TableName() string {
	return "rubro"
}

func init() {
	orm.RegisterModel(new(Rubro))
}

// AddRubro insert a new Rubro into database and returns
// last inserted Id on success.
func AddRubro(m *Rubro) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetRubroById retrieves Rubro by Id. Returns error if
// Id doesn't exist
func GetRubroById(id int) (v *Rubro, err error) {
	o := orm.NewOrm()
	v = &Rubro{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllRubro retrieves all Rubro matches certain condition. Returns empty list if
// no records exist
func GetAllRubro(query map[string]string, group []string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Rubro))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}
	var l []Rubro

	qs = qs.OrderBy(sortFields...).RelatedSel(5)

	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				//o.LoadRelated(&v, "ProductoRubro", 5, 0, 0, "-Activo")
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateRubro updates Rubro by Id and returns error if
// the record to be updated doesn't exist
func UpdateRubroById(m *Rubro) (err error) {
	o := orm.NewOrm()
	v := Rubro{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteRubro deletes Rubro by Id and returns error if
// the record to be deleted doesn't exist
func DeleteRubro(id int) (err error) {
	o := orm.NewOrm()
	v := Rubro{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Rubro{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// Generar arbol de rubros.
func ArbolRubrosMigracion() (rubros []map[string]interface{}, err error) {
	o := orm.NewOrm()
	var m []orm.Params
	//funcion para conseguir los rubros padre. OR (id not in (select DISTINCT rubro_padre from financiera.rubro_rubro))
	_, err = o.Raw(`
	SELECT DISTINCT ON (rubro.codigo) rubro.id as "idpsql", rubro.codigo as "_id",rubro.nombre as "nombre" , rubro.descripcion as "descripcion", rubro.unidad_ejecutora as "unidad_ejecutora"
	FROM public.rubro `).Values(&m)

	if err == nil {
		var res []interface{}
		err = formatdata.FillStruct(m, &res)
		done := make(chan interface{})
		defer close(done)
		resch := optimize.GenChanInterface(res...)

		charbolrubros := optimize.Digest(done, RamaRubrosMigracion, resch, nil)
		for data := range charbolrubros {
			if data != nil {
				rubros = append(rubros, data.(map[string]interface{})) //tomar valores del canal y agregarlos al array de hijos.
			}
		}
	} else {
		fmt.Println("error en ArbolRubrosMigracion: ", err.Error())
	}
	return
}

func RamaRubrosMigracion(forkin interface{}, params ...interface{}) (forkout interface{}) {
	fork := forkin.(map[string]interface{})
	fmt.Println(fork)
	o := orm.NewOrm()
	var m, n []orm.Params
	var res, res2 []interface{}
	//funcion para conseguir los hijos de los rubros padre.
	_, err := o.Raw(`SELECT rubro.codigo as "Codigo"
	  from public.rubro
	  join public.rama
		on  rama.rubro_hijo = rubro.id
	  WHERE rama.rubro_padre = ?`, fork["idpsql"]).Values(&m)
	// fmt.Println(m)
	if err == nil {
		var arr []interface{}
		var x map[string]interface{}
		for i := 0; i < len(m); i++ {
			if err = formatdata.FillStruct(m[i], &x); err == nil && x != nil {
				arr = append(arr, x["Codigo"])
			}
		}

		err = formatdata.FillStruct(arr, &res)
		fork["hijos"] = arr

		var arr2 []interface{}
		var y map[string]interface{}
		_, err2 := o.Raw(`SELECT rubro.codigo as "Codigo"
		from public.rubro
		join public.rama
		on  rama.rubro_padre = rubro.id
		WHERE rama.rubro_hijo = ?`, fork["idpsql"]).Values(&n)

		// fmt.Println("error en modelos : ", err2.Error())
		if err2 == nil {
			for i := 0; i < len(n); i++ {
				if err2 = formatdata.FillStruct(n[i], &y); err2 == nil && y != nil {
					arr2 = append(arr2, y["Codigo"])
				}
			}
			if len(arr2) != 0 {
				err2 = formatdata.FillStruct(arr2, &res2)
				fork["padre"] = arr2[0]
			} else {
				println("arreglo de padres vacio...")
			}
		} else {
			fmt.Println("error 2 en RamaRubrosMigracion: ", err2.Error())
		}

		return fork
	} else {
		fmt.Println("error 1 en RamaRubrosMigracion: ", err.Error())
	}
	return
}
