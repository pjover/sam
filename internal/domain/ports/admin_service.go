package ports

type AdminService interface {
	Backup() (string, error)
	CreateWorkingDirectory() (string, error)
}
