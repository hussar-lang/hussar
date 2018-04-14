# kmonkey :monkey:

Master: [![Build Status](https://travis-ci.org/kscarlett/kmonkey.svg?branch=master)](https://travis-ci.org/kscarlett/kmonkey)
[![Coverage Status](https://coveralls.io/repos/github/kscarlett/kmonkey/badge.svg?branch=master)](https://coveralls.io/github/kscarlett/kmonkey?branch=master)

Experimental: [![Build Status](https://travis-ci.org/kscarlett/kmonkey.svg?branch=experimental)](https://travis-ci.org/kscarlett/kmonkey)
[![Coverage Status](https://coveralls.io/repos/github/kscarlett/kmonkey/badge.svg?branch=experimental)](https://coveralls.io/github/kscarlett/kmonkey?branch=experimental)

 [![Documentation Status](https://readthedocs.org/projects/kmonkey/badge/?version=latest)](http://kmonkey.readthedocs.io/en/latest/?badge=latest) [![MIT Licence](https://badges.frapsoft.com/os/mit/mit.svg?v=103)](https://opensource.org/licenses/mit-license.php)

## What's this?
This is my implementation of the Monkey interpreter in Go, based on Thorsten Ball's [Writing an Interpreter in Go](https://interpreterbook.com/) with my own changes, extensions and improvements (coming soon).

Once the project is more feature complete, I start tagging releases, and the master branch will only contain stable releases. These releases will follow [SemVer 2.0](https://semver.org/spec/v2.0.0.html) and each major and minor update will be named as well. Development will then take place on what is currently the *experimental* branch.

## Planned/To Do
- [x] Add basic recursion.
- [ ] Write documentation.
- [ ] Refactor existing code for better organisation.
- [ ] Increase code coverage in critical areas.
- [ ] LLVM compilation.
- [ ] More types.
- [ ] Built in libraries e.g. string handling, filesystem I/O, networking.
- [ ] Ability to import user defined libraries.
- [ ] More robust interpreter/compiler/etc -> build system.
