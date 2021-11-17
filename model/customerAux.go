package model

import (
	"fmt"
	"strings"
)

func (c Customer) FirstAdult() Adult {
	for _, adult := range c.Adults {
		if adult.Role == "MOTHER" {
			return adult
		}
	}
	return c.Adults[0]
}

func (c Customer) FirstAdultNameWithCode() string {
	adult := c.FirstAdult()
	return fmt.Sprintf("%s %s (%d)", adult.Name, adult.Surname, c.Id())
}

func (c Customer) Id() int {
	return c.Children[0].Code / 10
}

func (c Customer) ChildrenNames(joinWith string) string {
	var names []string
	for _, child := range c.Children {
		names = append(names, child.NameWithCode())
	}
	return strings.Join(names, joinWith)
}

func (c Child) NameWithCode() string {
	return fmt.Sprintf("%s %s (%d)", c.Name, c.Surname, c.Code)
}
