help:
	@echo "soil: makefile available targets"
	@echo
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

vendor:
	@go mod vendor

deps:
	mkdir -p deps/bootstrap
	@wget -nv --show-progress $(shell curl -s https://api.github.com/repos/twbs/bootstrap/tags | jq -r ".[0].tarball_url") -O deps/bootstrap-latest.tar
	@tar xf deps/bootstrap-latest.tar --strip 1 -C deps/bootstrap
	@mkdir -p assets/js assets/css assets/scss/bootstrap
	@cp deps/bootstrap/dist/js/bootstrap.min.* assets/js
	@mv deps/bootstrap/scss/* assets/scss/bootstrap/
	@cp -R contrib/scss/* assets/scss
	sass --trace assets/scss/main.scss:assets/css/soil.css
	@cp -R contrib/favicon assets

build: vendor deps
	go build -ldflags="-s -w" soil.go

install: build
	@mkdir -p build/contrib
	@cp -R contrib/tls contrib/soil.yaml contrib/HOWTO build/contrib/
	@cp -R view build/view
	@cp -R assets build/assets
	@cp -v soil build/soil

run: vendor deps
	go run soil.go

clean:
	rm -f soil go.sum
	rm -rf build/
	rm -rf deps/ assets/
	rm -rf .sass-cache/
