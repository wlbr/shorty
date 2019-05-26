#LINKERFLAGS = -X main.Version=`git describe --tags --always --long --dirty` -X main.Buildstamp=`date -u '+%Y-%m-%d_%I:%M:%S_UTC'`
LINKERFLAGS = -X main.Buildstamp=`date -u '+%Y-%m-%d_%I:%M:%S_UTC'`
TARGET = shorty
PROJECTROOT = $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
DBPATH	=	$(PROJECTROOT)db/

define newline


endef

#for development I like to start postgress as a foreground process.
#and I like to use Iterm2 instead of the built-in OSX terminal.
define SCRIPT
	tell application "iTerm2"
		set dbWindow to (create window with default profile)
		tell current session of dbWindow
			set columns to 125
			set background color to {7500, 7500, 7500, 0}
			write text "postgres -D $(DBPATH)"
		end tell
	end tell
endef

all: clean build

.PHONY: clean
clean:
	rm -rf bin/
	rm -f $(TARGET)srv
	rm -f gotils/loglevel_string.go

generate:
	go generate gotils/loglevel.go

build: generate
	go build -ldflags "$(LINKERFLAGS)" -o $(TARGET)srv cmd/$(TARGET)srv/main.go

run: generate
	#go run -ldflags "$(LINKERFLAGS)" cmd/$(TARGET)/$(TARGET).go
	go run cmd/$(TARGET)srv/main.go -LogLevel All

debug: generate
	dlv debug cmd/$(TARGET)/$(TARGET).go -RunAsServer

buildcmd: clean generate
	mkdir -p bin
	for dir in cmd/* ; do \
		targetapp=$$(basename $$dir) ; \
		GOOS=linux GOARCH=amd64   go build -o bin/$${targetapp}_linux_x86_64 -ldflags "$(LINKERFLAGS)" cmd/$${targetapp}/main.go ; \
		#GOOS=darwin GOARCH=amd64  go build -o bin/$${targetapp}_osx_x86_64   -ldflags "$(LINKERFLAGS)" cmd/$${targetapp}/main.go ; \
		#GOOS=windows GOARCH=386   go build -o bin/$${targetapp}_win32        -ldflags "$(LINKERFLAGS)" cmd/$${targetapp}/main.go ; \
		#GOOS=windows GOARCH=amd64 go build -o bin/$${targetapp}_win64        -ldflags "$(LINKERFLAGS)" cmd/$${targetapp}/main.go ; \
	done

deploy: rbuild
	rsync -a --progress $(TARGET) $(HOST):./fcgi-bin/

dep:
	go get -u gopkg.in/gcfg.v1
	go get -u github.com/alvaroloes/enumer
	go get -u github.com/wlbr/templify
	go get -u github.com/gorilla/mux
	go get -u github.com/lib/pq


dbnew:
	rm -rf $(DBPATH)
	mkdir -p $(DBPATH)
	chmod 750 $(DBPATH)
	initdb -D $(DBPATH)

dbtermosx:
	#postgres -D $(DBPATH)
	#pg_ctl -D $(DBPATH) -l $(PROJECTROOT)db.log start
	#osascript -e 'tell app "Terminal" to do script "postgres -D $(DBPATH)"'

dbstart:
	echo '$(subst $(newline),\n,${SCRIPT})' > dbsstartcripttemp.scpt
	osascript dbsstartcripttemp.scpt
	rm -f dbsstartcripttemp.scpt

dbstop:
	pg_ctl stop -D $(DBPATH) -m fast

dbreset:
	psql postgres -f create-db.sql
	psql postgres -c 'CREATE EXTENSION "adminpack";'

dbrecreatetables:
	psql shorty -U shortyapp -f create-tables.sql

dbinit: dbreset dbrecreatetables


open:
	code . && open .
