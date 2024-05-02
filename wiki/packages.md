
Packages
========

Packages should/must be in a single folder:

https://go.dev/ref/spec#Package_clause

> An implementation may require that all source files for a package inhabit the same directory.

Package files that end up elsewhere become more or less  unreachable.


Internal packages
-----------------

https://go.dev/doc/go1.4#internalpackages

https://docs.google.com/document/d/1e8kOo3r51b2BWtTs_1uADIA5djfXhPT36s6eHVRIvaU/edit

> as of Go 1.4 the go command introduces a mechanism to define “internal” packages that may not be imported by packages outside the source subtree in which they reside.
> To create such a package, place it in a directory named internal or in a subdirectory of a directory named internal.