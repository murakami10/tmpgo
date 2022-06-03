package error

import (
	"errors"
	"fmt"
	"log"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrUnknown        = errors.New("unknown error")
)

func main() {
	u := UserService{}
	if err := u.Authenticate(); err != nil {
		log.Fatal()
	}
}

type UserService struct {
	userRepo userRepository
}

func (s *UserService) Authenticate() error {
	if err := s.userRepo.FindUser("hoge"); err != nil {
		return err
		if errors.Is(err, ErrRecordNotFound) {
			log.Fatal("ユーザが存在しない")
		} else {
			return fmt.Errorf("failed to userRepo.FindUser: %w", err)
		}
	}

	return nil
}

type userRepository struct {
	db dbHandler
}

func (r *userRepository) FindUser(id string) error {
	if err := r.db.Query("SELECT * FROM users WHERE id = ?", id); err != nil {
		return fmt.Errorf("faild to find user: %w", err)
	}
	return nil
}

type dbHandler struct{}

func (h dbHandler) Query(sql string, args ...interface{}) error {
	return ErrRecordNotFound
}
