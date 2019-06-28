package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaRubro_20190621_121535 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaRubro_20190621_121535{}
	m.Created = "20190621_121535"

	migration.Register("CrearTablaRubro_20190621_121535", m)
}

// Run the migrations
func (m *CrearTablaRubro_20190621_121535) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE plan_cuentas.rubro (id serial NOT NULL,organizacion integer NOT NULL,codigo varchar(50) NOT NULL,nombre varchar(250) NOT NULL,unidad_ejecutora integer NOT NULL,descripcion text,CONSTRAINT pk_rubro PRIMARY KEY (id),CONSTRAINT uq_organizacion_codigo_unidad_ejecutora_rubro UNIQUE (organizacion,unidad_ejecutora,codigo));")
	m.SQL("ALTER TABLE plan_cuentas.rubro ADD CONSTRAINT ck_codigo_rubro CHECK (((codigo)::text ~ '^([0-9]+-)+[0-9]+$'::text))")
	m.SQL("ALTER TABLE plan_cuentas.rubro OWNER TO desarrollooas;")

}

// Reverse the migrations
func (m *CrearTablaRubro_20190621_121535) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS plan_cuentas.rubro CASCADE;")
}
