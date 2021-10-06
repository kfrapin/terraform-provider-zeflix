terraform {
  required_providers {
    zeflix = {
      version = "0.1"
      source  = "hashicorp.com/edu/zeflix"
    }
  }
}

provider "zeflix" {
  // no configuration required
}