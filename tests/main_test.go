package vpntests

import (
	"log"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/retry"
	"github.com/gruntwork-io/terratest/modules/ssh"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

	dropletIP := terraform.Output(t, opts, "public_ip")
	require.NotEmpty(t, dropletIP)

	output := retry.DoWithRetry(t, "ping created droplet", 3, 10*time.Second, func() (string, error) {
		cmd := exec.Command("ping", "-c", "1", "-t", "1", dropletIP)
		output, err := cmd.CombinedOutput()
		log.Println(string(output))
		return string(output), err
	})
	require.Contains(t, output, "0.0% packet loss")

	sshKey := terraform.Output(t, opts, "ssh_key")
	require.NotEmpty(t, sshKey)

	output = ssh.CheckSshCommand(t, ssh.Host{
		Hostname:    dropletIP,
		SshUserName: "root",
		SshKeyPair: &ssh.KeyPair{
			PrivateKey: sshKey,
		},
	}, "docker ps")
	assert.Contains(t, output, "siomiz/softethervpn:alpine")
	assert.Contains(t, output, "schors/tgdante2:latest")

	vpnUser := terraform.Output(t, opts, "vpn_user")
	assert.NotEmpty(t, vpnUser)

	vpnPassword := terraform.Output(t, opts, "vpn_password")
	assert.NotEmpty(t, vpnPassword)

	vpnPSK := terraform.Output(t, opts, "vpn_psk")
	assert.NotEmpty(t, vpnPSK)

	socksUser := terraform.Output(t, opts, "socks_user")
	assert.NotEmpty(t, socksUser)

	socksPassword := terraform.Output(t, opts, "socks_password")
	assert.NotEmpty(t, socksPassword)
}
