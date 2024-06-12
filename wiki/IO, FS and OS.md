IO, FS and OS
=============


I'm still not sure which of these to use in what circumstances....
I think for things I'm interested in I should probably start with OS.
FS seems to more about filesystem virtualisation for things like testing/embedding, so not of immediate interest to me right now.



OS - operating system
---------------------


https://pkg.go.dev/os




IO
---

https://pkg.go.dev/io

>	Package io provides basic interfaces to I/O primitives. Its primary job is to wrap existing implementations of such primitives, such as those in package os, into shared public interfaces that abstract the functionality, plus some other related primitives.
>
>	Because these interfaces and primitives wrap lower-level operations with various implementations, unless otherwise informed clients should not assume they are safe for parallel execution.


FS - file system
----------------

https://pkg.go.dev/io/fs

> basic interfaces to a file system

https://www.gopherguides.com/articles/golang-1.16-io-fs-improve-test-performance

> the io/fs package allows for the implementation of a virtual file system


Other
-----

### Filepath
https://pkg.go.dev/path/filepath



### ioutil - deprecated
https://pkg.go.dev/io/ioutil










