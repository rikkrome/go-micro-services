package configs

import (
	"go.rikkrome/go-micro-services/services/auth-service/configs/databases"
)

func LoadConfigs() {
	databases.InitSQLDatabase()
}
