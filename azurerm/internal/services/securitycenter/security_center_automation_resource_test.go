package securitycenter_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/clients"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

func TestAccSecurityCenterAutomation_logicApp(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_security_center_automation", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testAccSecurityCenterAutomationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSecurityCenterAutomation_logicApp(data),
				Check: resource.ComposeTestCheckFunc(
					testAccSecurityCenterAutomationExists(data.ResourceName),
				),
			},
			data.ImportStep("action.0.trigger_url"), // trigger_url needs to be ignored
		},
	})
}

func TestAccSecurityCenterAutomation_logAnalytics(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_security_center_automation", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testAccSecurityCenterAutomationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSecurityCenterAutomation_logAnalytics(data),
				Check: resource.ComposeTestCheckFunc(
					testAccSecurityCenterAutomationExists(data.ResourceName),
				),
			},
			data.ImportStep(),
		},
	})
}

func TestAccSecurityCenterAutomation_eventHub(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_security_center_automation", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testAccSecurityCenterAutomationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSecurityCenterAutomation_eventHub(data),
				Check: resource.ComposeTestCheckFunc(
					testAccSecurityCenterAutomationExists(data.ResourceName),
				),
			},
			data.ImportStep("action.0.connection_string"), // trigger_url needs to be ignored
		},
	})
}

func TestAccSecurityCenterAutomation_requiresImport(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_security_center_automation", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testAccSecurityCenterAutomationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSecurityCenterAutomation_logicApp(data),
				Check: resource.ComposeTestCheckFunc(
					testAccSecurityCenterAutomationExists(data.ResourceName),
				),
			},
			data.RequiresImportErrorStep(testAccSecurityCenterAutomation_requiresImport),
		},
	})
}

func TestAccSecurityCenterAutomation_update(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_security_center_automation", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testAccSecurityCenterAutomationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSecurityCenterAutomation_logicApp(data),
				Check: resource.ComposeTestCheckFunc(
					testAccSecurityCenterAutomationExists(data.ResourceName),
				),
			},
			data.ImportStep("action.0.trigger_url"), // trigger_url needs to be ignored
			{
				Config: testAccSecurityCenterAutomation_ruleSingle(data),
				Check: resource.ComposeTestCheckFunc(
					testAccSecurityCenterAutomationExists(data.ResourceName),
				),
			},
			data.ImportStep("action.0.trigger_url"), // trigger_url needs to be ignored
			{
				Config: testAccSecurityCenterAutomation_logicApp(data),
				Check: resource.ComposeTestCheckFunc(
					testAccSecurityCenterAutomationExists(data.ResourceName),
				),
			},
			data.ImportStep("action.0.trigger_url"), // trigger_url needs to be ignored
		},
	})
}

func TestAccSecurityCenterAutomation_ruleSingle(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_security_center_automation", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testAccSecurityCenterAutomationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSecurityCenterAutomation_ruleSingle(data),
				Check: resource.ComposeTestCheckFunc(
					testAccSecurityCenterAutomationCountRules(data.ResourceName, 1, 1),
				),
			},
			data.ImportStep("action.0.trigger_url"), // trigger_url needs to be ignored
		},
	})
}

func TestAccSecurityCenterAutomation_ruleMulti(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_security_center_automation", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testAccSecurityCenterAutomationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSecurityCenterAutomation_ruleMulti(data),
				Check: resource.ComposeTestCheckFunc(
					testAccSecurityCenterAutomationCountRules(data.ResourceName, 1, 3),
				),
			},
			data.ImportStep("action.0.trigger_url"), // trigger_url needs to be ignored
		},
	})
}

func TestAccSecurityCenterAutomation_ruleSetMulti(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_security_center_automation", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testAccSecurityCenterAutomationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSecurityCenterAutomation_ruleSetMulti(data),
				Check: resource.ComposeTestCheckFunc(
					testAccSecurityCenterAutomationCountRules(data.ResourceName, 2, 4),
				),
			},
			data.ImportStep("action.0.trigger_url"), // trigger_url needs to be ignored
		},
	})
}

