package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaRama_20190621_121547 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaRama_20190621_121547{}
	m.Created = "20190621_121547"

	migration.Register("CrearTablaRama_20190621_121547", m)
}

// Run the migrations
func (m *CrearTablaRama_20190621_121547) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE plan_cuentas.rama (id serial NOT NULL,rubro_padre integer NOT NULL,rubro_hijo integer NOT NULL,CONSTRAINT pk_rama PRIMARY KEY (id));")
	m.SQL("ALTER TABLE plan_cuentas.rama ADD CONSTRAINT fk_rubro_hijo FOREIGN KEY (rubro_hijo) REFERENCES plan_cuentas.rubro(id) MATCH FULL;")
	m.SQL("ALTER TABLE plan_cuentas.rama ADD CONSTRAINT fk_rubro_padre FOREIGN KEY (rubro_padre) REFERENCES plan_cuentas.rubro(id) MATCH FULL;")
	m.SQL("ALTER TABLE plan_cuentas.rama OWNER TO desarrollooas;")

}

// Reverse the migrations
func (m *CrearTablaRama_20190621_121547) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("ALTER TABLE plan_cuentas.rama DROP CONSTRAINT fk_rubro_hijo FOREIGN KEY ;")
	m.SQL("ALTER TABLE plan_cuentas.rama DROP CONSTRAINT fk_rubro_padre FOREIGN KEY ;")
	m.SQL("DROP TABLE IF EXISTS plan_cuentas.rama CASCADE;")

}
