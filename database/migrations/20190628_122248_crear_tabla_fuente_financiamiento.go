package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaFuenteFinanciamiento_20190628_122248 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaFuenteFinanciamiento_20190628_122248{}
	m.Created = "20190628_122248"

	migration.Register("CrearTablaFuenteFinanciamiento_20190628_122248", m)
}

// Run the migrations
func (m *CrearTablaFuenteFinanciamiento_20190628_122248) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE plan_cuentas.fuente_financiamiento(id serial NOT NULL,nombre character varying(70) NOT NULL,descripcion character varying(200),fecha_creacion date,tipo_fuente_financiamiento integer NOT NULL,codigo character varying(10) NOT NULL,CONSTRAINT pk_fuente_financiamiento PRIMARY KEY (id));")
	m.SQL("ALTER TABLE plan_cuentas.fuente_financiamiento ADD CONSTRAINT fk_fuente_financiamiento_tipo_fuente_financiamiento FOREIGN KEY (tipo_fuente_financiamiento) REFERENCES plan_cuentas.tipo_fuente_financiamiento (id) MATCH FULL ON DELETE NO ACTION ON UPDATE NO ACTION;")
	m.SQL("ALTER TABLE plan_cuentas.fuente_financiamiento OWNER TO desarrollooas;")
}

// Reverse the migrations
func (m *CrearTablaFuenteFinanciamiento_20190628_122248) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS plan_cuentas.fuente_financiamiento CASCADE;")

}
