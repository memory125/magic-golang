## etcd注意事项
### go.mod
- 第一步：
`go mod init`
- 第二步
`go mod tidy`时，出现下述错误：
  `/etcd imports
  github.com/coreos/etcd/clientv3 tested by
  github.com/coreos/etcd/clientv3.test imports
  github.com/coreos/etcd/auth imports
  github.com/coreos/etcd/mvcc/backend imports
  github.com/coreos/bbolt: github.com/coreos/bbolt@v1.3.5: parsing go.mod:
  module declares its path as: go.etcd.io/bbolt
  but was required as: github.com/coreos/bbolt`
  
### 解决方案
- 修改go.mod
`go mod edit -replace github.com/coreos/bbolt@v1.3.4=go.etcd.io/bbolt@v1.3.4`
- 运行
`go mod tidy`
- 如果出现下述错误
`imports
  google.golang.org/grpc/naming: module google.golang.org/grpc@latest found (v1.32.0), but does not contain package google.golang.org/grpc/naming
  `
- 解决方案还是要修改go.mod
`go mod edit -replace google.golang.org/grpc@v1.32.0=google.golang.org/grpc@v1.26.0`
- 重新下载
`go mod tidy`