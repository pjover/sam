package cmd

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"sam/adm"
	"testing"
)

var testNote = "Test note"

func Test_ParseInsertConsumptionsArgs(t *testing.T) {

	t.Run("Should fail with less than 3 args", func(t *testing.T) {
		var actual, err = parseInsertConsumptionsArgs([]string{"1552"}, testNote)
		assert.Equal(t, adm.InsertConsumptionsArgs{}, actual)
		assert.Equal(t, errors.New("Introdueix més de 3 arguments, has introduit 1 arguments"), err)
	})

	t.Run("Should fail with invalid child code", func(t *testing.T) {
		var actual, err = parseInsertConsumptionsArgs([]string{"0.5", "MME", "2", "MME"}, testNote)
		assert.Equal(t, adm.InsertConsumptionsArgs{}, actual)
		assert.Equal(t, errors.New("El codi d'infant introduit és invàlid: 0.5"), err)
	})

	t.Run("Should fail with invalid unit format", func(t *testing.T) {
		var actual, err = parseInsertConsumptionsArgs([]string{"2220", "0,5", "MME", "2", "AGE"}, testNote)
		assert.Equal(t, adm.InsertConsumptionsArgs{}, actual)
		assert.Equal(t, errors.New("El número introduit és invàlid: 0,5"), err)
	})

	t.Run("Should fail with invalid unit format", func(t *testing.T) {
		var actual, err = parseInsertConsumptionsArgs([]string{"2220", "MME", "2"}, testNote)
		assert.Equal(t, adm.InsertConsumptionsArgs{}, actual)
		assert.Equal(t, errors.New("El número introduit és invàlid: MME"), err)
	})

	t.Run("Should fail with invalid product code", func(t *testing.T) {
		var actual, err = parseInsertConsumptionsArgs([]string{"2220", "2", "MMME"}, testNote)
		assert.Equal(t, adm.InsertConsumptionsArgs{}, actual)
		assert.Equal(t, errors.New("El codi de producte introduit és invàlid: MMME"), err)
	})

	t.Run("Should fail without last product code", func(t *testing.T) {
		var actual, err = parseInsertConsumptionsArgs([]string{"2220", "0.5", "MME", "2"}, testNote)
		assert.Equal(t, adm.InsertConsumptionsArgs{}, actual)
		assert.Equal(t, errors.New("No s'ha indroduit el codi del darrer producte"), err)
	})

	t.Run("Should fail with duplicated product code", func(t *testing.T) {
		var actual, err = parseInsertConsumptionsArgs([]string{"2220", "0.5", "MME", "2", "MME"}, testNote)
		assert.Equal(t, adm.InsertConsumptionsArgs{}, actual)
		assert.Equal(t, errors.New("Hi ha un codi de producte repetit"), err)
	})

	t.Run("Should parse with 1 consumption", func(t *testing.T) {
		var actual, err = parseInsertConsumptionsArgs([]string{"1552", "1", "QME"}, testNote)
		assert.Equal(t, adm.InsertConsumptionsArgs{Code: 1552, Consumptions: map[string]float64{"QME": 1}, Note: testNote}, actual)
		assert.Nil(t, err)
	})

	t.Run("Should parse with 2 consumption", func(t *testing.T) {
		var actual, err = parseInsertConsumptionsArgs([]string{"2220", "0.5", "MME", "2", "AGE"}, testNote)
		assert.Equal(t, adm.InsertConsumptionsArgs{Code: 2220, Consumptions: map[string]float64{"MME": 0.5, "AGE": 2}, Note: testNote}, actual)
		assert.Nil(t, err)
	})

	t.Run("Should parse with 3 consumption", func(t *testing.T) {
		var actual, err = parseInsertConsumptionsArgs([]string{"2220", "0.5", "MME", "2", "AGE", "1", "QME", "1", "BAB"}, testNote)
		assert.Equal(t, adm.InsertConsumptionsArgs{Code: 2220, Consumptions: map[string]float64{"MME": 0.5, "AGE": 2, "QME": 1, "BAB": 1}, Note: testNote}, actual)
		assert.Nil(t, err)
	})
}
