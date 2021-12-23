package model

import "fmt"

func (p Product) String() string {
	return fmt.Sprintf("%s  % 7.2f  %s", p.Id, p.Price, p.Name)
}
