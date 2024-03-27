terraform {
  required_providers {
    segment = {
      source  = "scentregroup/segment"
      version = "0.7.0"
    }
  }
}

provider "segment" {
  # Configuration options
}