func TestAccSecurityCenterAutomation_scopeMulti(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_security_center_automation", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testAccSecurityCenterAutomationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSecurityCenterAutomation_scopeMulti(data),
				Check: resource.ComposeTestCheckFunc(
					testAccSecurityCenterAutomationCountScopes(data.ResourceName, 3),
				),
			},
			data.ImportStep("action.0.trigger_url"), // trigger_url needs to be ignored
		},
	})
}

func TestAccSecurityCenterAutomation_actionMulti(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_security_center_automation", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testAccSecurityCenterAutomationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSecurityCenterAutomation_actionMulti(data),
				Check: resource.ComposeTestCheckFunc(
					testAccSecurityCenterAutomationCountActions(data.ResourceName, 2),
				),
			},
			data.ImportStep("action.0.trigger_url", "action.1.trigger_url"), // trigger_url needs to be ignored
		},
	})
}

func TestAccSecurityCenterAutomation_sourceMulti(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_security_center_automation", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testAccSecurityCenterAutomationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSecurityCenterAutomation_sourceMulti(data),
				Check: resource.ComposeTestCheckFunc(
					testAccSecurityCenterAutomationCountSources(data.ResourceName, 3),
					testAccSecurityCenterAutomationCountRules(data.ResourceName, 3, 3),
				),
			},
			data.ImportStep("action.0.trigger_url", "action.1.trigger_url"), // trigger_url needs to be ignored
		},
	})
}

func testAccSecurityCenterAutomationExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.AzureProvider.Meta().(*clients.Client).SecurityCenter.AutomationsClient
		ctx := acceptance.AzureProvider.Meta().(*clients.Client).StopContext

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		name := rs.Primary.Attributes["name"]
		resourceGroup, hasResourceGroup := rs.Primary.Attributes["resource_group_name"]
		if !hasResourceGroup {
			return fmt.Errorf("Bad: no resource group found in state for Security Center automation: %s", name)
		}

		resp, err := client.Get(ctx, resourceGroup, name)
		if err != nil {
			if utils.ResponseWasNotFound(resp.Response) {
				return fmt.Errorf("Bad: Security Center automation %q (resource group: %q) does not exist", name, resourceGroup)
			}

			return fmt.Errorf("Bad: Get on Security Center automation: %+v", err)
		}

		return nil
	}
}

func testAccSecurityCenterAutomationCountScopes(resourceName string, scopeCount int) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.AzureProvider.Meta().(*clients.Client).SecurityCenter.AutomationsClient
		ctx := acceptance.AzureProvider.Meta().(*clients.Client).StopContext

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		name := rs.Primary.Attributes["name"]
		resourceGroup, hasResourceGroup := rs.Primary.Attributes["resource_group_name"]
		if !hasResourceGroup {
			return fmt.Errorf("Bad: no resource group found in state for Security Center automation: %s", name)
		}

		resp, err := client.Get(ctx, resourceGroup, name)
		if err != nil {
			if utils.ResponseWasNotFound(resp.Response) {
				return fmt.Errorf("Bad: Security Center automation %q (resource group: %q) does not exist", name, resourceGroup)
			}

			return fmt.Errorf("Bad: Get on Security Center automation: %+v", err)
		}

		if len(*resp.AutomationProperties.Scopes) != scopeCount {
			return fmt.Errorf("Security Center automation doesn't have required number of scopes: got %d, wanted %d", len(*resp.AutomationProperties.Scopes), scopeCount)
		}

		return nil
	}
}

