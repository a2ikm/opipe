opipe
==========

Pass STDIN to [OS X's open command](https://developer.apple.com/library/mac/documentation/Darwin/Reference/ManPages/man1/open.1.html) as arguments.

**EDIT: `xargs` just do this like `find . | xargs open -a Atom`.**

Installation
------------

    $ go get github.com/a2ikm/opipe

Usage
------------

From pipe:

    $ find config/environments -name "*.rb" | opipe -a Atom

From redirection:

    $ find config/environments -name "*.rb" > files.txt
    $ opipe -a Atom < files.txt

Conflicts
------------

- opipe doesn't work with open's `-f` option.
