- model生成
goctl model mysql ddl -src user.sql -dir . -c
- rpc生成
goctl rpc proto -src user.proto -dir .
- api生成
goctl api go -api user.api -dir .