func testAccSecurityCenterAutomationCountRules(resourceName string, ruleSetCount int, ruleCount int) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.AzureProvider.Meta().(*clients.Client).SecurityCenter.AutomationsClient
		ctx := acceptance.AzureProvider.Meta().(*clients.Client).StopContext

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		name := rs.Primary.Attributes["name"]
		resourceGroup, hasResourceGroup := rs.Primary.Attributes["resource_group_name"]
		if !hasResourceGroup {
			return fmt.Errorf("Bad: no resource group found in state for Security Center automation: %s", name)
		}

		resp, err := client.Get(ctx, resourceGroup, name)
		if err != nil {
			if utils.ResponseWasNotFound(resp.Response) {
				return fmt.Errorf("Bad: Security Center automation %q (resource group: %q) does not exist", name, resourceGroup)
			}

			return fmt.Errorf("Bad: Get on Security Center automation: %+v", err)
		}

		actualRuleSetCount := 0
		actualRuleCount := 0
		for _, source := range *resp.AutomationProperties.Sources {
			actualRuleSetCount += len(*source.RuleSets)
			for _, ruleSet := range *source.RuleSets {
				actualRuleCount += len(*ruleSet.Rules)
			}
		}

		if actualRuleSetCount != ruleSetCount {
			return fmt.Errorf("Security Center automation doesn't have required number of rule sets: got %d, wanted %d", actualRuleSetCount, ruleSetCount)
		}
		if actualRuleCount != ruleCount {
			return fmt.Errorf("Security Center automation doesn't have required number of rules: got %d, wanted %d", actualRuleCount, ruleCount)
		}

		return nil
	}
}

func testAccSecurityCenterAutomationCountActions(resourceName string, actionCount int) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.AzureProvider.Meta().(*clients.Client).SecurityCenter.AutomationsClient
		ctx := acceptance.AzureProvider.Meta().(*clients.Client).StopContext

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		name := rs.Primary.Attributes["name"]
		resourceGroup, hasResourceGroup := rs.Primary.Attributes["resource_group_name"]
		if !hasResourceGroup {
			return fmt.Errorf("Bad: no resource group found in state for Security Center automation: %s", name)
		}

		resp, err := client.Get(ctx, resourceGroup, name)
		if err != nil {
			if utils.ResponseWasNotFound(resp.Response) {
				return fmt.Errorf("Bad: Security Center automation %q (resource group: %q) does not exist", name, resourceGroup)
			}

			return fmt.Errorf("Bad: Get on Security Center automation: %+v", err)
		}

		if len(*resp.AutomationProperties.Actions) != actionCount {
			return fmt.Errorf("Security Center automation doesn't have required number of actions: got %d, wanted %d", len(*resp.AutomationProperties.Actions), actionCount)
		}

		return nil
	}
}

func testAccSecurityCenterAutomationCountSources(resourceName string, sourceCount int) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.AzureProvider.Meta().(*clients.Client).SecurityCenter.AutomationsClient
		ctx := acceptance.AzureProvider.Meta().(*clients.Client).StopContext

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		name := rs.Primary.Attributes["name"]
		resourceGroup, hasResourceGroup := rs.Primary.Attributes["resource_group_name"]
		if !hasResourceGroup {
			return fmt.Errorf("Bad: no resource group found in state for Security Center automation: %s", name)
		}

		resp, err := client.Get(ctx, resourceGroup, name)
		if err != nil {
			if utils.ResponseWasNotFound(resp.Response) {
				return fmt.Errorf("Bad: Security Center automation %q (resource group: %q) does not exist", name, resourceGroup)
			}

			return fmt.Errorf("Bad: Get on Security Center automation: %+v", err)
		}

		if len(*resp.AutomationProperties.Sources) != sourceCount {
			return fmt.Errorf("Security Center automation doesn't have required number of sources: got %d, wanted %d", len(*resp.AutomationProperties.Sources), sourceCount)
		}

		return nil
	}
}

func testAccSecurityCenterAutomationDestroy(s *terraform.State) error {
	client := acceptance.AzureProvider.Meta().(*clients.Client).SecurityCenter.AutomationsClient
	ctx := acceptance.AzureProvider.Meta().(*clients.Client).StopContext

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "azurerm_security_center_automation" {
			continue
		}

		name := rs.Primary.Attributes["name"]
		resourceGroup := rs.Primary.Attributes["resource_group_name"]

		resp, err := client.Get(ctx, resourceGroup, name)
		if err != nil {
			if utils.ResponseWasNotFound(resp.Response) {
				return nil
			}

			return fmt.Errorf("Security Center automation still exists:\n%#v", resp)
		}
	}

	return nil
}

