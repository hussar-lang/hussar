# Contributing

This repository is open to contributions from anyone, but please discuss the change(s) you wish to make by opening an issue first.

If you are unsure where to start, you can look at issues tagged as [help wanted](https://github.com/hussar-lang/hussar/labels/help_wanted) or [good first issue](https://github.com/hussar-lang/hussar/labels/good_first_issue).

Please note we have a [code of conduct](.github/CODE_OF_CONDUCT.md), please follow it in all your interactions with the project.

## Workflow

We've started using git flow (with the aid of git-flow-avh) for the development of Hussar. If you are unfamiliar with git flow, read through [this very quick guide](https://danielkummer.github.io/git-flow-cheatsheet/). All contributions should be made as a feature, hotfix or bugfix as appropriate.

We also use [commitizen](https://github.com/commitizen/cz-cli) for our commit messages. If you don't want to install the tool, at least follow the correct style in your commit messages.

## Pull Request Process

1.  Use `gofmt -s`, `golint` and `go vet`. If [Go Report Card](https://goreportcard.com) gives your branch close to 100% (except for gocyclo), you should be good.
2.  Try to use a coding style consistent with the one already used.
3.  Ensure any added code is tested (where applicable).
4.  Ensure any additions are well documented. This means in the documentation as well as in your pull request.
5.  Any pull requests will be tested by Travis CI. All pull requests must build correctly.

If you are unsure about how to do any of the above, feel free to get in touch in the comments of your pull request to ask for help.
