terraform {
  required_providers {
    segment = {
      source  = "scentregroup/segment"
      version = "0.4.0"
    }
  }
}

provider "segment" {
  # Configuration options
}