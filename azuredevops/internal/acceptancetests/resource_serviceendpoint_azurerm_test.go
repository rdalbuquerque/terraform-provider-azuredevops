//go:build (all || resource_serviceendpoint_azurerm) && !exclude_serviceendpoints
// +build all resource_serviceendpoint_azurerm
// +build !exclude_serviceendpoints

package acceptancetests

import (
	"regexp"
	"strconv"
	"testing"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/microsoft/terraform-provider-azuredevops/azuredevops/internal/acceptancetests/testutils"
)

// validates that an apply followed by another apply (i.e., resource update) will be reflected in AzDO and the
// underlying terraform state.
func TestAccServiceEndpointAzureRm_CreateAndUpdate(t *testing.T) {
	// t.Skip("Skipping test TestAccServiceEndpointAzureRm_CreateAndUpdate: test resource limit")
	projectName := testutils.GenerateResourceName()
	serviceEndpointNameFirst := testutils.GenerateResourceName()
	serviceEndpointNameSecond := testutils.GenerateResourceName()
	serviceprincipalidFirst := uuid.New().String()
	serviceprincipalidSecond := uuid.New().String()
	serviceprincipalkeyFirst := uuid.New().String()
	serviceprincipalkeySecond := uuid.New().String()

	resourceType := "azuredevops_serviceendpoint_azurerm"
	tfSvcEpNode := resourceType + ".serviceendpointrm"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testutils.PreCheck(t, nil) },
		Providers:    testutils.GetProviders(),
		CheckDestroy: testutils.CheckServiceEndpointDestroyed(resourceType),
		Steps: []resource.TestStep{
			{
				Config: testutils.HclServiceEndpointAzureRMResource(projectName, serviceEndpointNameFirst, serviceprincipalidFirst, serviceprincipalkeyFirst),
				Check: resource.ComposeTestCheckFunc(
					testutils.CheckServiceEndpointExistsWithName(tfSvcEpNode, serviceEndpointNameFirst),
					resource.TestCheckResourceAttrSet(tfSvcEpNode, "project_id"),
					resource.TestCheckResourceAttrSet(tfSvcEpNode, "azurerm_spn_tenantid"),
					resource.TestCheckResourceAttr(tfSvcEpNode, "service_endpoint_name", serviceEndpointNameFirst),
					resource.TestCheckResourceAttrSet(tfSvcEpNode, "azurerm_subscription_id"),
					resource.TestCheckResourceAttrSet(tfSvcEpNode, "azurerm_subscription_name"),
					resource.TestCheckResourceAttr(tfSvcEpNode, "credentials.0.serviceprincipalid", serviceprincipalidFirst),
					resource.TestCheckResourceAttrSet(tfSvcEpNode, "credentials.0.serviceprincipalkey_hash"),
					resource.TestCheckResourceAttr(tfSvcEpNode, "credentials.0.serviceprincipalkey", serviceprincipalkeyFirst),
				),
			}, {
				Config: testutils.HclServiceEndpointAzureRMResource(projectName, serviceEndpointNameSecond, serviceprincipalidSecond, serviceprincipalkeySecond),
				Check: resource.ComposeTestCheckFunc(
					testutils.CheckServiceEndpointExistsWithName(tfSvcEpNode, serviceEndpointNameSecond),
					resource.TestCheckResourceAttrSet(tfSvcEpNode, "project_id"),
					resource.TestCheckResourceAttrSet(tfSvcEpNode, "azurerm_spn_tenantid"),
					resource.TestCheckResourceAttrSet(tfSvcEpNode, "azurerm_subscription_id"),
					resource.TestCheckResourceAttrSet(tfSvcEpNode, "azurerm_subscription_name"),
					resource.TestCheckResourceAttr(tfSvcEpNode, "service_endpoint_name", serviceEndpointNameSecond),
					resource.TestCheckResourceAttr(tfSvcEpNode, "credentials.0.serviceprincipalid", serviceprincipalidSecond),
					resource.TestCheckResourceAttrSet(tfSvcEpNode, "credentials.0.serviceprincipalkey_hash"),
					resource.TestCheckResourceAttr(tfSvcEpNode, "credentials.0.serviceprincipalkey", serviceprincipalkeySecond),
				),
			},
		},
	})
}

