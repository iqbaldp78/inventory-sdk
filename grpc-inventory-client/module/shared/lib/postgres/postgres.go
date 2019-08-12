//Package postgres used as driver to communicate with postgres db
package postgres

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/gocraft/dbr"

	//import postgres lib to this package
	_ "github.com/lib/pq"
)

//ConfigTest used as configuration for testing
type ConfigTest struct {
	DBUser     string `json:"DB_USER"`
	DBPassword string `json:"DB_PASSWORD"`
	DBHost     string `json:"DB_HOST"`
	DBPort     int    `json:"DB_PORT"`
	DBName     string `json:"DB_NAME"`
}

var conf *ConfigTest
var confString = `{"DB_USER":"postgres","DB_PASSWORD":"","DB_HOST":"localhost","DB_PORT":5432,"DB_NAME":"test"}`

//database used as interface to connect
var database *dbr.Session

const (
	existQuery     string = "SELECT EXISTS(SELECT datname FROM pg_catalog.pg_database WHERE lower(datname) = lower('%v'))"
	terminateQuery string = "SELECT pg_terminate_backend (pg_stat_activity.pid) FROM pg_stat_activity WHERE pg_stat_activity.datname = '%v'"
	dropQuery      string = "DROP DATABASE %v"
	createQuery    string = "CREATE DATABASE %v"
)

func init() {
	config := ConfigTest{}
	if os.Getenv("TEST_CONF") != "" {
		json.Unmarshal([]byte(os.Getenv("TEST_CONF")), &config)
	} else {
		json.Unmarshal([]byte(confString), &config)
	}

	conf = &config
}

//New used for setup connection to postgres
func New(dbname, user, password, host string, port int) error {
	// create a normal database connection through database/sql
	dsn := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%d sslmode=disable", dbname, user, password, host, port)
	// db, err := sql.Open("postgres", "dbname=dn_core user=postgres password=asdf1234 host=localhost port=5432 sslmode=disable")
	db, _ := dbr.Open("postgres", dsn, nil)

	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(10)

	database = db.NewSession(nil)
	err := database.Ping()
	if err != nil {
		return err
	}
	return nil
}

//GenerateTestDB used for generate db for testing
func GenerateTestDB() (*dbr.Session, error) {
	return setupTestDB()
}

func setupTestDB() (db *dbr.Session, err error) {
	var exist bool

	New("postgres", conf.DBUser, conf.DBPassword, conf.DBHost, conf.DBPort)
	db = GetDB()
	db.SelectBySql(fmt.Sprintf(existQuery, conf.DBName)).Load(&exist)

	if exist {
		db.SelectBySql(fmt.Sprintf(terminateQuery, conf.DBName)).Load(&exist)
		if _, err = db.Query(fmt.Sprintf(dropQuery, conf.DBName)); err != nil {
			return
		}
	} else {
		if _, err = db.Query(fmt.Sprintf(createQuery, conf.DBName)); err != nil {
			return
		}
	}

	db.Close()

	if !exist {
		New(conf.DBName, conf.DBUser, conf.DBPassword, conf.DBHost, conf.DBPort)
		db = GetDB()
		_, c, _, _ := runtime.Caller(0)
		dir := filepath.Dir(c)
		b, _ := ioutil.ReadFile(strings.Join([]string{dir, "test-db.sql"}, string(os.PathSeparator)))
		_, err = db.Query(string(b))
	}
	return
}

//GetDB used for get current database connection
func GetDB() *dbr.Session {
	return database
}
