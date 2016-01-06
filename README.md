mkpass
======

[![Build Status](https://travis-ci.org/andytom/mkpass.svg?branch=master)](https://travis-ci.org/andytom/mkpass)

A simple commandline password generator based on the
[xkcd - Password Strength](https://xkcd.com/936/) comic.


Requirements
------------
You need to have go installed to compile the binary.

By default mkpass reads the [Unix words](https://en.wikipedia.org/wiki/Words_%28Unix%29)
file (``/usr/share/dict/words``) as it's input file.


Installation
------------

Assuming that you have [$GOPATH configured](https://golang.org/doc/code.html#GOPATH)
can install using.

```
go install
```

If not you can compile and then complied binary into a directory in your ``$PATH``.

```
go build
cp mkpass ~/bin # For example
```


Usage
-----

You can run it with no arguments to generate a using the defaults.

```
mkpass
```

Full usage can be found using the built in help.

```
mkpass -h
```


Testing
-------

Testing can be run using the built go test framework.

```
go test
```

Automatic testing is also done using [Travis CI](https://travis-ci.org/andytom/mkpass)
against all major version since v1.1.


Licence
-------

See LICENCE.
