package migrate

import (
	"src.techknowlogick.com/xormigrate"
	"strconv"
	"xorm.io/xorm"
)

var migrates = make([]*xormigrate.Migration, 0)
var index = 0

func add(migrator xormigrate.MigrateFunc) {
	index++
	migrates = append(migrates, &xormigrate.Migration{
		ID:      "M_" + strconv.Itoa(index),
		Migrate: migrator,
	})
}

func Migrate(db *xorm.Engine) error {
	x := xormigrate.New(db, migrates)
	return x.Migrate()
}