func TestAccServiceEndpointAzureRm_CreateAndUpdate_WithValidate(t *testing.T) {
	// t.Skip("Skipping test TestAccServiceEndpointAzureRm_CreateAndUpdate: test resource limit")
	projectName := testutils.GenerateResourceName()
	serviceEndpointNameFirst := testutils.GenerateResourceName()
	serviceprincipalidFirst := uuid.New().String()
	serviceprincipalkeyFirst := uuid.New().String()
	validateFirst := false
	validateSecond := true

	resourceType := "azuredevops_serviceendpoint_azurerm"
	tfSvcEpNode := resourceType + ".serviceendpointrm"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testutils.PreCheck(t, nil) },
		Providers:    testutils.GetProviders(),
		CheckDestroy: testutils.CheckServiceEndpointDestroyed(resourceType),
		Steps: []resource.TestStep{
			{
				Config: testutils.HclServiceEndpointAzureRMResourceWithValidate(projectName, serviceEndpointNameFirst, serviceprincipalidFirst, serviceprincipalkeyFirst, validateFirst),
				Check: resource.ComposeTestCheckFunc(
					testutils.CheckServiceEndpointExistsWithName(tfSvcEpNode, serviceEndpointNameFirst),
					resource.TestCheckResourceAttrSet(tfSvcEpNode, "project_id"),
					resource.TestCheckResourceAttrSet(tfSvcEpNode, "azurerm_spn_tenantid"),
					resource.TestCheckResourceAttr(tfSvcEpNode, "service_endpoint_name", serviceEndpointNameFirst),
					resource.TestCheckResourceAttrSet(tfSvcEpNode, "azurerm_subscription_id"),
					resource.TestCheckResourceAttrSet(tfSvcEpNode, "azurerm_subscription_name"),
					resource.TestCheckResourceAttr(tfSvcEpNode, "credentials.0.serviceprincipalid", serviceprincipalidFirst),
					resource.TestCheckResourceAttrSet(tfSvcEpNode, "credentials.0.serviceprincipalkey_hash"),
					resource.TestCheckResourceAttr(tfSvcEpNode, "credentials.0.serviceprincipalkey", serviceprincipalkeyFirst),
					resource.TestCheckResourceAttr(tfSvcEpNode, "validate", strconv.FormatBool(validateFirst)),
				),
			}, {
				Config:      testutils.HclServiceEndpointAzureRMResourceWithValidate(projectName, serviceEndpointNameFirst, serviceprincipalidFirst, serviceprincipalkeyFirst, validateSecond),
				ExpectError: regexp.MustCompile("Failed to obtain the Json Web Token"),
			},
		},
	})
}

func TestAccServiceEndpointAzureRm_Create_WithValidate(t *testing.T) {
	// t.Skip("Skipping test TestAccServiceEndpointAzureRm_CreateAndUpdate: test resource limit")
	projectName := testutils.GenerateResourceName()
	serviceEndpointName := testutils.GenerateResourceName()
	serviceprincipalid := uuid.New().String()
	serviceprincipalkey := uuid.New().String()
	validate := true

	resourceType := "azuredevops_serviceendpoint_azurerm"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testutils.PreCheck(t, nil) },
		Providers:    testutils.GetProviders(),
		CheckDestroy: testutils.CheckServiceEndpointDestroyed(resourceType),
		Steps: []resource.TestStep{
			{
				Config:      testutils.HclServiceEndpointAzureRMResourceWithValidate(projectName, serviceEndpointName, serviceprincipalid, serviceprincipalkey, validate),
				ExpectError: regexp.MustCompile("Failed to obtain the Json Web Token"),
			},
		},
	})
}

