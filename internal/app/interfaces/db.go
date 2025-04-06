package interfaces

type DB interface {
	Instance() any
	Close()
}
