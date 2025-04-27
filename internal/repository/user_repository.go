package repository

import (
	"github.com/hashicorp/go-memdb"
	"github.com/sirupsen/logrus"
)

type UserRepository interface {
	SaveUsers(user []*User) error
}

type userRepository struct {
	WriteTransacation *memdb.Txn
	ReadTransaction   *memdb.Txn
}

func NewUserRepository(w, r *memdb.Txn) userRepository {
	return userRepository{
		WriteTransacation: w,
		ReadTransaction:   r,
	}
}

func (u userRepository) SaveUsers(users []*User) error {
	for _, usr := range users {
		if err := u.WriteTransacation.Insert("user", usr); err != nil {
			logrus.WithFields(logrus.Fields{
				"err":  err,
				"user": usr,
			}).Error("error to insert usr on database")

			u.WriteTransacation.Abort()
			return err
		}
	}

	u.WriteTransacation.Commit()

	logrus.Info("success to saving users")
	return nil
}
