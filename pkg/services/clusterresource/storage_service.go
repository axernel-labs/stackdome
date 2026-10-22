package clusterresource

import (
	"context"

	"github.com/axernel-labs/stackdome/pkg/models"
)

type StackStorageClusterResourceService interface {
	UpsertStorageInCluster(ctx context.Context, stackStorage *models.StackStorage) *ClusterResourceError
	DeleteStorageInCluster(ctx context.Context, stackStorage *models.StackStorage) *ClusterResourceError
}
