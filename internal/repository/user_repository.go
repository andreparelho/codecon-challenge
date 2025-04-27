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
	GetTopCountries(total int) ([]Countries, error)
	GetActiveUsers() ([]ActiveUsers, error)
	GetMembers() ([]TeamInsights, error)
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
		if superUser.Score > 900 && superUser.Active {
			users = append(users, *superUser)
		}
	}

	return users, nil
}

func (u userRepository) GetTopCountries(total int) ([]Countries, error) {
	defer func(begin time.Time) {
		logrus.WithFields(logrus.Fields{
			"timestamp": time.Since(begin),
		}).Info("success to get topcountries")
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

	var countries []Countries
	for country, c := range count {
		countries = append(countries, Countries{
			Country: country,
			Total:   c,
		})
	}

	sort.Slice(countries, func(i, j int) bool {
		return countries[i].Total > countries[j].Total
	})

	if total <= 0 {
		total = 5
	}

	return countries[:total], nil
}

func (u userRepository) GetActiveUsers() ([]ActiveUsers, error) {
	defer func(begin time.Time) {
		logrus.WithFields(logrus.Fields{
			"timestamp": time.Since(begin),
		}).Info("success to get active users")
	}(time.Now())

	var activeUsers []ActiveUsers

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

		for _, log := range user.Logs {
			count[log.Date]++
		}
	}

	for date, t := range count {
		activeUsers = append(activeUsers, ActiveUsers{
			Date:  date,
			Total: t,
		})
	}

	sort.Slice(activeUsers, func(i, j int) bool {
		return activeUsers[i].Total > activeUsers[j].Total
	})

	return activeUsers, nil
}

func (u userRepository) GetMembers() ([]TeamInsights, error) {
	defer func(begin time.Time) {
		logrus.WithFields(logrus.Fields{
			"timestamp": time.Since(begin),
		}).Info("success to get members")
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

	team := make(map[string]*TeamInsights)
	var teamInsights []TeamInsights
	for obj := it.Next(); obj != nil; obj = it.Next() {
		user := obj.(*User)

		ti, ok := team[user.Team.Name]
		if !ok {
			ti = &TeamInsights{
				Team: user.Team.Name,
			}

			team[user.Team.Name] = ti
		}

		ti.TotalMembers++
		if user.Team.Leader {
			ti.Leaders++
		}

		for _, cp := range user.Team.Projects {
			if cp.Completed {
				ti.CompletedProjects++
			}
		}
	}

	teamInsights = make([]TeamInsights, 0, len(team))
	for _, t := range team {
		teamInsights = append(teamInsights, *t)
	}

	return teamInsights, nil
}
