module github.com/galihrivanto/config

go 1.13

replace (
	github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)

require (
	github.com/bitly/go-simplejson v0.5.0
	github.com/coreos/etcd v3.3.18+incompatible
	github.com/fsnotify/fsnotify v1.4.7
	github.com/galihrivanto/runner v0.0.0-20190907143749-e9e818dd7303
	github.com/galihrivanto/x v0.0.0-20200305041419-b83c7ad2dab9
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/protobuf v1.3.4 // indirect
	github.com/imdario/mergo v0.3.8
	github.com/micro/go-micro v1.18.0
	go.uber.org/zap v1.14.0 // indirect
	golang.org/x/sys v0.0.0-20200302150141-5c8b2ff67527 // indirect
	google.golang.org/genproto v0.0.0-20200304201815-d429ff31ee6c // indirect
	google.golang.org/grpc v1.27.1 // indirect
	gopkg.in/yaml.v2 v2.2.2
	gopkg.in/yaml.v3 v3.0.0-20200121175148-a6ecf24a6d71
)
