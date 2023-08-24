package configs

import (
	"go.rikkrome/tokyo/services/auth-service/configs/databases"
)

func LoadConfigs() {
	databases.InitSQLDatabase()
}
