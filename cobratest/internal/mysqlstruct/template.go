package mysqlstruct

import (
	"fmt"
	"html/template"
	"os"

	"wm-motor.com/Infra/cobratest/internal/word"
)

//定义模板
const strcutTpl = `type {{.TableName | ToCamelCase}} struct{
{{range .Columns}}
{{ $length := len .Comment}}
{{if gt $length 0}} 
// {{.Comment}}
{{else}}
// {{.Name}}
{{end}}
{{ $typeLen := len .Type }}
 {{ if gt $typeLen 0 }}
 {{.Name | ToCamelCase}}	{{.Type}}	{{.Tag}}
 {{else}}
{{.Name}}
{{end}}
{{end}}

func (model {{.TableName | ToCamelCase}}) TableName() string {
	return "{{.TableName}}"
}`

//定义模板结构体
type StructTemplate struct {
	strcutTpl string
}

//定义接收数据库字段信息结构体
type StructColumn struct {
	Name    string
	Type    string
	Tag     string
	Comment string
}

//定义数据库查询得到的结构体
type StructTemplateDB struct {
	TableName string
	Columns   []*StructColumn
}

//初始化模板
func NewStructTemplate() *StructTemplate {
	return &StructTemplate{strcutTpl: strcutTpl}
}

//根据数据库列信息获取列相关信息
func (t *StructTemplate) AssemblyColumns(tbcolums []*TableColumn) []*StructColumn {

	tplColumns := make([]*StructColumn, 0, len(tbcolums))

	for _, column := range tbcolums {
		tag := fmt.Sprintf("`"+"json:"+"\"%s\""+"`", column.ColumnName)
		tplColumns = append(tplColumns, &StructColumn{
			Name:    column.ColumnName,
			Type:    column.ColumnType,
			Tag:     tag,
			Comment: column.ColumnComment,
		})
	}
	return tplColumns
}

//模板渲染
func (t *StructTemplate) Generate(tableName string, tplColumns []*StructColumn) error {

	funcMap := template.FuncMap{"ToCamelCase": word.UnderscoreToUpperCamelCase} //定义模板调用的函数
	tp1 := template.New("mysqlstruct")
	tp1, _ = tp1.Funcs(funcMap).Parse(t.strcutTpl)

	tp1DB := StructTemplateDB{
		TableName: tableName,
		Columns:   tplColumns,
	}
	err := tp1.Execute(os.Stdout, tp1DB) //调用模板渲染
	if err != nil {

		return err
	}
	return nil
}
