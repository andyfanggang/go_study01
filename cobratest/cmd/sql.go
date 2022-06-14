package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"wm-motor.com/Infra/cobratest/internal/mysqlstruct"
)

var (
	DBType    string
	Host      string
	UserName  string
	Password  string
	Charset   string
	dbName    string
	tableName string
)

var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "sql转换和处理",
	Long:  "sql转换和处理",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
var sql2structCmd = &cobra.Command{
	Use:   "struct",
	Short: "sql转换",
	Long:  "sql转换",
	Run: func(cmd *cobra.Command, args []string) {

		dbinfo := mysqlstruct.DBInfo{
			DBType:   DBType,
			Host:     Host,
			UserName: UserName,
			Password: Password,
			Charset:  Charset,
		}

		dbModel := mysqlstruct.NewDBModel(&dbinfo)
		err := dbModel.Connect()
		if err != nil {
			log.Fatal("DB Connect error!", err)
		}

		columns, err := dbModel.GetColumns(dbName, tableName)
		if err != nil {
			log.Fatal("DB GetColumn error!", err)
		}
		template := mysqlstruct.NewStructTemplate()
		templatecolums := template.AssemblyColumns(columns)
		err = template.Generate(tableName, templatecolums)
		if err != nil {

			log.Fatal("template Generate error!", err)

		}

	},
}

func init() {
	sqlCmd.AddCommand(sql2structCmd)
	sql2structCmd.Flags().StringVarP(&UserName, "username", "", "", "请输入数据库用户名")
	sql2structCmd.Flags().StringVarP(&Password, "password", "", "", "请输入数据库密码")
	sql2structCmd.Flags().StringVarP(&Host, "host", "", "127.0.0.1:3306", "请输入数据库的host")
	sql2structCmd.Flags().StringVarP(&DBType, "dbtype", "", "", "请输入数据库类型")
	sql2structCmd.Flags().StringVarP(&Charset, "charset", "", "", "请输入数据库编码")
	sql2structCmd.Flags().StringVarP(&dbName, "db", "", "", "请输入数据库名")
	sql2structCmd.Flags().StringVarP(&tableName, "table", "", "", "请输入数据库表")

}
