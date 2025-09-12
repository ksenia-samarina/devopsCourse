terraform {
  required_providers {
    vkcs = {
      source = "vk-cs/vkcs"
      version = "< 1.0.0"
    }
  }
}

provider "vkcs" {
  # Your user account.
  username = ""

  # The password of the account
  password = ""

  # The tenant token can be taken from the project Settings tab - > API keys.
  # Project ID will be our token.
  project_id = ""

  # Region name
  region = "RegionOne"

  auth_url = "https://infra.mail.ru:35357/v3/"
}
