package interfaceRepository

type IAdminRepository interface {
	GetPassword(string) (string, error)
}