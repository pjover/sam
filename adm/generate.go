package adm

import (
	"fmt"
	"sam/util"
)

type GenerateManager struct {
	postManager util.HttpPostManager
}

func NewGenerateManager() GenerateManager {
	return GenerateManager{
		util.NewHttpPostManager(),
	}
}

func (g GenerateManager) GenerateBdd() (string, error) {
	fmt.Println("Generant els rebuts ...")
	return "", nil
}
