<h1 align="center">
  <br>
  <a href="https://www.hussar.io"><img src="assets/hussar.png" alt="Hussar Logo" width="200"></a>
  <br>
  Hussar
  <br>
</h1>

<h4 align="center">A small programming language (very much under development).</h4>

<p align="center">
  <a href="https://travis-ci.org/hussar-lang/hussar">
    <img src="https://travis-ci.org/hussar-lang/hussar.svg?branch=develop"
         alt="Build Status">
  </a>
  <a href="https://codeclimate.com/github/hussar-lang/hussar/maintainability"><img src="https://api.codeclimate.com/v1/badges/7f5869af27d7d9ecf476/maintainability" /></a>
  <a href="https://opensource.org/licenses/mit-license.php">
    <img src="https://badges.frapsoft.com/os/mit/mit.svg?v=103"
         alt="License">
  </a>
</p>

<p align="center">
  <!--<a href="#key-features">Key Features</a> •-->
  <a href="#getting-started">Getting Started</a> •
  <a href="#usage">Usage</a> •
  <a href="#contributing">Contributing</a> •
  <a href="#planned">Planned Features</a> •
  <a href="#versioning">Versioning</a> •
  <a href="#license">License</a>
</p>

## What's this?

This is my implementation of the Monkey interpreter in Go, based on Thorsten Ball's [Writing an Interpreter in Go](https://interpreterbook.com/) with my own changes, extensions and improvements. There are a lot of big changes planned, so this will probably divert a lot from the original project.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development, as well as normal use of the Hussar tools.

### Prerequisites

To install Hussar for development, you need Go on your local machine. For instructions on how to install Go for your OS, follow [the guide on their website](https://golang.org/doc/install).

### Installing

For normal usage, you can install Hussar using Homebrew like this:

```console
$ brew install hussar-lang/tap/hussar
```

Alternatively, you can download a copy from the [releases](https://github.com/hussar-lang/hussar/releases).

#### For Development

The simplest method to download this project for development is with `go get`. Simply run `go get github.com/hussar-lang/hussar` in your terminal of choice, and the project will be downloaded to your Go path.

An alternative method is to clone the project from this repository directly. You can do this by running `git clone git@github.com:hussar-lang/hussar`. For the Go imports to work correctly without change, you will need to place the downloaded project in `$GOPATH/src/github.com/hussar-lang/hussar`. Once you're inside the project directory, you may have to run `go get ./...` to fetch any dependencies.

<!--If you only need the Hussar tools, without the code, you will soon be able to download and install it simply by executing `curl get.hussar.io | sh`, which will ask you for several options and install the right version automatically. -- shhhh, this is coming soon!  -->

## Usage

You can run scripts with the _run_ subcommand, while passing in the script in question.

```console
$ hussar run file.hss
```

Another option is to simply call the Hussar command without any subcommands, which will start the interactive mode (or REPL) like so:

```console
$ hussar
```

Once in the interactive mode, you can run code and get the result returned. You can exit this mode by calling `exit(0)` or by pressing control-c in your terminal.

For more information about the tooling and which commands are available to you, you can run the `env` subcommand:

```console
$ hussar env
arch:  amd64
os:    darwin
bin:   hussar
gc:    go1.12
vers:  0.4.0
```

## Contributing

Please read [CONTRIBUTING.md](.github/CONTRIBUTING.md) for details on how to contribute to this project and our code of conduct.

## Planned

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

We are now tagging all stable releases on the master branch. These releases follow [SemVer 2.0](https://semver.org/spec/v2.0.0.html). Development is taking place on the [develop branch](https://github.com/hussar-lang/hussar/tree/develop).

For the versions available, see the [releases in this repository](https://github.com/hussar-lang/hussar/releases).

## License

This project is licensed under the MIT license - see the [LICENSE](LICENSE) file for the details.
