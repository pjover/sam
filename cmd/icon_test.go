package cmd

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"sam/adm"
	"testing"
)

var testNote = "Test note"

func TestParseInsertConsumptionsArgs_3_args_minimum(t *testing.T) {
	var actual, err = parseInsertConsumptionsArgs([]string{"1552"}, testNote)
	assert.Equal(t, adm.InsertConsumptionsArgs{}, actual)
	assert.Equal(t, errors.New("Introdueix més de 3 arguments, has introduit 1 arguments"), err)
}

func TestParseInsertConsumptionsArgs_invalid_child_code(t *testing.T) {
	var actual, err = parseInsertConsumptionsArgs([]string{"0.5", "MME", "2", "MME"}, testNote)
	assert.Equal(t, adm.InsertConsumptionsArgs{}, actual)
	assert.Equal(t, errors.New("El codi d'infant introduit és invàlid: 0.5"), err)
}

func TestParseInsertConsumptionsArgs_invalid_unit_format_1(t *testing.T) {
	var actual, err = parseInsertConsumptionsArgs([]string{"2220", "0,5", "MME", "2", "AGE"}, testNote)
	assert.Equal(t, adm.InsertConsumptionsArgs{}, actual)
	assert.Equal(t, errors.New("El número introduit és invàlid: 0,5"), err)
}

func TestParseInsertConsumptionsArgs_invalid_unit_format_2(t *testing.T) {
	var actual, err = parseInsertConsumptionsArgs([]string{"2220", "MME", "2"}, testNote)
	assert.Equal(t, adm.InsertConsumptionsArgs{}, actual)
	assert.Equal(t, errors.New("El número introduit és invàlid: MME"), err)
}

func TestParseInsertConsumptionsArgs_invalid_product_code(t *testing.T) {
	var actual, err = parseInsertConsumptionsArgs([]string{"2220", "2", "MMME"}, testNote)
	assert.Equal(t, adm.InsertConsumptionsArgs{}, actual)
	assert.Equal(t, errors.New("El codi de producte introduit és invàlid: MMME"), err)
}

func TestParseInsertConsumptionsArgs_missing_last_product_code(t *testing.T) {
	var actual, err = parseInsertConsumptionsArgs([]string{"2220", "0.5", "MME", "2"}, testNote)
	assert.Equal(t, adm.InsertConsumptionsArgs{}, actual)
	assert.Equal(t, errors.New("No s'ha indroduit el codi del darrer producte"), err)
}

func TestParseInsertConsumptionsArgs_duplicated_product_code(t *testing.T) {
	var actual, err = parseInsertConsumptionsArgs([]string{"2220", "0.5", "MME", "2", "MME"}, testNote)
	assert.Equal(t, adm.InsertConsumptionsArgs{}, actual)
	assert.Equal(t, errors.New("Hi ha un codi de producte repetit"), err)
}

func TestParseInsertConsumptionsArgs_ok_1_consumption(t *testing.T) {
	var actual, err = parseInsertConsumptionsArgs([]string{"1552", "1", "QME"}, testNote)
	assert.Equal(t, adm.InsertConsumptionsArgs{Code: 1552, Consumptions: map[string]float64{"QME": 1}, Note: testNote}, actual)
	assert.Nil(t, err)
}

func TestParseInsertConsumptionsArgs_ok_2_consumption(t *testing.T) {
	var actual, err = parseInsertConsumptionsArgs([]string{"2220", "0.5", "MME", "2", "AGE"}, testNote)
	assert.Equal(t, adm.InsertConsumptionsArgs{Code: 2220, Consumptions: map[string]float64{"MME": 0.5, "AGE": 2}, Note: testNote}, actual)
	assert.Nil(t, err)
}

func TestParseInsertConsumptionsArgs_ok_3_consumption(t *testing.T) {
	var actual, err = parseInsertConsumptionsArgs([]string{"2220", "0.5", "MME", "2", "AGE", "1", "QME", "1", "BAB"}, testNote)
	assert.Equal(t, adm.InsertConsumptionsArgs{Code: 2220, Consumptions: map[string]float64{"MME": 0.5, "AGE": 2, "QME": 1, "BAB": 1}, Note: testNote}, actual)
	assert.Nil(t, err)
}
