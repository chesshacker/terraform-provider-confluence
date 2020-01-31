---
layout: "confluence"
page_title: "Confluence: confluence_attachment"
sidebar_current: "docs-confluence-resource-attachment"
description: |-
  Provides attachments in Confluence
---

# confluence_attachment

Adds a file attachment to the specified page.

## Example Usage

```hcl
resource confluence_attachment "default" {
  title = "example.txt"
  data  = "This is the contents of the example attachment."
  page  = confluence_content.default.id
}
resource confluence_content "default" {
  title = "Example Page"
  body  = "This page has a <ac:link><ri:attachment ri:filename=\"example.txt\"/><ac:plain-text-link-body><![CDATA[file attachment]]></ac:plain-text-link-body></ac:link>."
  space = "EXAMPLE"
}
```

## Argument Reference

The following arguments are supported:


* `data` - (Required) The contents of the file attached.

* `media_type` - (Optional) The [MIME type] of the attached file. Default is
  text/plain.

* `page` - (Required) The page to which this attachment is added.

* `title` - (Required) The title (or filename) of the attachment.

## Attributes Reference

This resource exports the following attributes:

* `data` - The contents of the file attached.

* `media_type` - The MIME type of the attached file.

* `page` - The page to which this attachment is added.

* `title` - The title (or filename) of the attachment.

* `version` - The version number of the attachment.

## Import

Attachment can be imported using the attachment id.

```
$ terraform import confluence_attachment.default {{id}}
```

[MIME type]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types