func testAccSecurityCenterAutomation_logicApp(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_logic_app_workflow" "test" {
  name                = "acctestlogicapp-%d"
  location            = "%s"
  resource_group_name = azurerm_resource_group.test.name
}

data "azurerm_client_config" "current" {
}

resource "azurerm_security_center_automation" "test" {
  name                = "acctestautomation-%d"
  location            = "%s"
  resource_group_name = azurerm_resource_group.test.name

  scopes = [
    "/subscriptions/${data.azurerm_client_config.current.subscription_id}"
  ]

  action {
    type        = "LogicApp"
    resource_id = azurerm_logic_app_workflow.test.id
    trigger_url = "https://example.net/this_is_never_validated_by_azure"
  }

  source {
    event_source = "Alerts"
  }
}
`, data.RandomInteger, data.Locations.Primary, data.RandomInteger, data.Locations.Primary, data.RandomInteger, data.Locations.Primary)
}

func testAccSecurityCenterAutomation_logAnalytics(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_log_analytics_workspace" "test" {
  name                = "acctestlogs-%d"
  location            = "%s"
  resource_group_name = azurerm_resource_group.test.name
  sku                 = "Free"
}

data "azurerm_client_config" "current" {
}

resource "azurerm_security_center_automation" "test" {
  name                = "acctestautomation-%d"
  location            = "%s"
  resource_group_name = azurerm_resource_group.test.name

  scopes = [
    "/subscriptions/${data.azurerm_client_config.current.subscription_id}"
  ]

  action {
    type        = "LogAnalytics"
    resource_id = azurerm_log_analytics_workspace.test.id
  }

  source {
    event_source = "Alerts"
  }
}
`, data.RandomInteger, data.Locations.Primary, data.RandomInteger, data.Locations.Primary, data.RandomInteger, data.Locations.Primary)
}

func testAccSecurityCenterAutomation_eventHub(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_eventhub_namespace" "test" {
  name                = "acctesteventhub-%d"
  location            = "%s"
  resource_group_name = azurerm_resource_group.test.name
  sku                 = "Basic"
  capacity            = 1
}

data "azurerm_client_config" "current" {
}

resource "azurerm_security_center_automation" "test" {
  name                = "acctestautomation-%d"
  location            = "%s"
  resource_group_name = azurerm_resource_group.test.name

  scopes = [
    "/subscriptions/${data.azurerm_client_config.current.subscription_id}"
  ]

  action {
    type              = "EventHub"
    resource_id       = azurerm_eventhub_namespace.test.id
    connection_string = azurerm_eventhub_namespace.test.default_primary_connection_string
  }

  source {
    event_source = "Alerts"
  }
}
`, data.RandomInteger, data.Locations.Primary, data.RandomInteger, data.Locations.Primary, data.RandomInteger, data.Locations.Primary)
}

func testAccSecurityCenterAutomation_requiresImport(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

resource "azurerm_security_center_automation" "import" {
  name                = azurerm_security_center_automation.test.name
  location            = azurerm_security_center_automation.test.location
  resource_group_name = azurerm_security_center_automation.test.resource_group_name

  scopes = [
    "/subscriptions/${data.azurerm_client_config.current.subscription_id}"
  ]

  action {
    type        = "LogicApp"
    resource_id = azurerm_logic_app_workflow.test.id
    trigger_url = "https://example.net/this_is_never_validated_by_azure"
  }

  source {
    event_source = "Alerts"
  }
}
`, testAccSecurityCenterAutomation_logicApp(data))
}

func testAccSecurityCenterAutomation_ruleSingle(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_logic_app_workflow" "test" {
  name                = "acctestlogicapp-%d"
  location            = "%s"
  resource_group_name = azurerm_resource_group.test.name
}

data "azurerm_client_config" "current" {
}

