package postgresql

import (
	"github.com/go-gdbc/gdbc"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func getDSN(t *testing.T, dataSourceUrl string) (string, error) {
	adapter := gdbc.GetDataSourceNameAdapter("postgresql")
	dataSource, err := gdbc.GetDataSource(dataSourceUrl)
	assert.Nil(t, err)
	return adapter.GetDataSourceName(dataSource)
}

func getDSNWithUser(t *testing.T, dataSourceUrl string, username string, password string) (string, error) {
	adapter := gdbc.GetDataSourceNameAdapter("postgresql")
	dataSource, err := gdbc.GetDataSource(dataSourceUrl, gdbc.Username(username), gdbc.Password(password))
	assert.Nil(t, err)
	return adapter.GetDataSourceName(dataSource)
}

func TestPostgresDataSourceNameAdapter_GetDataSourceNameWithWrongDatabaseName(t *testing.T) {
	dsn, err := getDSN(t, "gdbc:postgresql://localhost:3000/testdb/wrong?sslmode=disable&TimeZone=Asia/Shanghai")
	assert.NotNil(t, err)
	assert.Equal(t, "database name format is wrong : testdb/wrong", err.Error())
	assert.Empty(t, dsn)
}

func TestPostgresDataSourceNameAdapter_GetDataSourceNameWithoutUser(t *testing.T) {
	dsn, err := getDSN(t, "gdbc:postgresql://localhost:3000/testdb?sslmode=disable&TimeZone=Asia/Shanghai")
	assert.Nil(t, err)
	assert.True(t, strings.Contains(dsn, "host=localhost"))
	assert.True(t, strings.Contains(dsn, "port=3000"))
	assert.True(t, strings.Contains(dsn, "user= "))
	assert.True(t, strings.Contains(dsn, "password= "))
	assert.True(t, strings.Contains(dsn, "dbname=testdb"))
	assert.True(t, strings.Contains(dsn, "TimeZone=Asia/Shanghai"))
	assert.True(t, strings.Contains(dsn, "sslmode=disable"))
}

func TestPostgresDataSourceNameAdapter_GetDataSourceNameWithoutUserAndPort(t *testing.T) {
	dsn, err := getDSN(t, "gdbc:postgresql://localhost/testdb?sslmode=disable&TimeZone=Asia/Shanghai")
	assert.Nil(t, err)
	assert.True(t, strings.Contains(dsn, "host="+DefaultHost))
	assert.True(t, strings.Contains(dsn, "port="+DefaultPort))
	assert.True(t, strings.Contains(dsn, "user= "))
	assert.True(t, strings.Contains(dsn, "password= "))
	assert.True(t, strings.Contains(dsn, "dbname=testdb"))
	assert.True(t, strings.Contains(dsn, "TimeZone=Asia/Shanghai"))
	assert.True(t, strings.Contains(dsn, "sslmode=disable"))
}

func TestPostgresDataSourceNameAdapter_GetDataSourceNameWithoutUserAndHostAndPort(t *testing.T) {
	dsn, err := getDSN(t, "gdbc:postgresql:testdb?sslmode=disable&TimeZone=Asia/Shanghai")
	assert.Nil(t, err)
	assert.True(t, strings.Contains(dsn, "host="+DefaultHost))
	assert.True(t, strings.Contains(dsn, "port="+DefaultPort))
	assert.True(t, strings.Contains(dsn, "user= "))
	assert.True(t, strings.Contains(dsn, "password= "))
	assert.True(t, strings.Contains(dsn, "dbname=testdb"))
	assert.True(t, strings.Contains(dsn, "TimeZone=Asia/Shanghai"))
	assert.True(t, strings.Contains(dsn, "sslmode=disable"))
}

func TestPostgresDataSourceNameAdapter_GetDataSourceNameWithoutUserAndHostAndPortAndDatabase(t *testing.T) {
	dsn, err := getDSN(t, "gdbc:postgresql:?sslmode=disable&TimeZone=Asia/Shanghai")
	assert.Nil(t, err)
	assert.True(t, strings.Contains(dsn, "host="+DefaultHost))
	assert.True(t, strings.Contains(dsn, "port="+DefaultPort))
	assert.True(t, strings.Contains(dsn, "user= "))
	assert.True(t, strings.Contains(dsn, "password= "))
	assert.True(t, strings.Contains(dsn, "dbname= "))
	assert.True(t, strings.Contains(dsn, "TimeZone=Asia/Shanghai"))
	assert.True(t, strings.Contains(dsn, "sslmode=disable"))
}

func TestPostgresDataSourceNameAdapter_GetDataSourceNameWithFullFormat(t *testing.T) {
	dsn, err := getDSN(t, "gdbc:postgresql://username:password@localhost:3000/testdb?sslmode=disable&TimeZone=Asia/Shanghai")
	assert.Nil(t, err)
	assert.True(t, strings.Contains(dsn, "host=localhost"))
	assert.True(t, strings.Contains(dsn, "port=3000"))
	assert.True(t, strings.Contains(dsn, "user=username"))
	assert.True(t, strings.Contains(dsn, "password=password"))
	assert.True(t, strings.Contains(dsn, "dbname=testdb"))
	assert.True(t, strings.Contains(dsn, "TimeZone=Asia/Shanghai"))
	assert.True(t, strings.Contains(dsn, "sslmode=disable"))
}

func TestPostgresDataSourceNameAdapter_GetDataSourceNameWithFullFormatAndUserOptions(t *testing.T) {
	dsn, err := getDSNWithUser(t, "gdbc:postgresql://localhost:3000/testdb?sslmode=disable&TimeZone=Asia/Shanghai", "username", "password")
	assert.Nil(t, err)
	assert.True(t, strings.Contains(dsn, "host=localhost"))
	assert.True(t, strings.Contains(dsn, "port=3000"))
	assert.True(t, strings.Contains(dsn, "user=username"))
	assert.True(t, strings.Contains(dsn, "password=password"))
	assert.True(t, strings.Contains(dsn, "dbname=testdb"))
	assert.True(t, strings.Contains(dsn, "TimeZone=Asia/Shanghai"))
	assert.True(t, strings.Contains(dsn, "sslmode=disable"))
}
