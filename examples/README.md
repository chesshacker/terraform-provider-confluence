# Examples for the Confluence Terraform Provider

To successfully run any of these examples, you must provide information to
access the confluence API. To make that part easier, a template is provided.
Copy `secrets.template.env` in the root directory of this repository to
`secrets.env` and edit the values. Before running any of the examples, source
this file `source secrets.env`.

Sourcing the file sets environment variables for CONFLUENCE_SITE,
CONFLUENCE_USER, CONFLUENCE_TOKEN and CONFLUENCE_SPACE. Environment variables
are one way to configure provider and resource values. If no other value is
specified, these environment variables are used as default values. The first
three are used to configure the confluence provider, and the CONFLUENCE_SPACE is
used any time a resource needs to specify a space.

Configuring this will also provide the values you need to run acceptance tests
with `make testacc`.
