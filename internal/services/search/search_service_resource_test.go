package search_test

import (
	"context"
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/go-azure-sdk/resource-manager/search/2022-09-01/services"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azurerm/internal/clients"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/utils"
)

type SearchServiceResource struct{}

func TestAccSearchService_basicStandard(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_search_service", "test")
	r := SearchServiceResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data, "standard"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccSearchService_basicFree(t *testing.T) {
	// Regression test case for issue #10151
	data := acceptance.BuildTestData(t, "azurerm_search_service", "test")
	r := SearchServiceResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data, "free"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccSearchService_basicBasic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_search_service", "test")
	r := SearchServiceResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data, "basic"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccSearchService_requiresImport(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_search_service", "test")
	r := SearchServiceResource{}
	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data, "standard"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.RequiresImportErrorStep(r.requiresImport),
	})
}

func TestAccSearchService_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_search_service", "test")
	r := SearchServiceResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.complete(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccSearchService_update(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_search_service", "test")
	r := SearchServiceResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data, "standard"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.complete(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccSearchService_ipRules(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_search_service", "test")
	r := SearchServiceResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data, "standard"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.ipRules(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.basic(data, "standard"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccSearchService_identity(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_search_service", "test")
	r := SearchServiceResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data, "standard"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.identity(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.basic(data, "standard"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccSearchService_hostingMode(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_search_service", "test")
	r := SearchServiceResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.hostingMode(data, "standard3"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccSearchService_hostingModeInvalidSKU(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_search_service", "test")
	r := SearchServiceResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config:      r.hostingMode(data, "standard2"),
			Check:       acceptance.ComposeTestCheckFunc(),
			ExpectError: regexp.MustCompile("can only be defined if"),
		},
	})
}

func TestAccSearchService_partitionCountInvalidBySku(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_search_service", "test")
	r := SearchServiceResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config:      r.partitionCount(data, "basic", 3),
			Check:       acceptance.ComposeTestCheckFunc(),
			ExpectError: regexp.MustCompile("values greater than 1 cannot be set"),
		},
	})
}

func TestAccSearchService_partitionCountInvalidByInput(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_search_service", "test")
	r := SearchServiceResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config:      r.partitionCount(data, "standard", 5),
			Check:       acceptance.ComposeTestCheckFunc(),
			ExpectError: regexp.MustCompile("must be 1"),
		},
	})
}

func TestAccSearchService_replicaCount(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_search_service", "test")
	r := SearchServiceResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config:      r.replicaCount(data, "basic", 3),
			Check:       acceptance.ComposeTestCheckFunc(),
			ExpectError: regexp.MustCompile("cannot be greater than 1 for the"),
		},
	})
}

func TestAccSearchService_replicaCountInvalid(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_search_service", "test")
	// NOTE: The default quota for the 'free' sku is 1
	// but it is ok to use it here since it will never be created since
	// it will fail validation due to the replica count defined in the
	// test configuration...
	r := SearchServiceResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config:      r.replicaCount(data, "free", 2),
			Check:       acceptance.ComposeTestCheckFunc(),
			ExpectError: regexp.MustCompile("cannot be greater than 1 for the"),
		},
	})
}

func TestAccSearchService_cmkEnforcement(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_search_service", "test")
	r := SearchServiceResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.cmkEnforcement(data, true),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccSearchService_cmkEnforcementUpdate(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_search_service", "test")
	r := SearchServiceResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			// This should create a Search Service with the default 'cmk_enforcement_enabled' value of 'false'...
			Config: r.basic(data, "standard"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.cmkEnforcement(data, false),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.cmkEnforcement(data, true),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccSearchService_apiAccessControlRbacError(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_search_service", "test")
	r := SearchServiceResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config:      r.apiAccessControlBoth(data, true, "http401WithBearerChallenge"),
			Check:       acceptance.ComposeTestCheckFunc(),
			ExpectError: regexp.MustCompile("cannot be defined"),
		},
	})
}

