package adm

import (
	"fmt"
	"sam/comm"
)

func EditCustomer(customerCode int) error {
	url := fmt.Sprintf("http://localhost:8081/db/hobbit_prod/customer/%d", customerCode)
	return comm.OpenUrl(url)
}
