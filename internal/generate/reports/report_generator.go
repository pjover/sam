package reports

type ReportGenerator interface {
	Generate() (string, error)
}
