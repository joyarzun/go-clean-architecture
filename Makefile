.PHONY: test testcov opencov

test:
	ginkgo -v ./...

testcov:
	ginkgo -v -cover -covermode=count -coverpkg=./... ./...

showcov: testcov
	go tool cover -func coverprofile.out

opencov:
	go tool cover -html=coverprofile.out

govet:
	go vet ./...

staticcheck:
	staticcheck ./...

COV_THRESHOLD = 80
testcovci: testcov
	go tool cover -func coverprofile.out | grep total | awk -v COV_THRESHOLD=$(COV_THRESHOLD) '{print "scale=0;("substr($$3, 1, length($$3)-1) "-"COV_THRESHOLD")/1"}' | bc | xargs -I % test % -ge 0