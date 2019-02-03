variable "do_api_token" {}

variable "vpn_instance_name" {
  description = "Instance name to use"
  default = ""
}

variable "region" {
  description = "An identifier of DigitalOcean region to put VPN/proxy instance to"
  default = "nyc2"
}

provider "digitalocean" {
  token = "${var.do_api_token}"
}

resource "random_id" "vpn_instance_name" {
  prefix = "vpn-"
  byte_length = 4
}

resource "tls_private_key" "vpn" {
  algorithm = "RSA"
  rsa_bits  = "4096"
}

resource "digitalocean_ssh_key" "vpn" {
  name       = "VPN server"
  public_key = "${tls_private_key.vpn.public_key_openssh}"
}

resource "digitalocean_droplet" "vpn" {
  image  = "ubuntu-18-04-x64"
  name   = "${var.vpn_instance_name != "" ? var.vpn_instance_name : random_id.vpn_instance_name.hex}"
  size   = "s-1vcpu-1gb"
  region = "${var.region}"

  ssh_keys = ["${digitalocean_ssh_key.vpn.fingerprint}"]
}
