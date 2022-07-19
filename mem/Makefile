
coverage: test
	@echo "===> Creating coverage report"
	go tool cover -html=coverage.txt

# build:
# 	@echo "===> Building"
# 	CGO_ENABLED=0 GOOS=linux GARCH=amd64 go build -o app $(MAINFILES)


# run-local:
# 	@echo "===>  Running $(APP_NAME) locally"
# 	@GO111MODULE=on go run -race $(MAINFILES)

bench-mem:
	go test -bench=. -benchmem -run=^$ -benchmem -memprofile=mem.out

bench-cpu:
	go test -bench=. -run=^$ -benchmem -cpuprofile=cpu.out

pprof-cpu: bench-cpu
	go tool pprof -http=:8080 cpu.out

pprof-mem: bench-mem
	go tool pprof -http=:8080 mem.out
