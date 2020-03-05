module github.com/galihrivanto/config

go 1.13

require (
	github.com/bitly/go-simplejson v0.5.0
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869 // indirect
	github.com/coreos/etcd v3.3.18+incompatible
	github.com/coreos/go-systemd v0.0.0-00010101000000-000000000000 // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/fsnotify/fsnotify v1.4.7
	github.com/galihrivanto/x v0.0.0-20200305041419-b83c7ad2dab9
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/protobuf v1.3.4 // indirect
	github.com/google/uuid v1.1.1 // indirect
	github.com/imdario/mergo v0.3.8
	go.uber.org/zap v1.14.0 // indirect
	golang.org/x/sys v0.0.0-20200302150141-5c8b2ff67527 // indirect
	google.golang.org/genproto v0.0.0-20200304201815-d429ff31ee6c // indirect
	google.golang.org/grpc v1.27.1 // indirect
	gopkg.in/yaml.v2 v2.2.2
	gopkg.in/yaml.v3 v3.0.0-20200121175148-a6ecf24a6d71
)

replace github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
