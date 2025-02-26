
package pg

import (
	"context"
  "gitlab.techetronventures.com/core/backend/pkg/log"
	"github.com/uptrace/bun"
	model "gitlab.techetronventures.com/core/oms-user-management/internal/oms-user-management/models"
)

func (db *DB) Health() model.HealthRepository {
	return &Health{
  		db.DB,
  		db.log,
  }
}

func (db *Tx) Health() model.HealthRepository {
	return &Health{
  		db.Tx,
  		db.log,
  	}
}

type Health struct {
	bun.IDB
	*log.Logger
}

func (b *Health) Ping(ctx context.Context) error {
	var result int
  	// Run a simple query to check the connection
  	err := b.IDB.NewRaw("SELECT 1").Scan(ctx, &result)
  	if err != nil {
  		b.Logger.Error(ctx, err.Error())
  		return err
  	}
  	return nil
}

