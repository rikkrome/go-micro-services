package configs

import (
	"github.com/rikkrome/go-micro-services/services/auth-service/configs/databases"
)

func LoadConfigs() {
	databases.InitSQLDatabase()
}
