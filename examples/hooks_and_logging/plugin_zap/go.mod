module dropin_zap

go 1.18

require (
	github.com/tdx/franz-go v0.8.3
	github.com/tdx/franz-go/plugin/kzap v0.1.0
	go.uber.org/zap v1.21.0
)

require (
	github.com/klauspost/compress v1.15.6 // indirect
	github.com/pierrec/lz4/v4 v4.1.14 // indirect
	github.com/tdx/franz-go/pkg/kmsg v1.1.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	gopkg.in/yaml.v3 v3.0.0 // indirect
)

replace github.com/tdx/franz-go => ../../..
