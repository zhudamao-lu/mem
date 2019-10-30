package mem

import (
	"testing"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func insert() error {
	db, err := sql.Open("mysql", "mosalut:12345678@tcp(172.19.70.171)/jx?charset=utf8")
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("insert into account values(null, ?, md5(?), null, ?, null, ?, null, 0, 0, null, null)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec("account", "password", "email@email.com", "nick")
	if err != nil {
		return err
	}

	return nil
}

func TestMem(t *testing.T) {
	err := insert()
	if err != nil {
		dbError := &DbErr{}
		err := dbError.Mapping(err)
		if err != nil {
			t.Error(err.Error())
			return
		}

		t.Log(dbError.Number)
		t.Log(dbError.Message)
		t.Log(dbError.Error())
		return
	}
}
