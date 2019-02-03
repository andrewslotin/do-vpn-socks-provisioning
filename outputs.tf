output "public_ip" {
  value = "${digitalocean_droplet.vpn.ipv4_address}"
}
