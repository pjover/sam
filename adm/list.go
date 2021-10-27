package adm

import (
	"fmt"
	"sam/comm"
)

func ListEmails(ei1 bool, ei2 bool, ei3 bool) error {
	var group string
	if ei1 {
		group = "EI_1"
	} else if ei2 {
		group = "EI_2"
	} else if ei3 {
		group = "EI_3"
	} else {
		group = "ALL"
	}
	url := fmt.Sprintf("http://localhost:8080/lists/emails/%s", group)
	return comm.PrintUrl(url)
}

func ListCustomers() error {
	return comm.PrintUrl("http://localhost:8080/lists/customers")
}

func ListChildren() error {
	return comm.PrintUrl("http://localhost:8080/lists/children")
}

func ListAllCustomersConsumptions() error {
	fmt.Println("Llistant els consums de tots els clients")
	return nil
}

func ListCustomerConsumptions(customerCode int) error {
	fmt.Println("Llistant els consums del client", customerCode)
	return nil
}
