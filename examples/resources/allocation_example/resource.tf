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
    "10.0.0.0/16",
    "10.5.0.0/24"
  ]
}

resource "ipam_allocation" "example_0" {
  id            = "allocation_example_0"
  pool_name     = ipam_pool.example.name
  prefix_length = 24
}

resource "ipam_allocation" "example_1" {
  id            = "allocation_example_1"
  pool_name     = ipam_pool.example.name
  prefix_length = 27
}