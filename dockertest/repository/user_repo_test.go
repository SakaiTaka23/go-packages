package repository

import (
	"database/sql"
	"dockertest/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gotest.tools/v3/assert"
	"log"
	"os"
	"reflect"
	"testing"
	"time"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	var sqlDB *sql.DB
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	pool.MaxWait = 2 * time.Minute
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	runOputions := &dockertest.RunOptions{
		Repository: "mysql",
		Tag:        "8.0",
		Env: []string{
			"MYSQL_ROOT_PASSWORD=secret",
			"MYSQL_DATABASE=example",
		},
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(runOputions, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})

	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		time.Sleep(time.Second * 3)
		var err error
		sqlDB, err = sql.Open("mysql", fmt.Sprintf("root:secret@tcp(%s)/example", resource.GetHostPort("3306/tcp")))
		if err != nil {
			return err
		}

		if err = sqlDB.Ping(); err != nil {
			return err
		}

		db, err = gorm.Open(mysql.New(mysql.Config{
			Conn: sqlDB,
		}), &gorm.Config{})
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	_ = db.AutoMigrate(&model.User{})

	code := m.Run()

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func Test_sqlUser_Create(t *testing.T) {
	repo := NewUserRepository(db)
	id, err := repo.Create("John Doe")
	assert.NilError(t, err)
	assert.Assert(t, id != 0)
}

func Test_sqlUser_FindAll(t *testing.T) {
	createMockUsers(db)

	repo := NewUserRepository(db)
	users, err := repo.FindAll()

	expect := []model.User{
		{ID: 1, Name: "John Doe1"},
		{ID: 2, Name: "Jane Doe2"},
		{ID: 3, Name: "Jack Doe3"},
	}

	assert.NilError(t, err)
	assert.Assert(t, len(users) == 3)
	assert.Assert(t, reflect.DeepEqual(users, expect))
}

func createMockUsers(db *gorm.DB) {
	users := []*model.User{
		{ID: 1, Name: "John Doe1"},
		{ID: 2, Name: "Jane Doe2"},
		{ID: 3, Name: "Jack Doe3"},
	}
	db.Create(users)
}
