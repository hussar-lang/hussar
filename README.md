# kmonkey :monkey:

Master: [![Build Status](https://travis-ci.org/kscarlett/kmonkey.svg?branch=master)](https://travis-ci.org/kscarlett/kmonkey)
[![Coverage Status](https://coveralls.io/repos/github/kscarlett/kmonkey/badge.svg?branch=master)](https://coveralls.io/github/kscarlett/kmonkey?branch=master)

Development: [![Build Status](https://travis-ci.org/kscarlett/kmonkey.svg?branch=develop)](https://travis-ci.org/kscarlett/kmonkey)
[![Coverage Status](https://coveralls.io/repos/github/kscarlett/kmonkey/badge.svg?branch=develop)](https://coveralls.io/github/kscarlett/kmonkey?branch=develop)

[![Documentation Status](https://readthedocs.org/projects/kmonkey/badge/?version=latest)](http://kmonkey.readthedocs.io/en/latest/?badge=latest) [![MIT Licence](https://badges.frapsoft.com/os/mit/mit.svg?v=103)](https://opensource.org/licenses/mit-license.php)

## What's this?

This is my implementation of the Monkey interpreter in Go, based on Thorsten Ball's [Writing an Interpreter in Go](https://interpreterbook.com/) with my own changes, extensions and improvements.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development, as well as normal use of the kmonkey tools.

### Prerequisites

To install kmonkey for development, you need Go on your local machine. For instructions on how to install Go for your OS, follow [the guide on their website](https://golang.org/doc/install).

### Installing

The simplest method to download this project for development is with `go get`. Simply run `go get kscarlett.com/kmonkey` in your terminal of choice, and the project will be downloaded to your Go path.

An alternative method is to clone the project from this repository directly. You can do this by running `git clone git@github.com:kscarlett/kmonkey`. For the Go imports to work correctly without change, you will need to place the downloaded project in `$GOPATH/src/kscarlett.com/kmonkey`. Once you're inside the project directory, you may have to run `go get ./...` to fetch any dependencies.

<!--If you only need the kmonkey tools, without the code, you will soon be able to download and install it simply by executing `curl get.kmonkey.kscarlett.com | sh`, which will ask you for several options and install the right version automatically. -- shhhh, this is coming soon!  -->

### Usage

You can run scripts with the _run_ subcommand, while passing in the script in question.

```
kmonkey run file.km
```

Another option is to simply call the kmonkey command without any subcommands, which will start the interactive mode (or REPL) like so:

```
kmonkey
```

Once in the interactive mode, you can run code and get the result returned. You can exit this mode by calling `exit(0)` or by pressing control-c in your terminal.

## Contributing

Please read [CONTRIBUTING.md](.github/CONTRIBUTING.md) for details on how to contribute to this project and our code of conduct.

## Planned/To Do

- [x] Add basic recursion.
- [ ] Write documentation.
- [ ] Refactor existing code for better organisation.
- [ ] Increase code coverage in critical areas.
- [ ] (LLVM) compilation.
- [ ] More types.
- [ ] Built in libraries e.g. string handling, filesystem I/O, networking.
- [ ] Ability to import user defined libraries.
- [ ] More robust interpreter/compiler/etc -> build system.
- [ ] Automatic updating of tooling.

## Versioning

We are now tagging all stable releases on the master branch. These releases follow [SemVer 2.0](https://semver.org/spec/v2.0.0.html). Development is taking place on the [develop branch](https://github.com/kscarlett/kmonkey/tree/develop).

For the versions available, see the [releases in this repository](https://github.com/kscarlett/kmonkey/releases).

## License

This project is licensed under the MIT license - see the [LICENSE](LICENSE) file for the details.
