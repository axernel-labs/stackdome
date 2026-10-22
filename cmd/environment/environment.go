package environment

import (
	"context"

	"github.com/axernel-labs/stackdome/config"
	"github.com/axernel-labs/stackdome/pkg/auth"
	"github.com/axernel-labs/stackdome/pkg/clients/turnstile"
	"github.com/axernel-labs/stackdome/pkg/clustermanager"
	"github.com/axernel-labs/stackdome/pkg/computequota"
	"github.com/axernel-labs/stackdome/pkg/db"
	emailpkg "github.com/axernel-labs/stackdome/pkg/email"
	applogger "github.com/axernel-labs/stackdome/pkg/logger"
	"github.com/axernel-labs/stackdome/pkg/observability"
	"github.com/axernel-labs/stackdome/pkg/resourceaccess"
	"github.com/axernel-labs/stackdome/pkg/services"
	"github.com/axernel-labs/stackdome/pkg/signupprotection"
	"github.com/axernel-labs/stackdome/pkg/stores"
	"github.com/axernel-labs/stackdome/pkg/worker/workermanager"
	"github.com/openshift-online/ocm-sdk-go/leadership"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type EnvImpl interface {
	Init(context.Context) error
	InitDatabase(context.Context) error
	Environment() *Env
	Shutdown(context.Context) error
}

type Env struct {
	Name                        string
	Services                    Services
	DBSession                   db.SessionFactory
	Config                      *config.ApplicationConfig
	PlatformConfig              *config.PlatformConfig
	Clients                     Clients
	ClusterManager              clustermanager.ClusterManager
	WorkerManager               workermanager.WorkerManager
	ResourceAccessPolicyManager resourceaccess.ResourceAccessPolicyManager
	PermissionService           auth.PermissionService
	RefreshTokenStore           stores.RefreshTokenStore
	OAuthStateStore             stores.OAuthStateStore
	EmailService                emailpkg.EmailService
	LeadershipFlag              *leadership.Flag
	EncryptionService           services.EncryptionService
	ComputePolicy               computequota.Policy
	PasswordSignupProtection    signupprotection.PasswordSignupProtection
	SignupClientIPResolver      signupprotection.ClientIPResolver
	Logger                      applogger.Logger
	Observability               *observability.Metrics
}

type Clients struct {
	DefaultClusterClient client.Client
	TurnstileVerifier    turnstile.Verifier
}

type Services struct {
	UserService                 services.UserService
	OrganisationService         services.OrganisationService
	ClusterService              services.ClusterService
	StackStorageService         services.StackStorageService
	VolumeService               services.VolumeService
	StackService                services.StackService
	StackResourceService        services.StackResourceService
	ImageBuildService           services.ImageBuildService
	ClusterImageRegistryService services.ImageRegistryService
	StackDomainService          services.StackDomainsService
	OrganisationDomainService   services.OrganisationDomainsService
	NamespaceService            services.NamespaceService
	LoggingService              services.LoggingService
	MetricsService              services.MetricsService
	SecretService               services.SecretService
	CredentialResolver          services.CredentialResolver
	RegistryCredentialService   services.RegistryCredentialService
	GitIntegrationService       services.GitIntegrationService
	GitHubWebhookService        services.GitHubWebhookService
	EncryptionService           services.EncryptionService
	ObjectStoreService          services.ObjectStoreService
	PostgresAddonService        services.PostgresAddonService
	PostgresBackupService       services.PostgresBackupService
	APITokenService             services.APITokenService
	ProjectService              services.ProjectService
	OrgInviteService            services.OrgInviteService
	SignupService               services.SignupService
	StackReleaseService         services.StackReleaseService
	ReleaseEventRecorder        services.ReleaseEventRecorder
	ReferenceService            services.ReferenceService
	StackPreviewConfigService   services.StackPreviewConfigService
	PreviewStackService         services.PreviewStackService
}
