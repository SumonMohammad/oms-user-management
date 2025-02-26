package model

import (
	"context"
	"github.com/uptrace/bun"
	"gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
	"time"
)

type AuditLogs interface {
	CreateAuditLog(ctx context.Context, auditLog *AuditLog) error
	GetAuditLogs(ctx context.Context, req *grpc.GetAuditLogsRequest) ([]*AuditLog, int64, error)
	GetAuditLogById(ctx context.Context, req *grpc.GetAuditLogByIdRequest) (*grpc.GetAuditLogByIdResponse, error)
}

type AuditLog struct {
	bun.BaseModel `bun:"table:audit_log"`

	ID            int64     `json:"id" bun:"id,pk,autoincrement"`
	Type          string    `json:"type" bun:"type"`                       // E.g., login, password change
	AttemptByID   int64     `json:"attempt_by_id" bun:"attempt_by_id"`     // ID of trader, investor, or broker_admin
	AttemptByType string    `json:"attempt_by_type" bun:"attempt_by_type"` // Specifies "trader", "investor", or "broker_admin"
	IPAddress     string    `json:"ip_address" bun:"ip_address"`
	ActionType    string    `json:"action_type" bun:"action_type"`
	HTTPMethod    string    `json:"http_method" bun:"http_method"`
	Endpoint      string    `json:"endpoint" bun:"endpoint"`
	IsSuccess     bool      `json:"is_success" bun:"is_success"`
	Platform      string    `json:"platform" bun:"platform"`
	DeviceName    string    `json:"device_name" bun:"device_name"`
	DeviceType    string    `json:"device_type" bun:"device_type"`
	Description   string    `json:"description" bun:"description,type:json"`
	RequestBody   string    `json:"request_body" bun:"request_body,type:json"`
	ResponseBody  string    `json:"response_body" bun:"response_body,type:json"`
	CreatedAt     time.Time `json:"created_at" bun:"created_at,default:current_timestamp"`
	UpdatedAt     time.Time `json:"updated_at" bun:"updated_at,nullzero"`
}
