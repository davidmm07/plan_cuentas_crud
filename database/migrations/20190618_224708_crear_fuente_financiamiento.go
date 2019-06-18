package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearFuenteFinanciamiento_20190618_224708 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearFuenteFinanciamiento_20190618_224708{}
	m.Created = "20190618_224708"

	migration.Register("CrearFuenteFinanciamiento_20190618_224708", m)
}

// Run the migrations
func (m *CrearFuenteFinanciamiento_20190618_224708) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE tipo_fuente_financiamiento (id serial NOT NULL,	nombre varchar(20) NOT NULL,		descripcion varchar(250) NULL,codigo_abreviacion varchar(20) NULL,activo bool NOT NULL, numero_orden numeric(5,2) NULL,CONSTRAINT pk_tipo_fuente_financiamiento PRIMARY KEY (id));")
	m.SQL("CREATE TABLE fuente_financiamiento (id serial NOT NULL,descripcion varchar(250) NULL,nombre varchar(250) NOT NULL,codigo varchar(6) NOT NULL,tipo_fuente_financiamiento_id int4 NULL,CONSTRAINT pk_fuente_financiamiento PRIMARY KEY (id));")
	m.SQL("CREATE TABLE fuente_financiamiento_apropiacion (id serial NOT NULL,apropiacion_id int4 NOT NULL,fuente_financiamiento_id int4 NOT NULL,dependencia int4 NOT NULL,CONSTRAINT pk_fuente_financiamiento_apropiacion PRIMARY KEY (id));")
	m.SQL("ALTER TABLE fuente_financiamiento ADD CONSTRAINT fk_fuente_financiamiento_tipo_fuente_financiamiento FOREIGN KEY (tipo_fuente_financiamiento_id) REFERENCES tipo_fuente_financiamiento(id) ON UPDATE RESTRICT ON DELETE RESTRICT;")
	m.SQL("ALTER TABLE fuente_financiamiento_apropiacion ADD CONSTRAINT fk_fuente_financiamiento_apropiacion_apropiacion FOREIGN KEY (apropiacion_id) REFERENCES apropiacion(id) ON UPDATE RESTRICT ON DELETE RESTRICT;")
	m.SQL("ALTER TABLE fuente_financiamiento_apropiacion ADD CONSTRAINT fk_fuente_financiamiento_apropiacion_fuente_financiamiento FOREIGN KEY (fuente_financiamiento_id) REFERENCES fuente_financiamiento(id) ON UPDATE RESTRICT ON DELETE RESTRICT;")
}

// Reverse the migrations
func (m *CrearFuenteFinanciamiento_20190618_224708) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
