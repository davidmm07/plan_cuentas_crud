package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	appmessagemanager "github.com/udistrital/plan_cuentas_crud/managers/appMessageManager"
	rubromanager "github.com/udistrital/plan_cuentas_crud/managers/rubroManager"
	"github.com/udistrital/plan_cuentas_crud/models"

	"github.com/astaxie/beego"
)

// RamaController operations for Rama
type RamaController struct {
	beego.Controller
}

// URLMapping ...
func (c *RamaController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Rama
//@Param	parentId	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	body		body 	models.Rama	true		"body for Rama content"
// @Success 201 {int} models.Rama
// @Failure 403 body is empty
// @router / [post]
func (c *RamaController) Post() {
	var v models.Rama
	parentID := 0
	var err error

	if parentIdSTR := c.GetString("parentId"); parentIdSTR != "" {
		parentID, err = strconv.Atoi(parentIdSTR)
		if err != nil {
			beego.Error(err.Error())
			panic(appmessagemanager.ParamsErrorMessage())
		}
	}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if parentID == 0 {
			if _, err := models.AddRama(&v); err == nil {
				c.Data["json"] = v
			} else {
				c.Data["json"] = err
			}
		} else {
			rubromanager.RubroRelationRegistrator(parentID, &v)
			c.Data["json"] = v
		}

	} else {
		c.Data["json"] = err
	}
}

// GetOne ...
// @Title Get One
// @Description get Rama by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Rama
// @Failure 403 :id is empty
// @router /:id [get]
func (c *RamaController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetRamaById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
}

// GetAll ...
// @Title Get All
// @Description get Rama
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Rama
// @Failure 403
// @router / [get]
func (c *RamaController) GetAll() {
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
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllRama(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
}

// Put ...
// @Title Put
// @Description update the Rama
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Rama	true		"body for Rama content"
// @Success 200 {object} models.Rama
// @Failure 403 :id is not int
// @router /:id [put]
func (c *RamaController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Rama{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateRamaById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
}

// Delete ...
// @Title Delete
// @Description delete the Rama
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *RamaController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteRama(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
}

// DeleteRubroRelation ...
// @Title DeleteRubroRelation
// @Description delete the Rama
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /DeleteRubroRelation/:id/:ue [delete]
func (c *RamaController) DeleteRubroRelation() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	ueStr := c.Ctx.Input.Param(":ue")
	ue, _ := strconv.Atoi(ueStr)
	if err := models.DeleteRubroRelation(id); err == nil {
		go genRubrosTreeFile(int(ue))
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
}
