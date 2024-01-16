terraform {
  required_providers {
    segment = {
      source  = "scentregroup/segment"
      version = "0.3.2"
    }
  }
}

provider "segment" {
  # Configuration options
}