data "vkcs_compute_flavor" "compute" {
  name = var.compute_flavor
}

data "vkcs_images_image" "compute" {
  visibility = "public"
  default    = true
  properties = {
    mcs_os_distro  = "ubuntu"
    mcs_os_version = "24.04"
  }
}

resource "vkcs_compute_instance" "compute" {
  name              = "compute-instance"
  flavor_id         = data.vkcs_compute_flavor.compute.id
  key_pair          = var.key_pair_name
  security_groups   = ["default", "ssh", "registry_sg"]
  availability_zone = var.availability_zone_name

  block_device {
    uuid                  = data.vkcs_images_image.compute.id
    source_type           = "image"
    destination_type      = "volume"
    volume_type           = "ceph-ssd"
    volume_size           = 10
    boot_index            = 0
    delete_on_termination = true
  }

  network {
    uuid = vkcs_networking_network.example.id
  }

  depends_on = [
    vkcs_networking_network.example,
    vkcs_networking_subnet.example
  ]

  user_data = <<-EOF
        #cloud-config
        package_update: true
        package_upgrade: true
        packages:
          - docker.io

        runcmd:
          # включаем Docker
          - systemctl enable docker
          - systemctl start docker
          - sleep 15

          - mkdir -p /opt/registry/data /opt/registry/auth
          - docker run --rm --entrypoint htpasswd registry:2.7.0 -Bbn myuser mypassword > /opt/registry/auth/htpasswd

          - |
            cat <<EOL > /etc/systemd/system/registry.service
            [Unit]
            Description=Docker Registry
            After=docker.service
            Requires=docker.service

            [Service]
            Restart=always
            ExecStart=/usr/bin/docker run --rm --name registry -p 5000:5000 \
              -v /opt/registry/data:/var/lib/registry \
              -v /opt/registry/auth:/auth \
              -e REGISTRY_AUTH=htpasswd \
              -e REGISTRY_AUTH_HTPASSWD_REALM="basic-realm" \
              -e REGISTRY_AUTH_HTPASSWD_PATH=/auth/htpasswd \
              registry:2.7.0

            [Install]
            WantedBy=multi-user.target
            EOL
          - systemctl daemon-reload
          - systemctl enable registry
          - systemctl start registry
              EOF
}

resource "vkcs_networking_floatingip" "fip" {
  pool = data.vkcs_networking_network.extnet.name
}

resource "vkcs_compute_floatingip_associate" "fip" {
  floating_ip = vkcs_networking_floatingip.fip.address
  instance_id = vkcs_compute_instance.compute.id
}

output "instance_fip" {
  value = vkcs_networking_floatingip.fip.address
}