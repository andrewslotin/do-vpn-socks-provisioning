do-vpn-socks-provisioning
=========================

A Terraform module to create and provision DigtalOcean instance with IPSec VPN and SOCKS5 proxy servers in no time.

Installation
------------

Make sure you have [Terraform](http://terraform.io/) installed and available in your `$PATH`. You will also need an 
account on [DigitalOcean](https://m.do.co/c/50382994d4dd).

To create and provision a new server you need first create a new [personal access token](https://www.digitalocean.com/docs/api/create-personal-access-token/).
This module will use it to spawn a new instance, add a new SSH key to your account and use it to access the server.

```bash
# Clone this repository
git clone git@github.com:andrewslotin/do-vpn-socks-provisioning.git

# Switch to the directory containing Terraform code and apply the provisioning code
cd do-vpn-socks-provisioning
DO_API_TOKEN=<your token> make run
```

Once Terraform is done provisioning server you can see its public IP address as well as VPN and SOCKS5 credentials in its
outputs.

You can always retrieve them from the state file by switching again to the directory where you run `terraform apply` and 
executing `terraform output`.

### !!Warning!!

Your `.tfstate` file generated by Terraform and placed into the current working directory contains sensitive information, such as 
your private access token and SSH key pair. This information can be used by 3rd party to access your DigitalOcean account and
instance. Do not share this file or store it anywhere on Internet. If you absolutely have to do that, consider encrypting it and/or
use secure media.

Clean Up 
--------

Once you are done using VPN/SOCKS5 proxy, you can destroy the instance by switching to the directory where you run `make run`
and execute `DO_API_TOKEN=<your token> make stop`. This will destroy the instance and remove its SSH key from your account.

Credits
-------

This module pulls [`siomiz/softethervpn`](http://hub.docker.com/r/siomiz/softethervpn) and [`schors/tgdante2`](http://hub.docker.com/r/schors/tgdante2) 
Docker images to run VPN and SOCKS5 servers respectively.
