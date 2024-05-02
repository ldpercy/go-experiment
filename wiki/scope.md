Scope
=====

What's in/out of scope in various situations.

Affected by:
* package context
* visibility
* mutability
* function context


Packages
--------
As far as I can tell so far there's scope no boundaries between individual package files (as long as the package files are all in the same folder), ie they could all be in one file.

Splitting things up is generally bad, not sure how go decides which is the 'actual' package, probably just lexical/directory order.



### Local and Package variables

Vars declared outside functions are package vars and are accessible across the package.
I dont' think there's a way to set package vars from inside functions.


