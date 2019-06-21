package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearEsquemaPlanCuentas_20190621_121431 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearEsquemaPlanCuentas_20190621_121431{}
	m.Created = "20190621_121431"

	migration.Register("CrearEsquemaPlanCuentas_20190621_121431", m)
}

// Run the migrations
func (m *CrearEsquemaPlanCuentas_20190621_121431) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE SCHEMA plan_cuentas;")

}

// Reverse the migrations
func (m *CrearEsquemaPlanCuentas_20190621_121431) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP SCHEMA IF EXISTS plan_cuentas CASCADE;")

}
