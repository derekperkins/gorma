package gorma_test

import (
	"fmt"
	"testing"

	"github.com/goadesign/gorma"
	"github.com/goadesign/gorma/dsl"
)

func TestModelContext(t *testing.T) {
	sg := &gorma.RelationalModelDefinition{}
	sg.Name = "SG"

	c := sg.Context()
	exp := fmt.Sprintf("RelationalModel %#v", sg.Name)
	if c != exp {
		t.Errorf("Expected %s, got %s", exp, c)
	}

	sg.Name = ""

	c = sg.Context()
	exp = "unnamed RelationalModel"
	if c != exp {
		t.Errorf("Expected %s, got %s", exp, c)
	}
}

func TestModelDSL(t *testing.T) {
	sg := &gorma.RelationalModelDefinition{}
	f := func() {
		return
	}
	sg.DefinitionDSL = f
	c := sg.DSL()
	if c == nil {
		t.Errorf("Expected %s, got nil", f)
	}

}

func TestPKAttributesSingle(t *testing.T) {
	sg := &gorma.RelationalModelDefinition{}
	sg.RelationalFields = make(map[string]*gorma.RelationalFieldDefinition)
	f := makePK("id")
	sg.PrimaryKeys = append(sg.PrimaryKeys, f)
	sg.RelationalFields[f.Name] = f

	pka := sg.PKAttributes()

	if pka != "id int" {
		t.Errorf("Expected %s, got %s", "id int", pka)
	}

}
func TestPKAttributesMultiple(t *testing.T) {
	sg := &gorma.RelationalModelDefinition{}
	sg.RelationalFields = make(map[string]*gorma.RelationalFieldDefinition)
	f := makePK("Field1")
	sg.PrimaryKeys = append(sg.PrimaryKeys, f)
	sg.RelationalFields[f.Name] = f

	f2 := makePK("Field2")
	sg.RelationalFields[f2.Name] = f2
	sg.PrimaryKeys = append(sg.PrimaryKeys, f2)

	pka := sg.PKAttributes()

	if pka != "field1 int,field2 int" {
		t.Errorf("Expected %s, got %s", "field1 int,field2 int", pka)
	}

}
func makePK(name string) *gorma.RelationalFieldDefinition {

	f := &gorma.RelationalFieldDefinition{}
	f.Name = dsl.SanitizeFieldName(name)
	f.DatabaseFieldName = dsl.SanitizeDBFieldName(f.Name)
	f.Datatype = gorma.PKInteger
	return f

}
func TestPKWhereSingle(t *testing.T) {
	sg := &gorma.RelationalModelDefinition{}
	sg.RelationalFields = make(map[string]*gorma.RelationalFieldDefinition)
	f := &gorma.RelationalFieldDefinition{}
	f.Name = dsl.SanitizeFieldName("ID")
	f.DatabaseFieldName = dsl.SanitizeDBFieldName(f.Name)
	f.Datatype = gorma.PKInteger

	sg.RelationalFields[f.Name] = f
	sg.PrimaryKeys = append(sg.PrimaryKeys, f)

	pkw := sg.PKWhere()

	if pkw != "id = ?" {
		t.Errorf("Expected %s, got %s", "id = ?", pkw)
	}

}

func TestPKWhereMultiple(t *testing.T) {
	sg := &gorma.RelationalModelDefinition{}
	sg.RelationalFields = make(map[string]*gorma.RelationalFieldDefinition)
	f := makePK("Field1")
	sg.RelationalFields[f.Name] = f
	sg.PrimaryKeys = append(sg.PrimaryKeys, f)

	f2 := makePK("Field2")
	sg.RelationalFields[f2.Name] = f2
	sg.PrimaryKeys = append(sg.PrimaryKeys, f2)

	pkw := sg.PKWhere()

	if pkw != "field1 = ? and field2 = ?" {
		t.Errorf("Expected %s, got %s", "field1 = ? and field2 = ?", pkw)
	}

}

func TestPKWhereFieldsSingle(t *testing.T) {
	sg := &gorma.RelationalModelDefinition{}
	sg.RelationalFields = make(map[string]*gorma.RelationalFieldDefinition)
	f := &gorma.RelationalFieldDefinition{}
	f.Name = dsl.SanitizeFieldName("ID")
	f.DatabaseFieldName = dsl.SanitizeDBFieldName(f.Name)
	f.Datatype = gorma.PKInteger

	sg.RelationalFields[f.Name] = f
	sg.PrimaryKeys = append(sg.PrimaryKeys, f)

	pkw := sg.PKWhereFields()

	if pkw != "id" {
		t.Errorf("Expected %s, got %s", "id", pkw)
	}

}

func TestPKWhereFieldsMultiple(t *testing.T) {
	sg := &gorma.RelationalModelDefinition{}
	sg.RelationalFields = make(map[string]*gorma.RelationalFieldDefinition)
	f := makePK("Field1")
	sg.RelationalFields[f.Name] = f
	sg.PrimaryKeys = append(sg.PrimaryKeys, f)

	f2 := makePK("Field2")
	sg.RelationalFields[f2.Name] = f2
	sg.PrimaryKeys = append(sg.PrimaryKeys, f2)

	pkw := sg.PKWhereFields()

	if pkw != "field1,field2" {
		t.Errorf("Expected %s, got %s", "field1,field2", pkw)
	}

}

func TestPKUpdateFieldsSingle(t *testing.T) {
	sg := &gorma.RelationalModelDefinition{}
	sg.RelationalFields = make(map[string]*gorma.RelationalFieldDefinition)
	f := &gorma.RelationalFieldDefinition{}
	f.Name = dsl.SanitizeFieldName("ID")
	f.DatabaseFieldName = dsl.SanitizeDBFieldName(f.Name)
	f.Datatype = gorma.PKInteger

	sg.RelationalFields[f.Name] = f
	sg.PrimaryKeys = append(sg.PrimaryKeys, f)

	pkw := sg.PKUpdateFields("model")

	if pkw != "model.ID" {
		t.Errorf("Expected %s, got %s", "model.ID", pkw)
	}

}

func TestPKUpdateFieldsMultiple(t *testing.T) {
	sg := &gorma.RelationalModelDefinition{}
	sg.RelationalFields = make(map[string]*gorma.RelationalFieldDefinition)
	f := makePK("Field1")
	sg.RelationalFields[f.Name] = f
	sg.PrimaryKeys = append(sg.PrimaryKeys, f)

	f2 := makePK("Field2")
	sg.RelationalFields[f2.Name] = f2
	sg.PrimaryKeys = append(sg.PrimaryKeys, f2)

	pkw := sg.PKUpdateFields("model")

	if pkw != "model.Field1,model.Field2" {
		t.Errorf("Expected %s, got %s", "model.Field1,model.Field2", pkw)
	}

}
