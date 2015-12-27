mkpass
======


[![Build Status](https://travis-ci.org/andytom/mkpass.svg?branch=master)](https://travis-ci.org/andytom/mkpass)

A simple commandline password generator based on the
[xkcd - Password Strength](https://xkcd.com/936/) comic.

Installation
------------

Compile assuming that you have go installed.

``
go build
``

If you have [$GOPATH configured](https://golang.org/doc/code.html#GOPATH) you
can install using.

``
go install
``

If not you can copy the complied binary into a directory in your ``$PATH``.

By default mkpass reads the [Unix words](https://en.wikipedia.org/wiki/Words_%28Unix%29)
file (``/usr/share/dict/words``) as it's input file.

Usage
-----

You can run it with no arguments to generate a using the defaults.

``
mkpass
``

Full usage can be found using the built in help.

``
mkpass -h
``

Testing
-------

Testing can be run using the built go test framework.

``
go test
``

Licence
-------

See LICENCE.
