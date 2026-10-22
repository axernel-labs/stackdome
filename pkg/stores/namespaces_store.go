package stores

import (
	"context"

	"github.com/axernel-labs/stackdome/pkg/errors"
	"github.com/axernel-labs/stackdome/pkg/models"
)

type NamespacesStore interface {
	CreateWithTx(ctx context.Context, ns *models.Namespace) (*models.Namespace, *errors.ServiceError)
	Create(ctx context.Context, ns *models.Namespace) (*models.Namespace, *errors.ServiceError)
	Get(ctx context.Context, id string) (*models.Namespace, *errors.ServiceError)
	Delete(ctx context.Context, id string) *errors.ServiceError
	DeleteWithTx(ctx context.Context, id string) *errors.ServiceError
	Update(ctx context.Context, id string, ns *models.Namespace) (*models.Namespace, *errors.ServiceError)
	ListByOrganisation(ctx context.Context, organisationID string) ([]*models.Namespace, *errors.ServiceError)
	ListByStack(ctx context.Context, stackID string) ([]*models.Namespace, *errors.ServiceError)
}
