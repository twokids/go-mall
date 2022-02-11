# go-mall
a exercise for go zero mall


## model生成

cd service/user && goctl model mysql ddl -src ./model/user.sql -dir ./model -c
cd service/product && goctl model mysql ddl -src ./model/product.sql -dir ./model -c
cd service/order && goctl model mysql ddl -src ./model/order.sql -dir ./model -c
cd service/pay && goctl model mysql ddl -src ./model/pay.sql -dir ./model -c

cd  service/user
goctl api go -api ./api/user.api -dir ./api
goctl rpc proto -src ./rpc/user.proto -dir ./rpc

cd  service/product
goctl api go -api ./api/product.api -dir ./api
goctl rpc proto -src ./rpc/product.proto -dir ./rpc

cd  service/order
goctl api go -api ./api/order.api -dir ./api
goctl rpc proto -src ./rpc/order.proto -dir ./rpc

cd  service/pay
goctl api go -api ./api/pay.api -dir ./api
goctl rpc proto -src ./rpc/pay.proto -dir ./rpc