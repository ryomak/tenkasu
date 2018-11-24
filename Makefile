run:
	go run src/cmd/main.go

echo-coin:
	go run src/broker_control/binance/tool/echoPair.go src/broker_control/binance/tool/tool.go
	go run src/broker_control/binance/tool/writeRoute.go src/broker_control/binance/tool/tool.go

route:
	go run src/broker_control/binance/tool/writeRoute.go src/broker_control/binance/tool/tool.go
