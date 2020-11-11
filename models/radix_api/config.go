package radix_api

import (
	"encoding/json"
	"os"
	"path"

	jsonutils "github.com/equinor/radix-cost-allocation-api/utils/json"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	"k8s.io/client-go/util/homedir"
)

const (
	ContextProdction   = "production"
	ContextPlayground  = "playground"
	ContextDevelopment = "development"

	recommendedHomeDir  = ".radix"
	recommendedFileName = "config"

	clientID    = "ed6cb804-8193-4e55-9d3d-8b88688482b3"
	tenantID    = "3aa4a235-b6e2-48d5-9195-7fcf05b459b0"
	apiServerID = "a593a59c-8f76-490e-937b-a90779039a90"

	defaultContext = ContextProdction

	cfgContext      = "context"
	cfgClientID     = "client-id"
	cfgTenantID     = "tenant-id"
	cfgAccessToken  = "access-token"
	cfgRefreshToken = "refresh-token"
	cfgExpiresIn    = "expires-in"
	cfgExpiresOn    = "expires-on"
	cfgEnvironment  = "environment"
	cfgApiserverID  = "apiserver-id"
)

var (
	RecommendedConfigDir = path.Join(homedir.HomeDir(), recommendedHomeDir)
	RecommendedHomeFile  = path.Join(RecommendedConfigDir, recommendedFileName)
	ValidContexts        = []string{ContextProdction, ContextPlayground, ContextDevelopment}
)

type RadixConfig struct {
	CustomConfig  *CustomConfig  `json:"customConfig"`
	SessionConfig *SessionConfig `json:"sessionConfig"`
}

type CustomConfig struct {
	Context string `json:"Context"`
}

type SessionConfig struct {
	ClientID     string      `json:"clientID"`
	TenantID     string      `json:"tenantID"`
	APIServerID  string      `json:"apiServerID"`
	RefreshToken string      `json:"refreshToken"`
	AccessToken  string      `json:"accessToken"`
	ExpiresIn    json.Number `json:"expiresIn"`
	ExpiresOn    json.Number `json:"expiresOn"`
	Environment  string      `json:"environment"`
}

type RadixConfigAccess struct {
}

func (c RadixConfigAccess) GetStartingConfig() *clientcmdapi.AuthProviderConfig {
	var radixConfig *RadixConfig
	if _, err := os.Stat(RecommendedHomeFile); err == nil {
		radixConfig = &RadixConfig{}
		jsonutils.Load(RecommendedHomeFile, radixConfig)
	} else {
		radixConfig = &RadixConfig{
			CustomConfig: &CustomConfig{
				Context: defaultContext,
			},
			SessionConfig: &SessionConfig{
				ClientID:    clientID,
				TenantID:    tenantID,
				APIServerID: apiServerID,
			},
		}
	}

	authProvider := &clientcmdapi.AuthProviderConfig{
		Name:   "azure",
		Config: toMap(radixConfig),
	}

	return authProvider
}

func toMap(radixConfig *RadixConfig) map[string]string {
	config := make(map[string]string)
	if radixConfig.CustomConfig != nil {
		config[cfgContext] = radixConfig.CustomConfig.Context
	}

	config[cfgClientID] = radixConfig.SessionConfig.ClientID
	config[cfgTenantID] = radixConfig.SessionConfig.TenantID
	config[cfgApiserverID] = radixConfig.SessionConfig.APIServerID
	config[cfgRefreshToken] = radixConfig.SessionConfig.RefreshToken
	config[cfgAccessToken] = radixConfig.SessionConfig.AccessToken
	config[cfgExpiresIn] = radixConfig.SessionConfig.ExpiresIn.String()
	config[cfgExpiresOn] = radixConfig.SessionConfig.ExpiresOn.String()
	config[cfgEnvironment] = radixConfig.SessionConfig.Environment
	return config
}
