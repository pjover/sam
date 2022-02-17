package model

import (
	"fmt"
	"strings"
)

func (c Customer) String() string {
	return fmt.Sprintf("%d  %-25s  %-55s  %s", c.Id, c.FirstAdultName(), c.ChildrenNames(", "), c.InvoiceHolder.PaymentInfoFmt())
}

func (c Customer) FirstAdult() Adult {
	for _, adult := range c.Adults {
		if adult.Role == "MOTHER" {
			return adult
		}
	}
	return c.Adults[0]
}

func (c Customer) FirstAdultName() string {
	adult := c.FirstAdult()
	return fmt.Sprintf("%s %s", adult.Name, adult.Surname)
}

func (c Customer) FirstAdultNameWithId() string {
	adult := c.FirstAdult()
	return fmt.Sprintf("%s %s (%d)", adult.Name, adult.Surname, c.Id)
}

func (c Customer) ChildrenNames(joinWith string) string {
	var names []string
	for _, child := range c.Children {
		names = append(names, child.NameWithId())
	}
	return strings.Join(names, joinWith)
}

func (c Child) String() string {
	return fmt.Sprintf("%d  %-30s  %s  %s", c.Id, c.NameAndSurname(), c.Group, c.BirthDate.Format("2006-01-02"))
}

func (c Child) NameAndSurname() string {
	return fmt.Sprintf("%s %s", c.Name, c.Surname)
}

func (c Child) NameWithId() string {
	return fmt.Sprintf("%s %s (%d)", c.Name, c.Surname, c.Id)
}

func (i InvoiceHolder) PaymentInfoFmt() string {
	switch i.PaymentType {
	case "BANK_DIRECT_DEBIT":
		return fmt.Sprintf("Rebut %s", i.BankAccountFmt())
	case "BANK_TRANSFER":
		return fmt.Sprintf("Trans. %s", i.BankAccountFmt())
	case "CASH":
		return "Efectiu"
	case "VOUCHER":
		return "Xec escoleta"
	default:
		return "Indefinit"
	}
}

func (i InvoiceHolder) BankAccountFmt() string {
	if len(i.BankAccount) != 24 {
		return i.BankAccount
	}
	return fmt.Sprintf(
		"%s %s %s %s %s %s",
		i.BankAccount[0:4],
		i.BankAccount[4:8],
		i.BankAccount[8:12],
		i.BankAccount[12:16],
		i.BankAccount[16:20],
		i.BankAccount[20:24],
	)
}

func (i InvoiceHolder) Mail() string {
	return fmt.Sprintf("%s <%s>", i.Name, i.Email)
}

func (a Adult) MobilePhoneFmt() string {
	if len(a.MobilePhone) != 9 {
		return a.MobilePhone
	}
	return fmt.Sprintf(
		"%s %s %s",
		a.MobilePhone[0:3],
		a.MobilePhone[3:6],
		a.MobilePhone[6:9],
	)
}

func (a Adult) NameSurnameFmt() string {
	return fmt.Sprintf("%s %s", a.Name, a.Surname)
}
