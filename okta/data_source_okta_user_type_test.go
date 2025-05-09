package okta

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceOktaUserType_read(t *testing.T) {
	resourceName := fmt.Sprintf("data.%s.test", userType)
	mgr := newFixtureManager("data-sources", userType, t.Name())
	createUserType := mgr.GetFixtures("okta_user_type.tf", t)
	readNameConfig := mgr.GetFixtures("read_name.tf", t)
	readIdConfig := mgr.GetFixtures("read_id.tf", t)

	oktaResourceTest(t, resource.TestCase{
		PreCheck:                 testAccPreCheck(t),
		ErrorCheck:               testAccErrorChecks(t),
		ProtoV5ProviderFactories: testAccMergeProvidersFactories,
		Steps: []resource.TestStep{
			{
				Config: createUserType,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("okta_user_type.test", "id"),
				),
			},
			{
				Config: readNameConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			{
				Config: readIdConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(fmt.Sprintf("data.%s.test2", userType), "name"),
				),
			},
		},
	})
}
