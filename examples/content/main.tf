# Example: keep track of your pets in Confluence

resource "random_pet" "pets" {
  count     = 4
  separator = " "
  length    = 3
}

provider "confluence" {
  site  = var.site
  user  = var.user
  token = var.token
}

resource confluence_content "example" {
  space = var.space
  title = "My Pets"
  body = templatefile("${path.module}/example.tmpl", {
    pets = [for p in random_pet.pets : title(p.id)]
  })
}

terraform {
  required_version = "~> v0.12.0"
  required_providers {
    random = "~> 2.2"
  }
}

variable "site" {
  type = string
}

variable "user" {
  type = string
}

variable "token" {
  type = string
}

variable "space" {
  type = string
}

output "example" {
  value = confluence_content.example
}
