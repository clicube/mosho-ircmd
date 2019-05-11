fmt:
	go fmt ./...

build:
	mkdir -p ./bin
	go build -o ./bin/mosho-cmdsubd ./cmd
	cp serviceAccountKey.json onCmd.sh ./bin

clean:
	go clean ./...
	rm -rf ./bin

test:
	go test ./...

remote-build:
	mkdir -p ./bin
	GOOS=linux GOARCH=arm GOARM=6 go build -o ./bin/mosho-cmdsubd ./cmd

remote-copy: remote-build
	ssh raspi "mkdir -p services/mosho-cmdsub/bin"
	scp ./bin/mosho-cmdsubd raspi:services/mosho-cmdsub/bin/mosho-cmdsubd

remote-run: remote-copy
	ssh raspi "cd services/mosho-cmdsub/bin && ./mosho-cmdsubd"
