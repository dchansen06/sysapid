.SILENT: sysapid
.DEFAULT: all

all: sysapid

sysapid:
	echo If this does have the intended installation effect, configure your GOBIN and GOPATH as appropriate
	echo For intended effect, open a new terminal and edit your enviornment variables as needed
	echo Review https://pkg.go.dev/cmd/go#hdr-Environment_variables for more information
	echo A reasonable default might be GOBIN="~/bin" and GOPATH="~/.go", and make sure $GOBIN is on $PATH
	go mod tidy
	go install sysapid.go
	echo Run sudo ~/bin/sysapid to start the application (or just ~/bin/sysapid if it is an open port)
