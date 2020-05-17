module github.com/demonoid81/libnetwork

go 1.14

require github.com/docker/docker v1.13.1

replace go.etcd.io/bbolt => github.com/coreos/bbolt v1.3.4

replace github.com/docker/docker => github.com/docker/engine v17.12.0-ce-rc1.0.20200309214505-aa6a9891b09c+incompatible
