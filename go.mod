module entysquare/filer-backend

go 1.15

require (
	github.com/alibabacloud-go/darabonba-openapi v0.1.5
	github.com/alibabacloud-go/dm-20151123 v1.0.0
	github.com/alibabacloud-go/tea v1.1.15
	github.com/aliyun/aliyun-oss-go-sdk v2.1.8+incompatible
	github.com/baiyubin/aliyun-sts-go-sdk v0.0.0-20180326062324-cfa1a18b161f // indirect
	github.com/gorilla/mux v1.8.0
	github.com/lib/pq v1.9.0
	github.com/satori/go.uuid v1.2.0
	github.com/sirupsen/logrus v1.7.0
	github.com/tealeg/xlsx v1.0.5
	github.com/tidwall/gjson v1.9.0
	github.com/yunpian/yunpian-go-sdk v0.0.0-20171206021512-2193bf8a7459
	golang.org/x/net v0.0.0-20210415231046-e915ea6b2b7d // indirect
	golang.org/x/time v0.0.0-20210220033141-f8bda1e9f3ba // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

//replace github.com/entysquare/payment-sdk-go => ../payment-sdk-go
