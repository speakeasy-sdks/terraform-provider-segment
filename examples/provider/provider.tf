terraform {
  required_providers {
    segment = {
      source  = "scentregroup/segment"
      version = "0.3.4"
    }
  }
}

provider "segment" {
  # Configuration options
}