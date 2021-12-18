package model

import "fmt"

func (p Product) String() string {
	return fmt.Sprintf("%#v", p)
}