func TestAccSearchService_apiAccessControlUpdate(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_search_service", "test")
	r := SearchServiceResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data, "standard"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.apiAccessControlApiKeysOrRBAC(data, false),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.apiAccessControlApiKeysOrRBAC(data, true),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.apiAccessControlBoth(data, false, "http401WithBearerChallenge"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.apiAccessControlBoth(data, false, "http403"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.basic(data, "standard"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func (t SearchServiceResource) Exists(ctx context.Context, clients *clients.Client, state *pluginsdk.InstanceState) (*bool, error) {
	id, err := services.ParseSearchServiceID(state.ID)
	if err != nil {
		return nil, err
	}

	resp, err := clients.Search.ServicesClient.Get(ctx, *id, services.GetOperationOptions{})
	if err != nil {
		return nil, fmt.Errorf("%s was not found: %+v", *id, err)
	}

	return utils.Bool(resp.Model != nil), nil
}

func (SearchServiceResource) template(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azurerm_resource_group" "test" {
  name     = "acctestRG-search-%d"
  location = "%s"
}
`, data.RandomInteger, data.Locations.Primary)
}

func (r SearchServiceResource) basic(data acceptance.TestData, sku string) string {
	template := r.template(data)
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

%s

resource "azurerm_search_service" "test" {
  name                = "acctestsearchservice%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  sku                 = "%s"

  tags = {
    environment = "staging"
  }
}
`, template, data.RandomInteger, sku)
}

func (r SearchServiceResource) requiresImport(data acceptance.TestData) string {
	template := r.basic(data, "standard")
	return fmt.Sprintf(`
%s

resource "azurerm_search_service" "import" {
  name                = azurerm_search_service.test.name
  resource_group_name = azurerm_search_service.test.resource_group_name
  location            = azurerm_search_service.test.location
  sku                 = azurerm_search_service.test.sku

  tags = {
    environment = "staging"
  }
}
`, template)
}

func (r SearchServiceResource) complete(data acceptance.TestData) string {
	template := r.template(data)
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

%s

resource "azurerm_search_service" "test" {
  name                = "acctestsearchservice%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  sku                 = "standard"
  replica_count       = 2
  partition_count     = 3

  public_network_access_enabled = false
  cmk_enforcement_enabled       = false

  tags = {
    environment = "Production"
    residential = "Area"
  }
}
`, template, data.RandomInteger)
}

func (r SearchServiceResource) ipRules(data acceptance.TestData) string {
	template := r.template(data)
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

%s

resource "azurerm_search_service" "test" {
  name                = "acctestsearchservice%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  sku                 = "standard"

  allowed_ips = ["168.1.5.65", "1.2.3.0/24"]

  tags = {
    environment = "staging"
  }
}
`, template, data.RandomInteger)
}

func (r SearchServiceResource) identity(data acceptance.TestData) string {
	template := r.template(data)
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

%s

resource "azurerm_search_service" "test" {
  name                = "acctestsearchservice%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  sku                 = "standard"

  identity {
    type = "SystemAssigned"
  }

  tags = {
    environment = "staging"
  }
}
`, template, data.RandomInteger)
}

func (r SearchServiceResource) hostingMode(data acceptance.TestData, sku string) string {
	template := r.template(data)
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

%s

resource "azurerm_search_service" "test" {
  name                = "acctestsearchservice%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  sku                 = "%s"
  hosting_mode        = "highDensity"

  tags = {
    environment = "staging"
  }
}
`, template, data.RandomInteger, sku)
}

func (r SearchServiceResource) partitionCount(data acceptance.TestData, sku string, count int) string {
	template := r.template(data)
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

%s

resource "azurerm_search_service" "test" {
  name                = "acctestsearchservice%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  sku                 = "%s"
  partition_count     = %d
}
`, template, data.RandomInteger, sku, count)
}

func (r SearchServiceResource) replicaCount(data acceptance.TestData, sku string, count int) string {
	template := r.template(data)
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

%s

resource "azurerm_search_service" "test" {
  name                = "acctestsearchservice%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  sku                 = "%s"
  replica_count       = %d
}
`, template, data.RandomInteger, sku, count)
}

func (r SearchServiceResource) cmkEnforcement(data acceptance.TestData, enforceCmk bool) string {
	template := r.template(data)
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

%s

resource "azurerm_search_service" "test" {
  name                = "acctestsearchservice%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  sku                 = "standard"

  cmk_enforcement_enabled = %t

  tags = {
    environment = "staging"
  }
}
`, template, data.RandomInteger, enforceCmk)
}

func (r SearchServiceResource) apiAccessControlApiKeysOrRBAC(data acceptance.TestData, localAuthenticationDisabled bool) string {
	template := r.template(data)
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

%s

resource "azurerm_search_service" "test" {
  name                = "acctestsearchservice%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  sku                 = "standard"

  local_authentication_disabled = %t

  tags = {
    environment = "staging"
  }
}
`, template, data.RandomInteger, localAuthenticationDisabled)
}

func (r SearchServiceResource) apiAccessControlBoth(data acceptance.TestData, localAuthenticationDisabled bool, authenticationFailureMode string) string {
	template := r.template(data)
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

%s

resource "azurerm_search_service" "test" {
  name                = "acctestsearchservice%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  sku                 = "standard"

  local_authentication_disabled = %t
  authentication_failure_mode   = %s

  tags = {
    environment = "staging"
  }
}
`, template, data.RandomInteger, localAuthenticationDisabled, authenticationFailureMode)
}
