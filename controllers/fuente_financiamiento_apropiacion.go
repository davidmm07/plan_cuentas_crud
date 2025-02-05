package controllers

import (
	"encoding/json"
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	fuenteApropiacionManager "github.com/udistrital/plan_cuentas_crud/managers/fuenteApropiacionManager"
	"github.com/udistrital/plan_cuentas_crud/models"
	utilidades "github.com/udistrital/plan_cuentas_crud/utilidades"
	"github.com/udistrital/utils_oas/responseformat"
)

// FuenteFinanciamientoApropiacionController operations for FuenteFinanciamientoApropiacion
type FuenteFinanciamientoApropiacionController struct {
	beego.Controller
}

// URLMapping ...
func (c *FuenteFinanciamientoApropiacionController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("RegistrarMultiple", c.RegistrarMultiple)
}

// RegistrarMultiple ...
// @Title RegistrarMultiple
// @Description Crea una nueva fuente de financiamiento con la relación de sus rubros y sus dependencias a los rubros
// @Param	body		body 	models.FuenteFinanciamientoApropiacion	true		"body for FuenteFinanciamientoApropiacion content"
// @Success 201 {int} int[]
// @Failure 403 body is empty
// @router /registrar_multiple [post]
func (c *FuenteFinanciamientoApropiacionController) RegistrarMultiple() {
	var v []*models.FuenteFinanciamientoApropiacion

	defer utilidades.ErrorResponse(c.Controller)

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		log.Panicln(err.Error())
	}

	ids := fuenteApropiacionManager.RegistrarMultipleManager(v)

	response := make(map[string]interface{})
	response["Ids"] = ids
	responseformat.SetResponseFormat(&c.Controller, response, "", 200)
}

// Post ...
// @Title Post
// @Description create FuenteFinanciamientoApropiacion
// @Param	body		body 	models.FuenteFinanciamientoApropiacion	true		"body for FuenteFinanciamientoApropiacion content"
// @Success 201 {int} models.FuenteFinanciamientoApropiacion
// @Failure 403 body is empty
// @router / [post]
func (c *FuenteFinanciamientoApropiacionController) Post() {
	var v models.FuenteFinanciamientoApropiacion
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddFuenteFinanciamientoApropiacion(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
			panic(err.Error())
		}
	} else {
		c.Data["json"] = err.Error()
		panic(err.Error())
	}

}

// GetOne ...
// @Title Get One
// @Description get FuenteFinanciamientoApropiacion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.FuenteFinanciamientoApropiacion
// @Failure 403 :id is empty
// @router /:id [get]
func (c *FuenteFinanciamientoApropiacionController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetFuenteFinanciamientoApropiacionById(id)
	if err != nil {
		c.Data["json"] = err.Error()
		panic(err.Error())
	} else {
		c.Data["json"] = v
	}

}

// GetAll ...
// @Title Get All
// @Description get FuenteFinanciamientoApropiacion
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.FuenteFinanciamientoApropiacion
// @Failure 403
// @router / [get]
func (c *FuenteFinanciamientoApropiacionController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")

				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllFuenteFinanciamientoApropiacion(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
		panic(err.Error())
	} else {
		c.Data["json"] = l
	}

}

// Put ...
// @Title Put
// @Description update the FuenteFinanciamientoApropiacion
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.FuenteFinanciamientoApropiacion	true		"body for FuenteFinanciamientoApropiacion content"
// @Success 200 {object} models.FuenteFinanciamientoApropiacion
// @Failure 403 :id is not int
// @router /:id [put]
func (c *FuenteFinanciamientoApropiacionController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.FuenteFinanciamientoApropiacion{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateFuenteFinanciamientoApropiacionById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
			panic(err.Error())
		}
	} else {
		c.Data["json"] = err.Error()
		panic(err.Error())
	}

}

// Delete ...
// @Title Delete
// @Description delete the FuenteFinanciamientoApropiacion
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *FuenteFinanciamientoApropiacionController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteFuenteFinanciamientoApropiacion(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
		panic(err.Error())
	}

}
