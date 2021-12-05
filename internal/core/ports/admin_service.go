package ports

type AdminService interface {
	Backup() (string, error)
	CreateDirectory(previousMonth bool, nextMonth bool) (string, error)
}
