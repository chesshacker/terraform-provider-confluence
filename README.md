# Terraform Provider for Confluence

On more than one occassion, I have wished for an easy way to update Confluence
Cloud pages from code. Terraform seems like a helpful tool to accomplish this
task, but I didn't see any providers for Confluence... so, I am going to give it
a try.

From the Confluence website, I found I could [Manage API Tokens], creating a
token for use with the [Confluence REST API Reference]. To get started, I saved
my username, API token, and instance name in `secrets.env`. If you'd like to
follow along, you can copy secrets.template.env to secrets.env and replace the
placeholders with values for your environment. Your instance name is whatever
comes before atlassian.net in the URL of your Confluence site.

A quick test with curl shows it works!

```
source secrets.env
curl -s -u "${EMAIL}:${APIKEY}" \
  "https://${INSTANCE}.atlassian.net/wiki/rest/api/content/" | jq
```

Note, I am using [jq] above to make the output more readable. You can leave
off `| jq` if you prefer.

Hopefully that works for you!

## Terraform Provider Basics

Terraform's guide on [Writing Custom Providers] was very helpful. I was able to
throw together a skeleton of a provider. Next, I need to fill it with actual
Confluence calls.

To see it in action, you can run:

```
make
terraform init
terraform apply
```

[Confluence REST API Reference]: https://docs.atlassian.com/atlassian-confluence/REST/1000.0.0-SNAPSHOT/
[Manage API Tokens]: https://id.atlassian.com/manage/api-tokens
[jq]: https://stedolan.github.io/jq/
[Writing Custom Providers]: https://www.terraform.io/docs/extend/writing-custom-providers.html
