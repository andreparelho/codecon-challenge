package repository

import (
	"sort"
	"time"

	"github.com/hashicorp/go-memdb"
	"github.com/sirupsen/logrus"
)

type UserRepository interface {
	SaveUsers(user []*User) error
	GetSuperusers() ([]User, error)
	GetTopCountries() ([]string, error)
}

type userRepository struct {
	Database *memdb.MemDB
}

func NewUserRepository(db *memdb.MemDB) userRepository {
	return userRepository{
		Database: db,
	}
}

func (u userRepository) SaveUsers(users []*User) error {
	defer func(begin time.Time) {
		logrus.WithFields(logrus.Fields{
			"timestamp": time.Since(begin),
		}).Info("success to saving users")
	}(time.Now())

	txn := u.Database.Txn(true)

	for _, usr := range users {
		if err := txn.Insert("user", usr); err != nil {
			logrus.WithFields(logrus.Fields{
				"err":  err,
				"user": usr,
			}).Error("error to insert usr on database")

			txn.Abort()
			return err
		}
	}

	txn.Commit()

	return nil
}

func (u userRepository) GetSuperusers() ([]User, error) {
	defer func(begin time.Time) {
		logrus.WithFields(logrus.Fields{
			"timestamp": time.Since(begin),
		}).Info("success to get superusers")
	}(time.Now())

	txn := u.Database.Txn(false)
	defer txn.Abort()

	it, err := txn.Get("user", "id")
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("error to get users on database")

		return nil, err
	}

	var users []User
	for obj := it.Next(); obj != nil; obj = it.Next() {
		superUser := obj.(*User)
		if superUser.Score > 900 && !superUser.Active {
			users = append(users, *superUser)
		}
	}

	return users, nil
}

func (u userRepository) GetTopCountries() ([]string, error) {
	defer func(begin time.Time) {
		logrus.WithFields(logrus.Fields{
			"timestamp": time.Since(begin),
		}).Info("success to get superusers")
	}(time.Now())

	txn := u.Database.Txn(false)
	defer txn.Abort()

	it, err := txn.Get("user", "id")
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("error to get users on database")

		return nil, err
	}

	count := make(map[string]int)
	for obj := it.Next(); obj != nil; obj = it.Next() {
		user := obj.(*User)
		count[user.Country]++
	}

	var cf []CountriesFrequency
	for country, count := range count {
		cf = append(cf, CountriesFrequency{
			Country: country,
			Count:   count,
		})
	}

	sort.Slice(cf, func(i, j int) bool {
		return cf[i].Count > cf[j].Count
	})

	top := 5
	if len(cf) < 5 {
		top = len(cf)
	}

	var countries []string
	for i := 0; i < top; i++ {
		countries = append(countries, cf[i].Country)
	}

	return countries, nil
}
