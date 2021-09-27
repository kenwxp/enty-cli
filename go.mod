module entysquare/enty-cli

go 1.15

require (
	github.com/lib/pq v1.9.0
	github.com/satori/go.uuid v1.2.0
	github.com/sirupsen/logrus v1.7.0
	github.com/tealeg/xlsx v1.0.5
	github.com/tidwall/gjson v1.9.0
	github.com/urfave/cli/v2 v2.3.0 // indirect
	golang.org/x/sys v0.0.0-20210330210617-4fbd30eecc44 // indirect
)

//replace github.com/entysquare/payment-sdk-go => ../payment-sdk-go
