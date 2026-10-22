package bootstrap

import (
	"context"
	"fmt"

	"github.com/axernel-labs/stackdome/config"
	"github.com/axernel-labs/stackdome/pkg/auth"
	"github.com/axernel-labs/stackdome/pkg/errors"
	"github.com/axernel-labs/stackdome/pkg/logger"
	"github.com/axernel-labs/stackdome/pkg/models"
	"github.com/axernel-labs/stackdome/pkg/services"
)

type Spec struct {
	OrganisationService services.OrganisationService
	ClusterService      services.ClusterService
	PlatformConfig      *config.PlatformConfig
	ClusterConfig       *config.ClusterConfig
	Logger              logger.Logger
}

type Service struct {
	organisationService services.OrganisationService
	clusterService      services.ClusterService
	platformCfg         *config.PlatformConfig
	clusterCfg          *config.ClusterConfig
	logger              logger.Logger
}

func NewService(spec Spec) *Service {
	return &Service{
		organisationService: spec.OrganisationService,
		clusterService:      spec.ClusterService,
		platformCfg:         spec.PlatformConfig,
		clusterCfg:          spec.ClusterConfig,
		logger:              spec.Logger,
	}
}

// Run provisions the platform org and shared-compute cluster, then conditionally
// reconciles platform wildcard TLS.
// The platform org is infrastructure-only: no users, no projects.
func (s *Service) Run(ctx context.Context) error {
	if !s.clusterCfg.IsSet() {
		s.logger.Info(ctx, "no shared-compute cluster configured; skipping shared-compute bootstrap")
		return nil
	}

	org, err := s.ensurePlatformOrg(ctx)
	if err != nil {
		return err
	}

	systemIdentity := &auth.Identity{IsSystem: true, OrgID: org.ID}
	if s.platformCfg.PlatformTLSEnabled {
		systemIdentity.ContactEmail = s.platformCfg.Email
	}
	sysCtx := auth.SetIdentityInContext(ctx, systemIdentity)

	cluster, cErr := s.clusterService.InternalUpsertSharedComputeCluster(sysCtx, &models.Cluster{
		Name:           models.SharedComputeClusterName,
		OrganisationID: org.ID,
		SharedCompute:  true,
		ClusterURL:     s.clusterCfg.ClusterURL,
		ClusterCAData:  s.clusterCfg.ClusterCAData,
		Token:          s.clusterCfg.Token,
	})
	if cErr != nil {
		return fmt.Errorf("failed to upsert shared-compute cluster: %w", cErr)
	}

	if s.platformCfg.PlatformTLSEnabled {
		if err := s.clusterService.InternalEnsurePlatformWildcardTLS(sysCtx, cluster, s.platformCfg); err != nil {
			return fmt.Errorf("failed to create or update platform wildcard TLS: %w", err)
		}
	}

	return nil
}

func (s *Service) ensurePlatformOrg(ctx context.Context) (*models.Organisation, error) {
	org, gErr := s.organisationService.InternalGetPlatformOrg(ctx)
	if gErr == nil {
		return org, nil
	}
	if gErr.Code != errors.ErrorNotFound {
		return nil, fmt.Errorf("failed to look up platform org: %w", gErr)
	}

	created, cErr := s.organisationService.InternalCreate(ctx, &models.Organisation{
		Name:     models.PlatformOrganisationName,
		Platform: true,
	})
	if cErr != nil {
		return nil, fmt.Errorf("failed to create platform org: %w", cErr)
	}
	return created, nil
}
