package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type GenerarTablasPlanCuentasConDoc_20190625_124322 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &GenerarTablasPlanCuentasConDoc_20190625_124322{}
	m.Created = "20190625_124322"

	migration.Register("GenerarTablasPlanCuentasConDoc_20190625_124322", m)
}

// Run the migrations
func (m *GenerarTablasPlanCuentasConDoc_20190625_124322) Up() {
	file, err := ioutil.ReadFile("../files/generar_tablas_plan_cuentas_con_doc.up.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";\n")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}

}

// Reverse the migrations
func (m *GenerarTablasPlanCuentasConDoc_20190625_124322) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
