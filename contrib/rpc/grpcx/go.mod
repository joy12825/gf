module github.com/joy12825/gf/contrib/rpc/grpcx/v2

go 1.18

require (
	github.com/joy12825/gf/v2 v2.7.6
	go.opentelemetry.io/otel v1.14.0
	go.opentelemetry.io/otel/trace v1.14.0
	google.golang.org/grpc v1.57.2
	google.golang.org/protobuf v1.31.0
)

require (
	github.com/BurntSushi/toml v1.3.2 // indirect
	github.com/clbanning/mxj/v2 v2.7.0 // indirect
	github.com/emirpasic/gods v1.18.1 // indirect
	github.com/fatih/color v1.16.0 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/gorilla/websocket v1.5.1 // indirect
	github.com/grokify/html-strip-tags-go v0.1.0 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/olekukonko/tablewriter v0.0.5 // indirect
	github.com/rivo/uniseg v0.4.4 // indirect
	go.opentelemetry.io/otel/sdk v1.14.0 // indirect
	golang.org/x/net v0.24.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230525234030-28d5490b6b19 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	github.com/joy12825/gf/contrib/registry/file/v2 => ../../registry/file/
	github.com/joy12825/gf/v2 => ../../../
)
