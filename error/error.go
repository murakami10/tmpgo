package error

import (
	"log"

	"github.com/pkg/errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrUnknown        = errors.New("unknown error")
)

func main() {
	u := UserService{}
	if err := u.Authenticate(); err != nil {
		log.Fatal(err)
	}
}

type UserService struct {
	userRepo userRepository
}

func (s *UserService) Authenticate() error {
	if err := s.userRepo.FindUser("hoge"); err != nil {
		switch errors.Cause(err) {
		case ErrRecordNotFound:
			log.Print("ユーザ存在しない")
		default:
			return err
		}
	}

	return nil
}

type userRepository struct {
	db dbHandler
}

func (r *userRepository) FindUser(id string) error {
	if err := r.db.Query("SELECT * FROM users WHERE id = ?", id); err != nil {
		return err
	}
	return nil
}

type dbHandler struct{}

func (h dbHandler) Query(sql string, args ...interface{}) error {
	//return errors.WithStack(ErrRecordNotFound)
	return errors.WithStack(ErrUnknown)
}
