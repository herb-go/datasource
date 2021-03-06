package sqlitecolumns

import (
	"errors"
	"strings"

	_ "github.com/mattn/go-sqlite3" //sqlite driver

	"github.com/herb-go/datasource/sql/db"
	"github.com/herb-go/datasource/sql/db/columns"
)

//Column sqlite column struct
type Column struct {
	CID     int64
	Name    string
	Field   string
	Type    string
	NotNull string
	Default interface{}
	Key     string
}

// ConvertType convert culumn type to golang type.
func ConvertType(t string) (string, error) {
	ft := strings.Split(t, "(")[0]
	switch strings.ToUpper(ft) {
	case "BOOL":
		return "byte", nil
	case "SMALLINT", "MEDIUMINT", "INT", "INTEGER", "TINYINT", "INT2", "INT8":
		return "int", nil
	case "BIGINT", "INT64":
		return "int64", nil
	case "FLOAT":
		return "float32", nil
	case "DOUBLE", "DOUBLE PRECISION", "REAL":
		return "float64", nil
	case "DATETIME", "DATE":
		return "time.Time", nil
	case "CHAR", "VARCHAR", "CHARACTER", "NCHAR", "NVARCHAR", "TEXT":
		return "string", nil
	case "BLOB":
		return "[]byte", nil
	}
	return "", errors.New("sqlitecolumns:Column type " + t + " is not supported.")

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
	if (output.ColumnType == "int" || output.ColumnType == "int64") && c.Key == "1" {
		output.AutoValue = true
	}
	if c.Key == "1" {
		output.PrimayKey = true
	}
	if c.NotNull == "1" {
		output.NotNull = true
	}

	return output, nil
}

// Columns sqlite columns type
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
	rows, err := db.Query("PRAGMA table_info(" + table + ")")
	if err != nil {
		return err
	}
	defer rows.Close()
	*c = []Column{}
	for rows.Next() {
		column := Column{}
		if err := rows.Scan(&column.CID, &column.Field, &column.Type, &column.NotNull, &column.Default, &column.Key); err != nil {
			return err
		}
		*c = append(*c, column)
	}
	return nil
}

func init() {
	columns.Register("sqlite3", func() columns.Loader {
		return &Columns{}
	})
}
