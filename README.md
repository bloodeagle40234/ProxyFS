# ProxyFS
File Server for Swift Clusters

## Synopsis

ProxyFS provides a POSIX-ish hierarchical file system backed by a Swift cluster, making it available via SMB, NFS, and the Swift and S3 APIs.

## Installation

* Define your GOPATH as desired (where your bin/, /pkg, and /src/ directory trees will appear)
* cd $GOPATH
* mkdir -p src/github.com/swiftstack
* cd src/github.com/swiftstack
* git clone git@github.com:swiftstack/ProxyFS.git
* cd ProxyFS
* git submodule update --init --recursive

## Build

* Install/update to at least Go 1.8.3
* cd $GOPATH/src/github.com/swiftstack/ProxyFS
* ./regression_test.py

## Development

The officially supported development environment for ProxyFS is
[Runway](https://github.com/swiftstack/runway). Please read Runway's
documentation and make sure to use the ProxyFS manifest instead of the default
one.

## Contributors

 * balaji@swiftstack.com
 * bschatz@swiftstack.com
 * charmer@swiftstack.com
 * ed@swiftstack.com
 * gerard@swiftstack.com
 * kmalone@swiftstack.com
 * orion@swiftstack.com
 * zdavis@swiftstack.com

## License

See LICENSE file