resource "azurerm_security_center_automation" "test" {
  name                = "acctestautomation-%d"
  location            = "%s"
  resource_group_name = azurerm_resource_group.test.name

  scopes = [
    "/subscriptions/${data.azurerm_client_config.current.subscription_id}"
  ]

  action {
    type        = "LogicApp"
    resource_id = azurerm_logic_app_workflow.test.id
    trigger_url = "https://example.net/this_is_never_validated_by_azure"
  }

  source {
    event_source = "Alerts"
    rule_set {
      rule {
        property_path  = "properties.metadata.severity"
        operator       = "Equals"
        expected_value = "High"
        property_type  = "String"
      }
    }
  }

  description = "Security Center Automation Acc test"
  tags = {
    Env = "Test"
  }
}
`, data.RandomInteger, data.Locations.Primary, data.RandomInteger, data.Locations.Primary, data.RandomInteger, data.Locations.Primary)
}

func testAccSecurityCenterAutomation_scopeMulti(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_logic_app_workflow" "test" {
  name                = "acctestlogicapp-%d"
  location            = "%s"
  resource_group_name = azurerm_resource_group.test.name
}

data "azurerm_client_config" "current" {
}

resource "azurerm_security_center_automation" "test" {
  name                = "acctestautomation-%d"
  location            = "%s"
  resource_group_name = azurerm_resource_group.test.name

  scopes = [
    "/subscriptions/${data.azurerm_client_config.current.subscription_id}",
    "/subscriptions/${data.azurerm_client_config.current.subscription_id}/resourceGroups/test",
    "/subscriptions/${data.azurerm_client_config.current.subscription_id}/resourceGroups/test2"
  ]

  action {
    type        = "LogicApp"
    resource_id = azurerm_logic_app_workflow.test.id
    trigger_url = "https://example.net/this_is_never_validated_by_azure"
  }

  source {
    event_source = "Alerts"
    rule_set {
      rule {
        property_path  = "properties.metadata.severity"
        operator       = "Equals"
        expected_value = "High"
        property_type  = "String"
      }
    }
  }
}
`, data.RandomInteger, data.Locations.Primary, data.RandomInteger, data.Locations.Primary, data.RandomInteger, data.Locations.Primary)
}

func testAccSecurityCenterAutomation_ruleMulti(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_logic_app_workflow" "test" {
  name                = "acctestlogicapp-%d"
  location            = "%s"
  resource_group_name = azurerm_resource_group.test.name
}

data "azurerm_client_config" "current" {
}

resource "azurerm_security_center_automation" "test" {
  name                = "acctestautomation-%d"
  location            = "%s"
  resource_group_name = azurerm_resource_group.test.name

  scopes = [
    "/subscriptions/${data.azurerm_client_config.current.subscription_id}"
  ]

  action {
    type        = "LogicApp"
    resource_id = azurerm_logic_app_workflow.test.id
    trigger_url = "https://example.net/this_is_never_validated_by_azure"
  }

  source {
    event_source = "Alerts"
    rule_set {
      rule {
        property_path  = "properties.metadata.severity"
        operator       = "Equals"
        expected_value = "Low"
        property_type  = "String"
      }
      rule {
        property_path  = "properties.metadata.severity"
        operator       = "Equals"
        expected_value = "Medium"
        property_type  = "String"
      }
      rule {
        property_path  = "properties.metadata.severity"
        operator       = "Equals"
        expected_value = "High"
        property_type  = "String"
      }
    }
  }
}
`, data.RandomInteger, data.Locations.Primary, data.RandomInteger, data.Locations.Primary, data.RandomInteger, data.Locations.Primary)
}

func testAccSecurityCenterAutomation_ruleSetMulti(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_logic_app_workflow" "test" {
  name                = "acctestlogicapp-%d"
  location            = "%s"
  resource_group_name = azurerm_resource_group.test.name
}

data "azurerm_client_config" "current" {
}

resource "azurerm_security_center_automation" "test" {
  name                = "acctestautomation-%d"
  location            = "%s"
  resource_group_name = azurerm_resource_group.test.name

  scopes = [
    "/subscriptions/${data.azurerm_client_config.current.subscription_id}"
  ]

  action {
    type        = "LogicApp"
    resource_id = azurerm_logic_app_workflow.test.id
    trigger_url = "https://example.net/this_is_never_validated_by_azure"
  }

  source {
    event_source = "Alerts"
    rule_set {
      rule {
        property_path  = "properties.metadata.title"
        operator       = "Equals"
        expected_value = "Tony Iommi"
        property_type  = "String"
      }
      rule {
        property_path  = "properties.metadata.title"
        operator       = "Equals"
        expected_value = "Ozzy Osbourne"
        property_type  = "String"
      }
    }
    rule_set {
      rule {
        property_path  = "properties.metadata.title"
        operator       = "Equals"
        expected_value = "Bill Ward"
        property_type  = "String"
      }
      rule {
        property_path  = "properties.metadata.title"
        operator       = "Equals"
        expected_value = "Geezer Butler"
        property_type  = "String"
      }
    }
  }
}
`, data.RandomInteger, data.Locations.Primary, data.RandomInteger, data.Locations.Primary, data.RandomInteger, data.Locations.Primary)
}

