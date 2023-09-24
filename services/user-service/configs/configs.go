package configs

import (
	"go.rikkrome/go-micro-services/services/user-service/configs/databases"
)

func LoadConfigs() {
	databases.InitSQLDatabase()
}
