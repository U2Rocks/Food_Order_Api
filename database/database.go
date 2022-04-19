package database

import (
	"gorm.io/gorm"
)

var (
	DBC *gorm.DB // this db variable is used throughout the program to manipulate the sqlite3 db
)
