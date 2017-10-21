trun:
	make test
	go run main.go
run:
	go run main.go
test:
	go test -v ./...
govendor:
	govendor fetch +out
	git add vendor/
	git commit -m "[MODIFY]update vendor"
	git push