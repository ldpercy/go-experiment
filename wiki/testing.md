Testing
=======

https://pkg.go.dev/testing


Test function signature:

```go
func TestXxx(*testing.T)
```

Private testing in adjacent files:

```
xxx_test.go
```

Public-only testing in adjacent packages:

```
packagename_test
```