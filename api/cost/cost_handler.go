package cost

import (
	costModels "github.com/equinor/radix-cost-allocation-api/api/cost/models"
	"github.com/equinor/radix-cost-allocation-api/api/utils"
	"github.com/equinor/radix-cost-allocation-api/models"
	v1 "github.com/equinor/radix-operator/pkg/apis/radix/v1"
	crdUtils "github.com/equinor/radix-operator/pkg/apis/utils"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

// CostHandler Instance variables
type CostHandler struct {
	accounts models.Accounts
}

// Init Constructor
func Init(accounts models.Accounts) CostHandler {
	return CostHandler{
		accounts: accounts,
	}
}

func (costHandler CostHandler) getUserAccount() models.Account {
	return costHandler.accounts.UserAccount
}

func (costHandler CostHandler) getServiceAccount() models.Account {
	return costHandler.accounts.ServiceAccount
}

// GetTotalCost handler for GetTotalCost
func (costHandler CostHandler) GetTotalCost(appName string) (*costModels.Cost, error) {
	radixRegistration, err := costHandler.getServiceAccount().RadixClient.RadixV1().RadixRegistrations().Get(appName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	applicationRegistrationBuilder := NewBuilder()
	applicationRegistration := applicationRegistrationBuilder.
		withRadixRegistration(radixRegistration).
		Build()

	return &costModels.Cost{
		ApplicationName:    applicationRegistration.Name,
		ApplicationOwner:   applicationRegistration.Owner,
		ApplicationCreator: applicationRegistration.Creator}, nil
}

// ApplicationBuilder Handles construction of DTO
type ApplicationBuilder interface {
	withName(name string) ApplicationBuilder
	withOwner(owner string) ApplicationBuilder
	withCreator(creator string) ApplicationBuilder
	withWBS(string) ApplicationBuilder
	withRadixRegistration(*v1.RadixRegistration) ApplicationBuilder
	Build() costModels.ApplicationRegistration
}

type applicationBuilder struct {
	name         string
	owner        string
	creator      string
	repository   string
	sharedSecret string
	adGroups     []string
	publicKey    string
	privateKey   string
	cloneURL     string
	machineUser  bool
	wbs          string
}

func (rb *applicationBuilder) withAppRegistration(appRegistration *costModels.ApplicationRegistration) ApplicationBuilder {
	rb.withName(appRegistration.Name)
	rb.withRepository(appRegistration.Repository)
	rb.withSharedSecret(appRegistration.SharedSecret)
	rb.withAdGroups(appRegistration.AdGroups)
	rb.withPublicKey(appRegistration.PublicKey)
	rb.withPrivateKey(appRegistration.PrivateKey)
	rb.withOwner(appRegistration.Owner)
	rb.withWBS(appRegistration.WBS)
	return rb
}

func (rb *applicationBuilder) withRadixRegistration(radixRegistration *v1.RadixRegistration) ApplicationBuilder {
	rb.withName(radixRegistration.Name)
	rb.withCloneURL(radixRegistration.Spec.CloneURL)
	rb.withSharedSecret(radixRegistration.Spec.SharedSecret)
	rb.withAdGroups(radixRegistration.Spec.AdGroups)
	rb.withPublicKey(radixRegistration.Spec.DeployKeyPublic)
	rb.withOwner(radixRegistration.Spec.Owner)
	rb.withCreator(radixRegistration.Spec.Creator)
	rb.withMachineUser(radixRegistration.Spec.MachineUser)
	rb.withWBS(radixRegistration.Spec.WBS)

	// Private part of key should never be returned
	return rb
}

func (rb *applicationBuilder) withName(name string) ApplicationBuilder {
	rb.name = name
	return rb
}

func (rb *applicationBuilder) withOwner(owner string) ApplicationBuilder {
	rb.owner = owner
	return rb
}

func (rb *applicationBuilder) withCreator(creator string) ApplicationBuilder {
	rb.creator = creator
	return rb
}

func (rb *applicationBuilder) withRepository(repository string) ApplicationBuilder {
	rb.repository = repository
	return rb
}

func (rb *applicationBuilder) withCloneURL(cloneURL string) ApplicationBuilder {
	rb.cloneURL = cloneURL
	return rb
}

func (rb *applicationBuilder) withSharedSecret(sharedSecret string) ApplicationBuilder {
	rb.sharedSecret = sharedSecret
	return rb
}

func (rb *applicationBuilder) withAdGroups(adGroups []string) ApplicationBuilder {
	rb.adGroups = adGroups
	return rb
}

func (rb *applicationBuilder) withPublicKey(publicKey string) ApplicationBuilder {
	rb.publicKey = strings.TrimSuffix(publicKey, "\n")
	return rb
}

func (rb *applicationBuilder) withPrivateKey(privateKey string) ApplicationBuilder {
	rb.privateKey = strings.TrimSuffix(privateKey, "\n")
	return rb
}

func (rb *applicationBuilder) withDeployKey(deploykey *utils.DeployKey) ApplicationBuilder {
	if deploykey != nil {
		rb.publicKey = deploykey.PublicKey
		rb.privateKey = deploykey.PrivateKey
	}

	return rb
}

func (rb *applicationBuilder) withMachineUser(machineUser bool) ApplicationBuilder {
	rb.machineUser = machineUser
	return rb
}

func (rb *applicationBuilder) withWBS(wbs string) ApplicationBuilder {
	rb.wbs = wbs
	return rb
}

func (rb *applicationBuilder) Build() costModels.ApplicationRegistration {
	repository := rb.repository
	if repository == "" {
		repository = crdUtils.GetGithubRepositoryURLFromCloneURL(rb.cloneURL)
	}

	return costModels.ApplicationRegistration{
		Name:         rb.name,
		Repository:   repository,
		SharedSecret: rb.sharedSecret,
		AdGroups:     rb.adGroups,
		PublicKey:    rb.publicKey,
		PrivateKey:   rb.privateKey,
		Owner:        rb.owner,
		Creator:      rb.creator,
		MachineUser:  rb.machineUser,
		WBS:          rb.wbs,
	}
}

func (rb *applicationBuilder) BuildRR() (*v1.RadixRegistration, error) {
	builder := crdUtils.NewRegistrationBuilder()

	radixRegistration := builder.
		WithPublicKey(rb.publicKey).
		WithPrivateKey(rb.privateKey).
		WithName(rb.name).
		WithRepository(rb.repository).
		WithSharedSecret(rb.sharedSecret).
		WithAdGroups(rb.adGroups).
		WithOwner(rb.owner).
		WithCreator(rb.creator).
		WithMachineUser(rb.machineUser).
		WithWBS(rb.wbs).
		BuildRR()

	return radixRegistration, nil
}

// NewBuilder Constructor for application builder
func NewBuilder() ApplicationBuilder {
	return &applicationBuilder{}
}

// AnApplicationRegistration Constructor for application builder with test values
func AnApplicationRegistration() ApplicationBuilder {
	return &applicationBuilder{
		name:    "my-app",
		owner:   "a_test_user@equinor.com",
		creator: "a_test_user@equinor.com",
	}
}
