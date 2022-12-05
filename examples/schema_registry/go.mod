module schema_registry

go 1.18

replace (
	github.com/tdx/franz-go => ../..
	github.com/tdx/franz-go/pkg/sr => ../../pkg/sr
)

require (
	github.com/hamba/avro v1.8.0
	github.com/tdx/franz-go v1.5.3
	github.com/tdx/franz-go/pkg/sr v0.0.0-20220525220550-fe9d5da387a2
)

require (
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.15.6 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pierrec/lz4/v4 v4.1.14 // indirect
	github.com/tdx/franz-go/pkg/kmsg v1.1.0 // indirect
)
