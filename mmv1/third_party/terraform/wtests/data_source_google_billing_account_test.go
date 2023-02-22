package google_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceGoogleBillingAccount_byFullName(t *testing.T) {
	billingId := acctest.GetTestBillingAccountFromEnv(t)
	name := "billingAccounts/" + billingId

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:  func() { acctest.TestAccPreCheck(t) },
		Providers: acctest.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckGoogleBillingAccount_byName(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.google_billing_account.acct", "id", billingId),
					resource.TestCheckResourceAttr("data.google_billing_account.acct", "name", name),
					resource.TestCheckResourceAttr("data.google_billing_account.acct", "open", "true"),
				),
			},
		},
	})
}

func TestAccDataSourceGoogleBillingAccount_byShortName(t *testing.T) {
	billingId := acctest.GetTestBillingAccountFromEnv(t)
	name := "billingAccounts/" + billingId

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:  func() { acctest.TestAccPreCheck(t) },
		Providers: acctest.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckGoogleBillingAccount_byName(billingId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.google_billing_account.acct", "id", billingId),
					resource.TestCheckResourceAttr("data.google_billing_account.acct", "name", name),
					resource.TestCheckResourceAttr("data.google_billing_account.acct", "open", "true"),
				),
			},
		},
	})
}

func TestAccDataSourceGoogleBillingAccount_byFullNameClosed(t *testing.T) {
	billingId := acctest.GetTestBillingAccountFromEnv(t)
	name := "billingAccounts/" + billingId

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:  func() { acctest.TestAccPreCheck(t) },
		Providers: acctest.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckGoogleBillingAccount_byNameClosed(name),
				ExpectError: regexp.MustCompile("Billing account not found: " + name),
			},
		},
	})
}

func TestAccDataSourceGoogleBillingAccount_byDisplayName(t *testing.T) {
	name := acctest.RandString(t, 16)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:  func() { acctest.TestAccPreCheck(t) },
		Providers: acctest.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckGoogleBillingAccount_byDisplayName(name),
				ExpectError: regexp.MustCompile("Billing account not found: " + name),
			},
		},
	})
}

func testAccCheckGoogleBillingAccount_byName(name string) string {
	return fmt.Sprintf(`
data "google_billing_account" "acct" {
  billing_account = "%s"
}
`, name)
}

func testAccCheckGoogleBillingAccount_byNameClosed(name string) string {
	return fmt.Sprintf(`
data "google_billing_account" "acct" {
  billing_account = "%s"
  open            = false
}
`, name)
}

func testAccCheckGoogleBillingAccount_byDisplayName(name string) string {
	return fmt.Sprintf(`
data "google_billing_account" "acct" {
  display_name = "%s"
}
`, name)
}