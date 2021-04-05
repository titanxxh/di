#   di(dynamic injection)
this lib can be used to dynamically inject a dependency from a dynamic lib.

for detail, check out [this blog](https://titanxxh.com/2021/03/31/golang-dynamic-injection/) (Chinese)

this lib combines 
-   golang [plugin](https://golang.org/pkg/plugin/) mechanism 
-   and a tricky [reflection](api.go) mechanism 
to implement dynamic dependency injection in golang.

check [example](example/greeter/main.go) code. 

we use [example/build.sh](example/build.sh) to build 2 plugins, 
one is Chinese Greeter and the other is English Greeter.

the both plugins implements [Greeter](example/api.go) interface.

we can run the [main](example/greeter/main.go) with 2 args.
-   first is the `so` file path, `./en/so`
-   second is the package path, 
    -   `github.com/titanxxh/di/example/en.english`, if your package name is same as your folder name.
    -   `*github.com/titanxxh/di/example/cn/v1.(cn).Chinese`, if your package name is different from your folder name.