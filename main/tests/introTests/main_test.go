package introtests

import (
	"testing"
	"context"
	//"os"
	"github.com/jmoiron/sqlx"
	"github.com/Mamvriyskiy/database_course/main/migrations"
	"github.com/Mamvriyskiy/database_course/main/containers"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type MyIntroTestsSuite struct {
	suite.Suite
}

// Сработает один раз перед запуском сьюта
func  (s *MyIntroTestsSuite) BeforeAll(t provider.T) {
}

// Сработает один раз после того, как все тесты завершатся
func (s *MyIntroTestsSuite) AfterAll(t provider.T) {
}

// Будет срабатывать каждый раз перед началом теста
func  (s *MyIntroTestsSuite) BeforeEach(t provider.T) {
	t.Epic("My Epic")
	t.Feature("My Feature")
	// и так далее
}

// Будет срабатывать каждый раз после окончания теста
func (s *MyIntroTestsSuite) AfterEach(t provider.T) {
}

var connDB *sqlx.DB

func TestSuiteRunner(t *testing.T) {
	dbTestContainers, db, err := containers.SetupTestDataBase()

	if err != nil {
		panic(err)
	}
	defer dbTestContainers.Terminate(context.Background())

	connDB = db
	err = migrations.MigrationsTestDataBase(connDB, "../data/data.sql")
	if err != nil {
		panic(err)
	}

	suite.RunSuite(t, new(MyIntroTestsSuite))
}
