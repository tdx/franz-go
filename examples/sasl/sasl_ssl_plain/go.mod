module sasl_ssl_plain

go 1.18

require github.com/tdx/franz-go v0.8.3

require (
	github.com/klauspost/compress v1.15.6 // indirect
	github.com/pierrec/lz4/v4 v4.1.14 // indirect
	github.com/tdx/franz-go/pkg/kmsg v1.1.0 // indirect
)

replace github.com/tdx/franz-go => ../../..
