package serviceendpoint

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataServiceEndpointGithub() *schema.Resource {
	r := dataSourceGenBaseServiceEndpointResource(dataSourceServiceEndpointGithubRead)
	return r
}

func dataSourceServiceEndpointGithubRead(d *schema.ResourceData, m interface{}) error {
	serviceEndpoint, projectID, err := dataSourceGetBaseServiceEndpoint(d, m)
	if err != nil {
		return err
	}
	if serviceEndpoint != nil {
		doBaseFlattening(d, serviceEndpoint, projectID)
		return nil
	}
	return fmt.Errorf("Error looking up service endpoint!")
}
