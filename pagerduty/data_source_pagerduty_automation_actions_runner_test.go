package pagerduty

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDataSourcePagerDutyAutomationActionsRunner_Basic(t *testing.T) {
	name := fmt.Sprintf("tf-%s", acctest.RandString(5))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourcePagerDutyAutomationActionsRunnerConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccDataSourcePagerdutyAutomationActionsRunner("pagerduty_automation_actions_runner.test", "data.pagerduty_automation_actions_runner.foo"),
				),
			},
		},
	})
}

func testAccDataSourcePagerdutyAutomationActionsRunner(src, n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		srcR := s.RootModule().Resources[src]
		srcA := srcR.Primary.Attributes

		r, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		a := r.Primary.Attributes

		if a["id"] == "" {
			return fmt.Errorf("No Runner ID is set")
		}

		testAtts := []string{"id", "name", "type", "runner_type", "creation_time", "last_seen", "description", "runbook_base_uri", "teams"}

		for _, att := range testAtts {
			if a[att] != srcA[att] {
				return fmt.Errorf("Expected the runner %s to be: %s, but got: %s", att, srcA[att], a[att])
			}
		}

		return nil
	}
}

func testAccDataSourcePagerDutyAutomationActionsRunnerConfig(name string) string {
	return fmt.Sprintf(`
resource "pagerduty_automation_actions_runner" "test" {
  name = "%s"
  runner_type = "runbook"
  runbook_base_uri = "cat-cat"
  runbook_api_key = "secret"
}

data "pagerduty_automation_actions_runner" "foo" {
  id = pagerduty_automation_actions_runner.test.id
}
`, name)
}
