package model

import (
  "context"
  "github.com/uptrace/bun"
  "time"
)

type HealthRepository interface {
  Ping(ctx context.Context) error
}

type Health struct {
  bun.BaseModel `bun:"table:health"`

  ID        int64     `json:"id" bun:"id,pk,autoincrement"`
  Status    string    `json:"status" bun:"status,default:'HEALTHY'"`
  CreatedAt time.Time `json:"created_at" bun:"created_at,default:current_timestamp"`
  UpdatedAt time.Time `json:"updated_at" bun:"updated_at,nullzero"`
  DeletedAt time.Time `json:"-" bun:"deleted_at,nullzero"`  // Excluded from JSON
}
