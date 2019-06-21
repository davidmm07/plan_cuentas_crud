package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaEstadoApropiacion_20190621_121633 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaEstadoApropiacion_20190621_121633{}
	m.Created = "20190621_121633"

	migration.Register("CrearTablaEstadoApropiacion_20190621_121633", m)
}

// Run the migrations
func (m *CrearTablaEstadoApropiacion_20190621_121633) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE plan_cuentas.estado_apropiacion (id serial NOT NULL,nombre varchar NOT NULL,descripcion varchar NOT NULL,codigo_abreviacion varchar NOT NULL,numero_orden integer,activo bool NOT NULL,CONSTRAINT pk_estado_apropiacion PRIMARY KEY (id));")

}

// Reverse the migrations
func (m *CrearTablaEstadoApropiacion_20190621_121633) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS plan_cuentas.estado_apropiacion CASCADE;")

}
