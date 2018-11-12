all:
	go build -o ../../bin/example .

mac:
	GOOS=darwin GOARCH=amd64 go build -o ../../bin/example-darwin-amd64 .

linux:
	GOOS=linux GOARCH=386 go build -o ../../bin/example-linux-386 .

linux64:
	GOOS=linux GOARCH=amd64 go build -o ../../bin/example-linux-amd64 .

windows:
	GOOS=windows GOARCH=386 go build -o ../../bin/example-windows-386.exe .

windows64:
	GOOS=windows GOARCH=amd64 go build -o ../../bin/example-windows-amd64.exe .
