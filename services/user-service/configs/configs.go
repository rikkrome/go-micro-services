package configs

import (
	"github.com/rikkrome/go-micro-services/services/user-service/configs/databases"
)

func LoadConfigs() {
	databases.InitSQLDatabase()
}
