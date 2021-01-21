package postgresql

import (
	"errors"
	"github.com/go-gdbc/gdbc"
	_ "github.com/jackc/pgx/v4"
	"strings"
)

const DefaultHost = "localhost"
const DefaultPort = "5432"
const DefaultUser = "postgres"

func init() {
	gdbc.Register("pgx", "postgresql", &PostgresDataSourceNameAdapter{})
}

type PostgresDataSourceNameAdapter struct {
}

func (dsnAdapter PostgresDataSourceNameAdapter) GetDataSourceName(dataSource gdbc.DataSource) (string, error) {
	dsn := ""
	host := DefaultHost
	port := DefaultPort
	user := DefaultUser
	password := ""
	databaseName := ""

	dataSourceUrl := dataSource.GetURL()
	if dataSourceUrl.Opaque != "" {
		databaseName = dataSourceUrl.Opaque
	} else {
		if dataSourceUrl.Hostname() != "" {
			host = dataSourceUrl.Hostname()
		}

		if dataSourceUrl.Port() != "" {
			port = dataSourceUrl.Port()
		}

		if dataSourceUrl.User != nil {
			if dataSourceUrl.User.Username() != "" {
				user = dataSourceUrl.User.Username()
			}
			userPassword, _ := dataSourceUrl.User.Password()
			if userPassword != "" {
				password = userPassword
			}
		} else {
			if dataSource.GetUsername() != "" {
				user = dataSource.GetUsername()
			}
			if dataSource.GetPassword() != "" {
				password = dataSource.GetPassword()
			}
		}

		if dataSourceUrl.Path != "" {
			databaseName = dataSourceUrl.Path
		}
	}

	if strings.HasPrefix(databaseName, "/") {
		databaseName = databaseName[1:]
	}

	if strings.Contains(databaseName, "/") {
		return "", errors.New("database name format is wrong : " + databaseName)
	}

	dsn = dsn + "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + databaseName

	arguments := dataSourceUrl.Query()
	if arguments != nil {
		for argumentName, values := range arguments {
			dsn = dsn + " " + argumentName + "=" + values[0]
		}
	}
	return dsn, nil
}
