package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaFuenteFinanciamientoApropiacion_20190628_123019 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaFuenteFinanciamientoApropiacion_20190628_123019{}
	m.Created = "20190628_123019"

	migration.Register("CrearTablaFuenteFinanciamientoApropiacion_20190628_123019", m)
}

// Run the migrations
func (m *CrearTablaFuenteFinanciamientoApropiacion_20190628_123019) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE plan_cuentas.fuente_financiamiento_apropiacion(id serial NOT NULL,apropiacion_id integer NOT NULL,fuente_financiamiento_id integer NOT NULL,dependencia integer NOT NULL,CONSTRAINT pk_fuente_financiamiento_apropiacion PRIMARY KEY (id));")
	m.SQL("ALTER TABLE plan_cuentas.fuente_financiamiento_apropiacion ADD CONSTRAINT fk_fuente_financiamiento_apropiacion_apropiacion FOREIGN KEY (apropiacion_id) REFERENCES plan_cuentas.apropiacion (id) MATCH FULL ON DELETE NO ACTION ON UPDATE NO ACTION;")
	m.SQL("ALTER TABLE plan_cuentas.fuente_financiamiento_apropiacion ADD CONSTRAINT fk_fuente_financiamiento_apropiacion_fuente_financiamiento FOREIGN KEY (fuente_financiamiento_id) REFERENCES plan_cuentas.fuente_financiamiento (id) MATCH FULL ON DELETE NO ACTION ON UPDATE NO ACTION;")
	m.SQL("ALTER TABLE plan_cuentas.fuente_financiamiento_apropiacion OWNER TO desarrollooas;")
}

// Reverse the migrations
func (m *CrearTablaFuenteFinanciamientoApropiacion_20190628_123019) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS plan_cuentas.fuente_financiamiento_apropiacion CASCADE;")

}
