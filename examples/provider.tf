terraform {
  required_providers {
    zeflix = {
      version = "0.1"
      source  = "hashicorp.com/edu/zeflix"
    }
  }
}

provider "zeflix" {
  api_endpoint = "http://localhost:8080"
  api_token = "MySecureToken"
}