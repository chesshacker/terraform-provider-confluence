# Contributing to this Project

Contributions are welcome! The backlog is kept as GitHub issues.

## Checking for problems

Several checks and tests are run as part of the CI/CD. Detected problems will
prevent the merge of the PR. Contributors can run the same checks locally:

  ```bash
  make check
  make test
  make testacc
  ```

Note: You will need to configure several environment variables specific to your
confluence setup before you can run the acceptance tests. To make that part
easier, I have included a template. Copy `secrets.template.env` to `secrets.env`
and edit its values. Then run `source secrets.env` before running `make testacc`
command.

The linter used by this project is [revive](https://revive.run/docs). The
[revive.toml](revive.toml) file can be adjusted if the default rules seem to be
problematic.

## Contributing new features

* Please search for an existing issue before creating a new one

* Simple code changes can be submitted directly as PRs. If your contribution is
  more involved, please open an issue first for discussion.

## Testing the docs

This project uses MkDocs to build and deploy documentation to GitHub Pages. This
process is automated by GitHub Actions.

To serve documentation locally requires Python 3.x. To install the dependencies
and serve the docs:

```
pip install -r requirements.txt
mkdocs serve
```

## Releases

This project hasn't committed to a versioned release yet, but plans to do
releases in the future.
