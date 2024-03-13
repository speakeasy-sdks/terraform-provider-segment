terraform {
  required_providers {
    segment = {
      source  = "scentregroup/segment"
      version = "0.5.1"
    }
  }
}

provider "segment" {
  # Configuration options
}