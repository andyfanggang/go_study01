package mysqlstruct

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//定义连接数据的信息结构体
type DBInfo struct {
	DBType   string
	Host     string
	UserName string
	Password string
	Charset  string
}

//定义数据库连接池
type DBModel struct {
	DBEngine *sql.DB
	DBInfo   *DBInfo
}

//定义数据库列信息结构体,通过information_schema数据库获取
type TableColumn struct {
	ColumnName    string
	DateType      string
	IsNullable    string
	ColumnKey     string
	ColumnType    string
	ColumnComment string
}

//数据库中类型与go类型转换
var DBTypeToStructType = map[string]string{
	"int":        "int32",
	"tinyint":    "int8",
	"smallint":   "int",
	"mediumint":  "int64",
	"bigint":     "int64",
	"bit":        "int",
	"bool":       "bool",
	"enum":       "string",
	"set":        "string",
	"varchar":    "string",
	"char":       "string",
	"tinytext":   "string",
	"mediumtext": "string",
	"text":       "string",
	"longtext":   "string",
	"blob":       "string",
	"tinyblob":   "string",
	"mediumblob": "string",
	"longblob":   "string",
	"date":       "time.Time",
	"datetime":   "time.Time",
	"timestamp":  "time.Time",
	"time":       "time.Time",
	"float":      "float64",
	"double":     "float64",
}

func NewDBModel(info *DBInfo) *DBModel {

	return &DBModel{DBInfo: info}

}

//定义连接数据库的函数
func (m *DBModel) Connect() error {
	var err error
	dsn := "%s:%s@tcp(%s)/information_schema?" + "charset=%s&parseTime=True&loc=Local" //连接数据库information_schema

	s := fmt.Sprintf(
		dsn,
		m.DBInfo.UserName,
		m.DBInfo.Password,
		m.DBInfo.Host,
		m.DBInfo.Charset,
	)
	m.DBEngine, err = sql.Open(m.DBInfo.DBType, s)
	if err != nil {
		log.Fatalf("打开数据库失败!" + err.Error())
		return err
	}
	return nil
}

//定义获取列信息的函数
func (m *DBModel) GetColumns(dbName string, tableName string) ([]*TableColumn, error) {
	sqlStr := "SELECT COLUMN_NAME,DATA_TYPE,IS_NULLABLE,COLUMN_KEY,COLUMN_TYPE,COLUMN_COMMENT FROM columns WHERE TABLE_SCHEMA=? and TABLE_NAME=?"

	rows, err := m.DBEngine.Query(sqlStr, dbName, tableName)

	if err != nil {
		fmt.Println("query failed!")
		return nil, err
	}

	defer rows.Close() //must close the rows

	var columns []*TableColumn

	for rows.Next() {

		var column TableColumn
		err = rows.Scan(&column.ColumnName, &column.DateType, &column.IsNullable, &column.ColumnKey, &column.ColumnType, &column.ColumnComment)

		if err != nil {
			return nil, err
		}

		columns = append(columns, &column)
	}

	return columns, nil
}
