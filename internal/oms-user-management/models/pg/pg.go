
package pg

import (
  "context"
  "database/sql"

  "github.com/uptrace/bun"
  "gitlab.techetronventures.com/core/backend/pkg/bundb"
  "gitlab.techetronventures.com/core/backend/pkg/log"
  model "gitlab.techetronventures.com/core/oms-user-management/internal/oms-user-management/models"
)

// compiler time check
var _ model.DB = (*DB)(nil)

// DB is a database representation
type DB struct {
  // *bun.IDB
  *bun.DB // underlying go-pg DB wrapper instance
  log     *log.Logger
}

// Tx represents transactions
type Tx struct {
  *bun.Tx
  log *log.Logger
}

// compiler time check
var _ model.Repository = (*Tx)(nil)

// New DB with given configurations and logger.
func New(conf *bundb.Config, logger *log.Logger) (db *DB, err error) {

  var pg *bundb.DB
  if pg, err = bundb.New(conf); err != nil {
    return
  }

  db = new(DB)
  db.DB = pg.DB // embed
  db.registerTables()
  db.log = logger.Named("db_model")
  db.log.Info(context.Background(), "db initialization done")
  return
}

// register all tables for relations
func (db *DB) registerTables() {

}

// InTx runs given function in SQL-transaction.
func (db *DB) InTx(ctx context.Context, txFunc model.TxFunc) (err error) {

  err = db.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) (err error) {
    return txFunc(ctx, &Tx{
      &tx,
      db.log,
    })
  })
  return
}
