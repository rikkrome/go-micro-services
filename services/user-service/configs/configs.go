package configs

import (
	"go.rikkrome/tokyo/services/user-service/configs/databases"
)

func LoadConfigs() {
	databases.InitSQLDatabase()
}
