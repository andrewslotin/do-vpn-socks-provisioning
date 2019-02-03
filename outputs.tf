output "public_ip" {
  value = "${digitalocean_droplet.vpn.ipv4_address}"
}

output "ssh_key" {
  value = "${tls_private_key.vpn.private_key_pem}"
}
