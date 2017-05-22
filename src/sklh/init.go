package sklh

// Init : used to initiate all required func in package
func Init() (err error) {
	// initiate prepared statement, file name : `stmt.go`
	err = InitStatement()
	return err
}
