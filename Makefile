help:
	@echo "soil makefile available targets:"
	@echo "	run	- runs this application"
	@echo "	build	- creates soil binary"
	@echo "	install	- makes installation to build/"
	@echo "	clean	- removes soil binary and build/ directory"
	@echo "	fork	- creates changes in fork, required parameter: repo=%account/%repo"

fork:
	@echo "fork: replacing 'arapov/soil' with '$(repo)' in go sources..."
	@find . -type f -name '*.go' | xargs sed -i 's/arapov\/soil/$(subst /,\/,$(repo))/g'
	@sed -i 's/arapov\/soil/$(subst /,\/,$(repo))/g' go.mod
	@echo "fork: use 'git diff' to see the changes, 'git commit' to apply or 'git checkout -f' to revert"

install: build
	@mkdir -p build/
	@cp -vR examples build/examples
	@cp -vR view build/view
	@cp -vR asset build/asset
	@cp -v soil build/soil

build:
	go build -ldflags="-s -w"

run:
	@go run soil.go

.PHONY: clean
clean:
	@rm -f soil
	@rm -rf build/
