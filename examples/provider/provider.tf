terraform {
  required_providers {
    segment = {
      source  = "scentregroup/segment"
      version = "0.3.3"
    }
  }
}

provider "segment" {
  # Configuration options
}