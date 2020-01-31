provider "confluence" {
  site  = var.site
  user  = var.user
  token = var.token
}

resource confluence_attachment "example" {
  title = "example.txt"
  data  = "This is the contents of the example attachment."
  page  = confluence_content.example.id
}

resource confluence_content "example" {
  title = "Example Page"
  body  = "This page has a <ac:link><ri:attachment ri:filename=\"example.txt\"/><ac:plain-text-link-body><![CDATA[file attachment]]></ac:plain-text-link-body></ac:link>."
  space = var.space
}

terraform {
  required_version = "~> v0.12.0"
  required_providers {}
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

output "example_content" {
  value = confluence_content.example
}

output "example_attachment" {
  value = confluence_attachment.example
}
