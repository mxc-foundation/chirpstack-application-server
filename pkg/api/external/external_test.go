package external

import (
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/suite"

	"github.com/mxc-foundation/chirpstack-application-server/pkg/storage"
	"github.com/mxc-foundation/chirpstack-application-server/pkg/test"
)

// DatabaseTestSuiteBase provides the setup and teardown of the database
// for every test-run.
type DatabaseTestSuiteBase struct {
	tx *storage.TxLogger
}

// SetupSuite is called once before starting the test-suite.
func (b *DatabaseTestSuiteBase) SetupSuite() {
	conf := test.GetConfig()
	if err := storage.Setup(conf); err != nil {
		panic(err)
	}
}

// SetupTest is called before every test.
func (b *DatabaseTestSuiteBase) SetupTest() {
	tx, err := storage.DB().Beginx()
	if err != nil {
		panic(err)
	}
	b.tx = tx

	test.MustFlushRedis(storage.RedisPool())
	test.MustResetDB(storage.DB().DB)
}

// TearDownTest is called after every test.
func (b *DatabaseTestSuiteBase) TearDownTest() {
	if err := b.tx.Rollback(); err != nil {
		panic(err)
	}
}

// Tx returns a database transaction (which is rolled back after every
// test).
func (b *DatabaseTestSuiteBase) Tx() sqlx.Ext {
	return b.tx
}

type APITestSuite struct {
	suite.Suite
	DatabaseTestSuiteBase
}

func TestAPI(t *testing.T) {
	suite.Run(t, new(APITestSuite))
}
