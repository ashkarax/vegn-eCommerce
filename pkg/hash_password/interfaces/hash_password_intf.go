package interfaceHashPass

type IHashPass interface {
	HashPassword(password string) string
	CompairPassword(hashedPassword string, plainPassword string) error
}
