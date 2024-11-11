package sql

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/authorizerdev/authorizer/internal/graph/model"
	"github.com/authorizerdev/authorizer/internal/models/schemas"
)

// AddWebhookLog to add webhook log
func (p *provider) AddWebhookLog(ctx context.Context, webhookLog *schemas.WebhookLog) (*model.WebhookLog, error) {
	if webhookLog.ID == "" {
		webhookLog.ID = uuid.New().String()
	}

	webhookLog.Key = webhookLog.ID
	webhookLog.CreatedAt = time.Now().Unix()
	webhookLog.UpdatedAt = time.Now().Unix()
	res := p.db.Clauses(
		clause.OnConflict{
			DoNothing: true,
		}).Create(&webhookLog)
	if res.Error != nil {
		return nil, res.Error
	}

	return webhookLog.AsAPIWebhookLog(), nil
}

// ListWebhookLogs to list webhook logs
func (p *provider) ListWebhookLogs(ctx context.Context, pagination *model.Pagination, webhookID string) (*model.WebhookLogs, error) {
	var webhookLogs []schemas.WebhookLog
	var result *gorm.DB
	var totalRes *gorm.DB
	var total int64

	if webhookID != "" {
		result = p.db.Where("webhook_id = ?", webhookID).Limit(int(pagination.Limit)).Offset(int(pagination.Offset)).Order("created_at DESC").Find(&webhookLogs)
		totalRes = p.db.Where("webhook_id = ?", webhookID).Model(&schemas.WebhookLog{}).Count(&total)
	} else {
		result = p.db.Limit(int(pagination.Limit)).Offset(int(pagination.Offset)).Order("created_at DESC").Find(&webhookLogs)
		totalRes = p.db.Model(&schemas.WebhookLog{}).Count(&total)
	}

	if result.Error != nil {
		return nil, result.Error
	}

	if totalRes.Error != nil {
		return nil, totalRes.Error
	}

	paginationClone := pagination
	paginationClone.Total = total

	responseWebhookLogs := []*model.WebhookLog{}
	for _, w := range webhookLogs {
		responseWebhookLogs = append(responseWebhookLogs, w.AsAPIWebhookLog())
	}
	return &model.WebhookLogs{
		WebhookLogs: responseWebhookLogs,
		Pagination:  paginationClone,
	}, nil
}
