# pdf

如果要增加权限

```
go get github.com/akavel/rsrc
rsrc -manifest pdf.exe.manifest
然后go build -o pdf.exe main.go
生成的就带权限了
```

