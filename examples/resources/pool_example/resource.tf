terraform {
  required_providers {
    ipam = {
      source = "cthiel42/tf-ipam"
      version = "1.0.2"
    }
  }
}

provider "ipam" {}

resource "ipam_pool" "example" {
  name = "pool_example"
  cidrs = [
    "10.0.0.0/24",
    "10.5.0.0/24"
  ]
}
