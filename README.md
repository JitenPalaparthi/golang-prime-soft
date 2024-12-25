## Golang

- TO run go application

```
go run main.go
```

- To build go application

```
go build main.go
```
- To stripe down the binary size

```
go build -ldflags="-w -s" -o hello1 main.go
```

- To get the OS/ARCH information

```
 go tool dist list
 ```
- To create cross compilation and build

```
GOOS=android GOARCH=arm64 go build -o build/demo-android-arm64 -ldflags="-w -s"  main.go
```

## Keywords

break,case, const,continue,default, if, else, fallthrough, for,func,goto, import, map,package,range,return,switch,var (18 out of 25)

## builtin

append,cap,clear,complex,copy, delete,len,imag, make, max,min,new,print,println,real (15 out of 18)

## packages

1. standard packages  : Packages those ship with Golang 
2. User defined packages : Packages those are created by the developer in same or for the same project
3. Third Party packages : Packages are created by third parties and developers uses them.

- GOROOT
- GOPATH
- GOBIN

