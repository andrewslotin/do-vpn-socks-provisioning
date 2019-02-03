package vpntests

import (
	"os"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

var APIAccessToken string

func TestMain(m *testing.M) {
	APIAccessToken = os.Getenv("DO_API_TOKEN")
	if APIAccessToken == "" {
		panic("missing DO_API_TOKEN")
	}

	os.Exit(m.Run())
}

func TestVPNProvisioning(t *testing.T) {
	opts := &terraform.Options{
		TerraformDir: "../",
		Vars: map[string]interface{}{
			"do_api_token": APIAccessToken,
		},
	}

	defer terraform.Destroy(t, opts)

	terraform.InitAndApply(t, opts)
}
