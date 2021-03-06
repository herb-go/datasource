package mysqlcolumns

import (
	"errors"
	"strings"

	_ "github.com/go-sql-driver/mysql" //mysql driver

	"github.com/herb-go/datasource/sql/db"
	"github.com/herb-go/datasource/sql/db/columns"
)

//Column mysql column struct
type Column struct {
	// Field  column name
	Field string
	// Type column type
	Type string
	// IsNull if column can be null
	IsNull string
	// Key  if column is primary key
	Key string
	// Default default value
	Default interface{}
	// Extra extra data
	Extra string
}

// ConvertType convert culumn type to golang type.
func ConvertType(t string) (string, error) {
	ft := strings.Split(t, "(")[0]
	switch strings.ToUpper(ft) {
	case "TINYINT", "BIT", "BOOL":
		return "byte", nil
	case "SMALLINT", "MEDIUMINT", "INT", "INTEGER":
		return "int", nil
	case "BIGINT":
		return "int64", nil
	case "FLOAT":
		return "float32", nil
	case "DOUBLE", "DOUBLE PRECISION":
		return "float64", nil
	case "DATETIME", "TIMESTAMP":
		return "time.Time", nil
	case "CHAR", "VARCHAR", "TINYTEXT", "TEXT", "MEDIUMTEXT", "LONGTEXT", "DECIMAL":
		return "string", nil
	case "BINARY", "VARBINARY", "TINYBLOB", "BLOB", "MEDIUMBLOB", "LONGBLOB":
		return "[]byte", nil
	}
	return "", errors.New("mysqlColumn:column type " + t + " is not supported.")

}

// Convert convert MysqlColumn to commn column
func (c *Column) Convert() (*columns.Column, error) {
	output := &columns.Column{}
	output.Field = c.Field
	t, err := ConvertType(c.Type)
	output.ColumnType = t
	if err != nil {
		return nil, err
	}
	if output.ColumnType == "time.Time" && c.Default != nil {
		output.AutoValue = true
	}
	if strings.Contains(c.Extra, "auto_increment") {
		output.AutoValue = true
	}
	if strings.Contains(c.Key, "PRI") {
		output.PrimayKey = true
	}
	if c.IsNull == "NO" {
		output.NotNull = true
	}

	return output, nil
}

// Columns mysql columns type
type Columns []Column

// Columns return loaded columns
func (c *Columns) Columns() ([]*columns.Column, error) {
	output := []*columns.Column{}
	for _, v := range *c {
		column, err := v.Convert()
		if err != nil {
			return nil, err
		}
		output = append(output, column)
	}
	return output, nil
}

// Load load columns with given database and table name
func (c *Columns) Load(conn db.Database, table string) error {
	db := conn.DB()
	rows, err := db.Query("desc " + table)
	if err != nil {
		return err
	}
	defer rows.Close()
	*c = []Column{}
	for rows.Next() {
		column := Column{}
		if err := rows.Scan(&column.Field, &column.Type, &column.IsNull, &column.Key, &column.Default, &column.Extra); err != nil {
			return err
		}
		*c = append(*c, column)
	}
	return nil
}

func init() {
	columns.Register("mysql", func() columns.Loader {
		return &Columns{}
	})
}
