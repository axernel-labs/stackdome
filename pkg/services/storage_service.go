package services

import (
	"context"

	"github.com/axernel-labs/stackdome/pkg/errors"
	"github.com/axernel-labs/stackdome/pkg/models"
	"github.com/axernel-labs/stackdome/pkg/services/clusterresource"
)

type StackStorageService interface {
	Get(ctx context.Context, ID string, userID string) (*models.StackStorage, *errors.ServiceError)
	GetByID(ctx context.Context, ID string) (*models.StackStorage, *errors.ServiceError)
	GetbyWorkspaceName(ctx context.Context, workspaceName string, userID string) (*models.StackStorage, *errors.ServiceError)
	InternalGet(ctx context.Context, ID string) (*models.StackStorage, *errors.ServiceError)
	ListByOrganisationID(ctx context.Context, organisationID string) ([]*models.StackStorage, *errors.ServiceError)
	ListByUserID(ctx context.Context, userID string) ([]*models.StackStorage, *errors.ServiceError)
	Create(ctx context.Context, spec *models.StackStorage, userID string) (*models.StackStorage, *errors.ServiceError)
	Update(ctx context.Context, ID string, userID string, spec *models.StackStorage) (*models.StackStorage, *errors.ServiceError)
	UpdateStatus(ctx context.Context, ID string, status *models.StackStorageStatus) *errors.ServiceError
	InjectClusterResourceService(clusterResourceService clusterresource.StackStorageClusterResourceService)
	ListVolumes(ctx context.Context, workspaceStorageID string, userID string) ([]*models.Volume, *errors.ServiceError)
	Delete(ctx context.Context, ID string, userID string) *errors.ServiceError
	MarkAsSynced(ctx context.Context, userID string, storageID string, volumeID string) *errors.ServiceError
}
