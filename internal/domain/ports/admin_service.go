package ports

type AdminService interface {
	Backup() (string, error)
	CreateDirectories() (string, error)
}
