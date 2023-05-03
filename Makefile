#集成ci检查
# make verifiers

getdeps:
   @mkdir -p ${GOPATH}/bin
   @which golangci-lint 1>/dev/null || (echo "Installing golangci-lint" && go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.52.2)
lint:

   @echo "Running $@ check"
   @GO111MODULE=on ${GOPATH}/bin/golangci-lint cache clean
   @GO111MODULE=on ${GOPATH}/bin/golangci-lint run --timeout=5m --config ./.golangci.yml
verifiers: getdeps lint