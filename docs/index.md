---
layout: "confluence"
page_title: "Provider: Confluence"
sidebar_current: "docs-confluence-index"
description: |-
  The Confluence provider is used to interact with Cloud Confluence.
  It can automate the publishing of content and is often used to publish
  information about other resources created in terraform.
---

# Confluence Provider

The Confluence provider is used to interact with the Cloud Confluence. The
provider needs to be configured with the proper credentials before it can be
used.

Use the navigation to the left to read about the available data sources.

## Example Usage

```hcl
provider "confluence" {
  site  = "my-site"
  user  = "my-user"
  token = "my-token"
}

resource confluence_content "default" {
  space = "MYSPACE"
  title = "Example Page"
  body  = "<p>This page was built with Terraform</p>"
}
```

## Authentication

Static credentials must be passed to the provider block.

## Argument Reference

* `site` - (Required) The site is the name of the site and appears in your wiki
  URL (https://*my-site*.atlassian.net/wiki/spaces/my-space/). This can also be
  set via the `CONFLUENCE_SITE` environment variable.

* `user` - (Required) The user is your user's email address. This can also be
  set via the `CONFLUENCE_USER` environment variable.

* `token` - (Required) The token is a secret every user can generate. It is
  similar to a password and should be treated as such. This can also be set via
  the `CONFLUENCE_TOKEN` environment variable. See [Manage your
  account](https://id.atlassian.com/manage/api-tokens).
