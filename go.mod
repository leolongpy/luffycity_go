module luffycity_go

go 1.16

require (
	github.com/Shopify/sarama v1.19.0
	github.com/asim/go-micro/v3 v3.6.0
	github.com/axgle/mahonia v0.0.0-20180208002826-3358181d7394
	github.com/coreos/etcd v3.3.17+incompatible
	github.com/gin-gonic/gin v1.7.2
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/go-sql-driver/mysql v1.6.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/hpcloud/tail v1.0.0
	github.com/jmoiron/sqlx v1.3.4
	github.com/micro/go-plugins/registry/etcd v0.0.0-20200119172437-4fe21aa238fd
	github.com/orcaman/concurrent-map v0.0.0-20210501183033-44dafcb38ecc
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/prometheus/client_golang v1.11.0 // indirect
	github.com/satori/go.uuid v1.2.0
	github.com/sirupsen/logrus v1.8.1
	github.com/sony/sonyflake v1.0.0
	go.uber.org/zap v1.17.0 // indirect
	golang.org/x/net v0.0.0-20210908191846-a5e095526f91
	golang.org/x/sys v0.0.0-20210909193231-528a39cd75f3 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20210909211513-a8c4777a87af // indirect
	google.golang.org/grpc v1.40.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/yaml.v2 v2.4.0
	sigs.k8s.io/yaml v1.2.0 // indirect
)

replace (
	github.com/coreos/bbolt v1.3.4 => go.etcd.io/bbolt v1.3.4
	github.com/micro/go-micro => github.com/Lofanmi/go-micro v1.16.1-0.20210804063523-68bbf601cfa4
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)
