go build -buildmode=plugin -o en.so ./en/so
go build -buildmode=plugin -o cn.so ./cn/so
go build -o greeter-di ./greeter