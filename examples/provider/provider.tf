terraform {
  required_providers {
    segment = {
      source  = "scentregroup/segment"
      version = "0.6.1"
    }
  }
}

provider "segment" {
  # Configuration options
}