fmt:
	go fmt ./...

build:
	mkdir -p ./bin
	go build -o ./bin/mosho-ircmdd ./cmd
	cp serviceAccountKey.json ir_pattern.json ./bin

run: build
	cd bin && ./mosho-ircmdd

clean:
	go clean ./...
	rm -rf ./bin

test:
	go test ./...

remote-build:
	mkdir -p ./bin
	GOOS=linux GOARCH=arm GOARM=6 go build -o ./bin/mosho-ircmdd ./cmd

remote-copy: remote-build
	ssh raspi "mkdir -p services/mosho-ircmd/bin"
	scp bin/mosho-ircmdd ir_pattern.json serviceAccountKey.json raspi:services/mosho-ircmd/bin/
	scp mosho-ircmdd.service raspi:services/mosho-ircmd/

remote-install: remote-copy
	ssh raspi "cd services/mosho-ircmd && sudo cp mosho-ircmdd.service /etc/systemd/system && sudo systemctl enable mosho-ircmdd && sudo service mosho-ircmdd start"
