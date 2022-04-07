package test_data

import "github.com/pjover/sam/internal/domain/model"

var ProductTST, _ = model.NewProduct(
	"TST",
	"Test product",
	"TstProduct",
	10.9,
	0.0,
	false,
)

var ProductXXX, _ = model.NewProduct(
	"XXX",
	"XXX product",
	"XxxProduct",
	9.1,
	0.0,
	false,
)

var ProductYYY, _ = model.NewProduct(
	"YYY",
	"YYY product",
	"YyyProduct",
	5,
	0.1,
	false,
)