func testAccSecurityCenterAutomation_actionMulti(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_logic_app_workflow" "test" {
  name                = "acctestlogicapp-%d"
  location            = "%s"
  resource_group_name = azurerm_resource_group.test.name
}

resource "azurerm_logic_app_workflow" "test2" {
  name                = "acctestlogicapp2-%d"
  location            = "%s"
  resource_group_name = azurerm_resource_group.test.name
}

data "azurerm_client_config" "current" {
}

resource "azurerm_security_center_automation" "test" {
  name                = "acctestautomation-%d"
  location            = "%s"
  resource_group_name = azurerm_resource_group.test.name

  scopes = [
    "/subscriptions/${data.azurerm_client_config.current.subscription_id}"
  ]

  action {
    type        = "LogicApp"
    resource_id = azurerm_logic_app_workflow.test.id
    trigger_url = "https://example.net/this_is_never_validated_by_azure"
  }

  action {
    type        = "LogicApp"
    resource_id = azurerm_logic_app_workflow.test2.id
    trigger_url = "https://example.net/this_is_also_never_validated_by_azure"
  }

  source {
    event_source = "Alerts"
  }
}
`, data.RandomInteger, data.Locations.Primary, data.RandomInteger, data.Locations.Primary, data.RandomInteger, data.Locations.Primary, data.RandomInteger, data.Locations.Primary)
}

func testAccSecurityCenterAutomation_sourceMulti(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_logic_app_workflow" "test" {
  name                = "acctestlogicapp-%d"
  location            = "%s"
  resource_group_name = azurerm_resource_group.test.name
}

data "azurerm_client_config" "current" {
}

resource "azurerm_security_center_automation" "test" {
  name                = "acctestautomation-%d"
  location            = "%s"
  resource_group_name = azurerm_resource_group.test.name

  scopes = [
    "/subscriptions/${data.azurerm_client_config.current.subscription_id}"
  ]

  action {
    type        = "LogicApp"
    resource_id = azurerm_logic_app_workflow.test.id
    trigger_url = "https://example.net/this_is_never_validated_by_azure"
  }

  source {
    event_source = "Alerts"
    rule_set {
      rule {
        property_path  = "properties.metadata.severity"
        operator       = "Equals"
        expected_value = "Low"
        property_type  = "String"
      }
    }
  }

  source {
    event_source = "Assessments"
    rule_set {
      rule {
        property_path  = "properties.metadata.severity"
        operator       = "Equals"
        expected_value = "Low"
        property_type  = "String"
      }
    }
  }

  source {
    event_source = "SubAssessments"
    rule_set {
      rule {
        property_path  = "properties.metadata.severity"
        operator       = "Equals"
        expected_value = "Low"
        property_type  = "String"
      }
    }
  }
}
`, data.RandomInteger, data.Locations.Primary, data.RandomInteger, data.Locations.Primary, data.RandomInteger, data.Locations.Primary)
}
