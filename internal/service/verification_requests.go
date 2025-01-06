package service

import (
	"context"
	"fmt"

	"github.com/authorizerdev/authorizer/internal/graph/model"
	"github.com/authorizerdev/authorizer/internal/utils"
)

// VerificationRequests is used to get all verification requests
// Permission: authorizer:admin
func (s *service) VerificationRequests(ctx context.Context, params *model.PaginatedInput) (*model.VerificationRequests, error) {
	log := s.Log.With().Str("func", "VerificationRequests").Logger()
	gc, err := utils.GinContextFromContext(ctx)
	if err != nil {
		log.Debug().Err(err).Msg("Failed to get GinContext")
		return nil, err
	}
	if !s.TokenProvider.IsSuperAdmin(gc) {
		log.Debug().Msg("Not logged in as super admin")
		return nil, fmt.Errorf("unauthorized")
	}

	pagination := utils.GetPagination(params)
	res, err := s.StorageProvider.ListVerificationRequests(ctx, pagination)
	if err != nil {
		log.Debug().Err(err).Msg("failed ListVerificationRequests")
		return nil, err
	}

	return res, nil
}