func TestAccServiceEndpointAzureRm_MgmtGrpCreateAndUpdate(t *testing.T) {
	t.Skip("Skipping test TestAccServiceEndpointAzureRm_MgmtGrpCreateAndUpdate: test resource limit")
	projectName := testutils.GenerateResourceName()
	serviceEndpointName := testutils.GenerateResourceName()
	serviceprincipalid := uuid.New().String()
	serviceprincipalkey := uuid.New().String()

	resourceType := "azuredevops_serviceendpoint_azurerm"
	tfSvcEpNode := resourceType + ".serviceendpointrm"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testutils.PreCheck(t, nil) },
		Providers:    testutils.GetProviders(),
		CheckDestroy: testutils.CheckServiceEndpointDestroyed(resourceType),
		Steps: []resource.TestStep{
			{
				Config: testutils.HclServiceEndpointAzureRMResourceWithMG(projectName, serviceEndpointName, serviceprincipalid, serviceprincipalkey),
				Check: resource.ComposeTestCheckFunc(
					testutils.CheckServiceEndpointExistsWithName(tfSvcEpNode, serviceEndpointName),
					resource.TestCheckResourceAttrSet(tfSvcEpNode, "project_id"),
					resource.TestCheckResourceAttrSet(tfSvcEpNode, "azurerm_spn_tenantid"),
					resource.TestCheckResourceAttr(tfSvcEpNode, "service_endpoint_name", serviceEndpointName),
					resource.TestCheckResourceAttrSet(tfSvcEpNode, "azurerm_management_group_id"),
					resource.TestCheckResourceAttrSet(tfSvcEpNode, "azurerm_management_group_name"),
					resource.TestCheckResourceAttr(tfSvcEpNode, "credentials.0.serviceprincipalid", serviceprincipalid),
					resource.TestCheckResourceAttrSet(tfSvcEpNode, "credentials.0.serviceprincipalkey_hash"),
					resource.TestCheckResourceAttr(tfSvcEpNode, "credentials.0.serviceprincipalkey", serviceprincipalkey),
				),
			},
		},
	})
}

func TestAccServiceEndpointAzureRm_AutomaticCreateAndUpdate(t *testing.T) {
	t.Skip("Skipping test TestAccServiceEndpointAzureRm_AutomaticCreateAndUpdate: test resource limit")

	projectName := testutils.GenerateResourceName()
	serviceEndpointName := testutils.GenerateResourceName()

	resourceType := "azuredevops_serviceendpoint_azurerm"
	tfSvcEpNode := resourceType + ".serviceendpointrm"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testutils.PreCheck(t, nil) },
		Providers:    testutils.GetProviders(),
		CheckDestroy: testutils.CheckServiceEndpointDestroyed(resourceType),
		Steps: []resource.TestStep{
			{
				Config: testutils.HclServiceEndpointAzureRMAutomaticResourceWithProject(projectName, serviceEndpointName),
				Check: resource.ComposeTestCheckFunc(
					testutils.CheckServiceEndpointExistsWithName(tfSvcEpNode, serviceEndpointName),
					resource.TestCheckResourceAttrSet(tfSvcEpNode, "project_id"),
					resource.TestCheckResourceAttrSet(tfSvcEpNode, "azurerm_spn_tenantid"),
					resource.TestCheckResourceAttrSet(tfSvcEpNode, "azurerm_subscription_id"),
					resource.TestCheckResourceAttrSet(tfSvcEpNode, "azurerm_subscription_name"),
					resource.TestCheckResourceAttr(tfSvcEpNode, "service_endpoint_name", serviceEndpointName),
					resource.TestCheckNoResourceAttr(tfSvcEpNode, "credentials.0"),
				),
			},
			{
				Config: testutils.HclServiceEndpointAzureRMAutomaticResourceWithProject(projectName, serviceEndpointName),
				Check: resource.ComposeTestCheckFunc(
					testutils.CheckServiceEndpointExistsWithName(tfSvcEpNode, serviceEndpointName),
					resource.TestCheckResourceAttrSet(tfSvcEpNode, "project_id"),
					resource.TestCheckResourceAttrSet(tfSvcEpNode, "azurerm_spn_tenantid"),
					resource.TestCheckResourceAttrSet(tfSvcEpNode, "azurerm_subscription_id"),
					resource.TestCheckResourceAttrSet(tfSvcEpNode, "azurerm_subscription_name"),
					resource.TestCheckResourceAttr(tfSvcEpNode, "service_endpoint_name", serviceEndpointName),
					resource.TestCheckNoResourceAttr(tfSvcEpNode, "credentials.0"),
				),
			},
		},
	})
}
