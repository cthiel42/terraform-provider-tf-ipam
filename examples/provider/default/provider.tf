// By default, the provider will store data at 
// .terraform/ipam_data.json in the current working
// directory. The file_storage example show how to
// customize this location. 

terraform {
  required_providers {
    ipam = {
      source = "cthiel42/tf-ipam"
      version = "1.0.2"
    }
  }
}

provider "ipam" {}
