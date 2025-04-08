package interfaces

type Bcrypt interface {
	Hash(data string) (string, error)
	CompareHash(hash, data string) bool
}
