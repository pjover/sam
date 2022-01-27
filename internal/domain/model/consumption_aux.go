package model

import "fmt"

func (c Consumption) String() string {
	return fmt.Sprintf("%d  %s  %4.1f  %s  %-5v  %s", c.ChildCode, c.YearMonth, c.Units, c.ProductID, c.IsRectification, c.InvoiceCode)
}
