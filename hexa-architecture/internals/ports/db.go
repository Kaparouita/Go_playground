package ports

type DbRepository interface {
	ConnectDb()
	CloseDb() error
}
