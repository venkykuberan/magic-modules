package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccLoggingBucketConfigFolder_basic(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"folder_name":   "tf-test-" + randString(t, 10),
		"org_id":        getTestOrgFromEnv(t),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccLoggingBucketConfigFolder_basic(context, 30),
			},
			{
				ResourceName:            "google_logging_folder_bucket_config.basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"folder"},
			},
			{
				Config: testAccLoggingBucketConfigFolder_basic(context, 40),
			},
			{
				ResourceName:            "google_logging_folder_bucket_config.basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"folder"},
			},
		},
	})
}

func TestAccLoggingBucketConfigProject_basic(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"project_name":  "tf-test-" + randString(t, 10),
		"org_id":        getTestOrgFromEnv(t),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccLoggingBucketConfigProject_basic(context, 30),
			},
			{
				ResourceName:            "google_logging_project_bucket_config.basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"project"},
			},
			{
				Config: testAccLoggingBucketConfigProject_basic(context, 40),
			},
			{
				ResourceName:            "google_logging_project_bucket_config.basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"project"},
			},
		},
	})
}

func TestAccLoggingBucketConfigBillingAccount_basic(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":        randString(t, 10),
		"billing_account_name": "billingAccounts/" + getTestBillingAccountFromEnv(t),
		"org_id":               getTestOrgFromEnv(t),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccLoggingBucketConfigBillingAccount_basic(context, 30),
			},
			{
				ResourceName:            "google_logging_billing_account_bucket_config.basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"billing_account"},
			},
			{
				Config: testAccLoggingBucketConfigBillingAccount_basic(context, 40),
			},
			{
				ResourceName:            "google_logging_billing_account_bucket_config.basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"billing_account"},
			},
		},
	})
}

func TestAccLoggingBucketConfigOrganization_basic(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"org_id":        getTestOrgFromEnv(t),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccLoggingBucketConfigOrganization_basic(context, 30),
			},
			{
				ResourceName:            "google_logging_organization_bucket_config.basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"organization"},
			},
			{
				Config: testAccLoggingBucketConfigOrganization_basic(context, 40),
			},
			{
				ResourceName:            "google_logging_organization_bucket_config.basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"organization"},
			},
		},
	})
}

func testAccLoggingBucketConfigFolder_basic(context map[string]interface{}, retention int) string {
	return fmt.Sprintf(Nprintf(`
resource "google_folder" "default" {
	display_name = "%{folder_name}"
	parent       = "organizations/%{org_id}"
}

resource "google_logging_folder_bucket_config" "basic" {
	folder    = google_folder.default.name
	location  = "global"
	retention_days = %d
	description = "retention test %d days"
	bucket_id = "_Default"
}
`, context), retention, retention)
}

func testAccLoggingBucketConfigProject_basic(context map[string]interface{}, retention int) string {
	return fmt.Sprintf(Nprintf(`
 resource "google_project" "default" {
 	project_id = "%{project_name}"
 	name       = "%{project_name}"
 	org_id     = "%{org_id}"
 }

resource "google_logging_project_bucket_config" "basic" {
	project    = google_project.default.name
	location  = "global"
	retention_days = %d
	description = "retention test %d days"
	 bucket_id = "_Default"
}
`, context), retention, retention)
}

func TestAccLoggingBucketConfig_CreateBuckets_withCustomId(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":        randString(t, 10),
		"billing_account_name": "billingAccounts/" + getTestBillingAccountFromEnv(t),
		"org_id":               getTestOrgFromEnv(t),
		"project_name":         "tf-test-" + randString(t, 10),
		"bucket_id":            "tf-test-bucket-" + randString(t, 10),
	}

	configlst := getLoggingBucketConfigs(context)

	for res, config := range configlst {
		vcrTest(t, resource.TestCase{
			PreCheck:  func() { testAccPreCheck(t) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config: config,
				},
				{
					ResourceName:            fmt.Sprintf("google_logging_%s_bucket_config.basic", res),
					ImportState:             true,
					ImportStateVerify:       true,
					ImportStateVerifyIgnore: []string{res},
				},
			},
		})
	}
}

func testAccLoggingBucketConfigBillingAccount_basic(context map[string]interface{}, retention int) string {
	return fmt.Sprintf(Nprintf(`

data "google_billing_account" "default" {
	billing_account = "%{billing_account_name}"
}

resource "google_logging_billing_account_bucket_config" "basic" {
	billing_account    = data.google_billing_account.default.billing_account
	location  = "global"
	retention_days = %d
	description = "retention test %d days"
	bucket_id = "_Default"
}
`, context), retention, retention)
}

func testAccLoggingBucketConfigOrganization_basic(context map[string]interface{}, retention int) string {
	return fmt.Sprintf(Nprintf(`
data "google_organization" "default" {
	organization = "%{org_id}"
}

resource "google_logging_organization_bucket_config" "basic" {
	organization    = data.google_organization.default.organization
	location  = "global"
	retention_days = %d
	description = "retention test %d days"
	bucket_id = "_Default"
}
`, context), retention, retention)
}

func getLoggingBucketConfigs(context map[string]interface{}) map[string]string {

	return map[string]string{
		"organization": Nprintf(`data "google_organization" "default" {
					organization = "%{org_id}"
				}
				
				resource "google_logging_organization_bucket_config" "basic" {
					organization    = data.google_organization.default.organization
					location  = "global"
					retention_days = 10
					description = "retention test 10 days"
					bucket_id = "%{bucket_id}"
				}`, context),

		"billing_account": Nprintf(`data "google_billing_account" "default" {
					billing_account = "%{billing_account_name}"
				}
			
				resource "google_logging_billing_account_bucket_config" "basic" {
					billing_account    = data.google_billing_account.default.billing_account
					location  = "global"
					retention_days = 10
					description = "retention test 10 days"
					bucket_id = "%{bucket_id}"
				}`, context),

		"folder": Nprintf(`resource "google_folder" "default" {
					display_name = "%{folder_name}"
					parent       = "organizations/%{org_id}"
				}
				
				resource "google_logging_folder_bucket_config" "basic" {
					folder    = google_folder.default.name
					location  = "global"
					retention_days = 10
					description = "retention test 10 days"
					bucket_id = "%{bucket_id}"
				}`, context),

		"project": Nprintf(`resource "google_project" "default" {
				project_id = "%{project_name}"
				name       = "%{project_name}"
				org_id     = "%{org_id}"
			}
			
			resource "google_logging_project_bucket_config" "basic" {
				project    = google_project.default.name
				location  = "global"
				retention_days = 10
				description = "retention test 10 days"
				bucket_id = "%{bucket_id}"
			}`, context),
	}

}
