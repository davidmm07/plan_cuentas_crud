package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaTipoFuenteFinanciamiento_20190628_120748 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaTipoFuenteFinanciamiento_20190628_120748{}
	m.Created = "20190628_120748"

	migration.Register("CrearTablaTipoFuenteFinanciamiento_20190628_120748", m)
	
}

// Run the migrations
func (m *CrearTablaTipoFuenteFinanciamiento_20190628_120748) Up() {
	m.SQL("CREATE TABLE plan_cuentas.tipo_fuente_financiamiento(id serial NOT NULL,nombre character varying(70) NOT NULL,descripcion character varying(300), CONSTRAINT pk_tipo_fuente_financiamiento PRIMARY KEY (id));") 
	m.SQL("ALTER TABLE plan_cuentas.tipo_fuente_financiamiento OWNER TO desarrollooas;")
}

// Reverse the migrations
func (m *CrearTablaTipoFuenteFinanciamiento_20190628_120748) Down() {
	m.SQL("DROP TABLE IF EXISTS plan_cuentas.tipo_fuente_financiamiento CASCADE;")
}
