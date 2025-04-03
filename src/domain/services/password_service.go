package services

type PasswordService interface {
	Hash(password string) (string, error)
}
