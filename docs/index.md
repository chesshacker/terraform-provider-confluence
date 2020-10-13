---
layout: "confluence"
page_title: "Provider: Confluence"
sidebar_current: "docs-confluence-index"
description: |-
  The Confluence provider is used to interact with Confluence.
  It can automate the publishing of content and is often used to publish
  information about other resources created in terraform.
---

# Confluence Provider

The Confluence provider is used to interact with the Confluence. The
provider needs to be configured with the proper credentials before it can be
used.

Use the navigation to the left to read about the available data sources.

## Example Usage

```hcl
provider "confluence" {
  site  = "my-site.atlassian.net"
  user  = "my-user"
  token = "my-token"
}

resource confluence_content "default" {
  space = "MYSPACE"
  title = "Example Page"
  body  = "<p>This page was built with Terraform<p>"
}
```

## Authentication

Static credentials must be passed to the provider block.

## Argument Reference

- `site` - (Required) For Confluence Cloud: The site is the name of the site
  and appears in your wiki URL (https://_my-site.atlassian.net_/wiki/spaces/my-space/).
  For Confluence Server users this should be the hostname of your Confluence
  instance that can receive `/rest/api` requests. This can also be set via the
  `CONFLUENCE_SITE` environment variable.

- `site_schema` - (Optional) Set the schema for connecting to the REST API.
  Defaults to `https`. This can also be set via the `CONFLUENCE_SITE_SCHEMA`
  environment variable.

- `public_site` - (Optional) For Confluence Server instances where your
  Confluence site URL is different than the hostname that serves REST API
  requests. Defaults to `site` if not set. This can also be set via the
  `CONFLUENCE_PUBLIC_SITE` environment variable.

- `public_site_schema` - (Optional) Set the schema for generated public URLs.
  Defaults to `https`. This can also be set via the `CONFLUENCE_PUBLIC_SITE_SCHEMA`
  environment variable.

- `user` - (Required) For Confluence Cloud the user is your user's email
  address. For Confluence Server this is the username of the user to login.
  This can also be set via the `CONFLUENCE_USER` environment variable.

- `token` - (Required) For Confluence Cloud the token is a secret every user
  can generate. It is similar to a password and should be treated as such. For
  Confluence Server, this is the password of the user. This can also be set via
  the `CONFLUENCE_TOKEN` environment variable. For Cloud token details, see
  [Manage your account](https://id.atlassian.com/manage/api-tokens).
