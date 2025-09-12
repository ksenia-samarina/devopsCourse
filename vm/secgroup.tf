# resource "vkcs_networking_secgroup" "ssh" {
#   name        = "ssh"
#   description = "Allow SSH access"
# }
#
# resource "vkcs_networking_secgroup_rule" "ssh_rule" {
#   security_group_id = vkcs_networking_secgroup.ssh.id
#   direction         = "ingress"
#   ethertype         = "IPv4"
#   protocol          = "tcp"
#   port_range_min    = 22
#   port_range_max    = 22
#   remote_ip_prefix  = "0.0.0.0/0"
# }

resource "vkcs_networking_secgroup" "registry_sg" {
  name        = "registry_sg"
  description = "Security group for Docker Registry"
}

resource "vkcs_networking_secgroup_rule" "registry_tcp" {
  security_group_id = vkcs_networking_secgroup.registry_sg.id
  direction         = "ingress"
  protocol          = "tcp"
  port_range_min    = 5000
  port_range_max    = 5000
  remote_ip_prefix  = "0.0.0.0/0"
  description       = "Allow Docker Registry"
}