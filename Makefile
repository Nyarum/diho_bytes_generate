
install:
	cd cmd/diho_bytes_generate && go install

generate:
	go generate ./...

.PHONY: install generate