# Terraform Provider for Confluence

On more than one occassion, I have wished for an easy way to update Confluence
Cloud pages from code. Terraform seems like a helpful tool to accomplish this
task, but I didn't see any providers for Confluence, so I built one. This works
for me, but it is pretty raw. I plan to overhall the file structure to make it
match most of the conventions used by the official terraform providers.

## How to use it?

To build the terraform provider, you should have a recent version of Go
installed. Building with `make` will create the terraform-provider-confluence
binary in this directory. Terraform looks in the current directory for custom
providers. If you wish to use it in other terraform projects, `make install`
will copy it to your `~/terraform.d/plugins` directory, making it generally
available from any directory on your computer.

The provider requires a few pieces of information to authenticate with
Confluence. To try out the `example.tf` file, you need to copy the
`terraform.template.tfvars` file to `terraform.tfvars` and replace the
placeholders with your settings. The `instance` and `space` can be found in the
URL on confluence.

```
https://*your_instance*.atlassian.net/wiki/spaces/*your_space*/...
```

`user` is your email, and `token` can be generated in [Your Account
Management](https://id.atlassian.com/manage/api-tokens).


After building the provider and creating `terraform.tfvars` you are ready to run
the example, which will build a Confluence page containing a list of your pets.
See `example.tf` to understand how the provider and page are constructed. Then
run the following commands:

```
terraform init
terraform apply
```

Note that the outputs include a URL you can use to visit your page directly.
Feel free to update the `example.tf` or `example.tmpl` files and re-apply to see
your changes applied to the Confluence page. When you are done, you can run
`terraform destroy` to delete the Confluence page.

## Helpful Resources

* [Confluence REST API Reference](https://docs.atlassian.com/atlassian-confluence/REST/1000.0.0-SNAPSHOT/)
* [Confluence Basic Auth for REST APIs](https://developer.atlassian.com/cloud/confluence/basic-auth-for-rest-apis/)
* [Confluence Storage Format](https://confluence.atlassian.com/doc/confluence-storage-format-790796544.html)
* [Terraform Provider Plugins](https://www.terraform.io/docs/plugins/provider.html)
* [Extending Terraform Docs](https://www.terraform.io/docs/extend/index.html)
