package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaApropiacion_20190621_125201 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaApropiacion_20190621_125201{}
	m.Created = "20190621_125201"

	migration.Register("CrearTablaApropiacion_20190621_125201", m)
}

// Run the migrations
func (m *CrearTablaApropiacion_20190621_125201) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE plan_cuentas.apropiacion (id serial NOT NULL,vigencia integer NOT NULL,valor numeric(20,7) NOT NULL,estado_apropiacion_id integer NOT NULL,rubro_id integer,CONSTRAINT pk_apropiacion PRIMARY KEY (id));")
	m.SQL("ALTER TABLE plan_cuentas.apropiacion ADD CONSTRAINT fk_apropiacion_rubro FOREIGN KEY (rubro_id) REFERENCES plan_cuentas.rubro(id) MATCH FULL;")
	m.SQL("ALTER TABLE plan_cuentas.apropiacion ADD CONSTRAINT fk_apropiacion_estado_apropiacion FOREIGN KEY (estado_apropiacion_id) REFERENCES plan_cuentas.estado_apropiacion(id) ON UPDATE CASCADE ON DELETE RESTRICT;")
	m.SQL("ALTER TABLE plan_cuentas.apropiacion OWNER TO desarrollooas;")
}

// Reverse the migrations
func (m *CrearTablaApropiacion_20190621_125201) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("ALTER TABLE plan_cuentas.apropiacion DROP CONSTRAINT fk_apropiacion_rubro ;")
	m.SQL("ALTER TABLE plan_cuentas.apropiacion DROP CONSTRAINT fk_apropiacion_estado_apropiacion ;")
	m.SQL("DROP TABLE IF EXISTS plan_cuentas.apropiacion CASCADE;")

}
