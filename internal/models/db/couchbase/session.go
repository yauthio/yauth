package couchbase

import (
	"context"
	"time"

	"github.com/couchbase/gocb/v2"
	"github.com/google/uuid"

	"github.com/authorizerdev/authorizer/internal/models/schemas"
)

// AddSession to save session information in database
func (p *provider) AddSession(ctx context.Context, session *schemas.Session) error {
	if session.ID == "" {
		session.ID = uuid.New().String()
		session.CreatedAt = time.Now().Unix()

		session.UpdatedAt = time.Now().Unix()
	}
	insertOpt := gocb.InsertOptions{
		Context: ctx,
	}
	_, err := p.db.Collection(schemas.Collections.Session).Insert(session.ID, session, &insertOpt)
	if err != nil {
		return err
	}
	return nil
}

// DeleteSession to delete session information from database
func (p *provider) DeleteSession(ctx context.Context, userId string) error {
	return nil
}
