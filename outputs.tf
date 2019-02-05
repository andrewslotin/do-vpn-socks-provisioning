output "public_ip" {
  value = "${digitalocean_droplet.vpn.ipv4_address}"
}

output "ssh_key" {
  value = "${tls_private_key.vpn.private_key_pem}"
}

output "vpn_user" {
  value = "${random_string.vpn_user.result}"
}

output "vpn_password" {
  value = "${random_string.vpn_password.result}"
}

output "vpn_psk" {
  value = "${random_string.vpn_psk.result}"
}

output "socks_user" {
  value = "${random_string.socks_user.result}"
}

output "socks_password" {
  value = "${random_string.socks_password.result}"
}
