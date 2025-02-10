CLI_ARCHIVE_NAME="ubctl"
SHELL = /bin/bash -e

.PHONY: version
version:
	@nix run .#ubctl version

.PHONY: nixdev
nixdev:
	@nix develop .#dev

.PHONY: nixfmt
nixfmt:
	@nix fmt

.PHONY: upgrade
upgrade:
	@go get -x -v -u ./... && make tidy

.PHONY: install
install:
	@go get -x -v ./... && make tidy

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: clean
clean:
	@go clean -x -i -r -cache -modcache
	@rm -rf ./dist

.PHONY: build
build:
	@goreleaser build --snapshot --verbose --clean

.PHONY: tag
tag:
	@git tag -f "$$(go run ./src/main.go version)"

.PHONY: release.local
release.local:
	@\
		CLI_ARCHIVE_NAME="$(CLI_ARCHIVE_NAME)" \
		SKIP_GITHUB="true" \
		SKIP_DOCKER="true" \
		goreleaser release \
			--snapshot \
			--verbose \
			--clean

.PHONY: release.github.strict
release.github.strict: tag
	@\
		CLI_ARCHIVE_NAME="$(CLI_ARCHIVE_NAME)" \
		SKIP_GITHUB="false" \
		SKIP_DOCKER="true" \
		goreleaser release \
			--skip=docker \
			--verbose \
			--clean

.PHONY: release.github
release.github: tag
	@\
		CLI_ARCHIVE_NAME="$(CLI_ARCHIVE_NAME)" \
		SKIP_GITHUB="false" \
		SKIP_DOCKER="true" \
		goreleaser release \
			--skip=validate,docker \
			--verbose \
			--clean

.PHONY: release.docker.strict
release.docker.strict: tag
	@\
		CLI_ARCHIVE_NAME="$(CLI_ARCHIVE_NAME)" \
		SKIP_GITHUB="true" \
		SKIP_DOCKER="false" \
		goreleaser release \
			--verbose \
			--clean

.PHONY: release.docker
release.docker: tag
	@\
		CLI_ARCHIVE_NAME="$(CLI_ARCHIVE_NAME)" \
		SKIP_GITHUB="true" \
		SKIP_DOCKER="false" \
		goreleaser release \
			--skip=validate \
			--verbose \
			--clean

.PHONY: release.all.strict
release.all.strict: tag
	@\
		CLI_ARCHIVE_NAME="$(CLI_ARCHIVE_NAME)" \
		SKIP_GITHUB="false" \
		SKIP_DOCKER="false" \
		goreleaser release \
			--verbose \
			--clean

.PHONY: release.all
release.all: tag
	@\
		CLI_ARCHIVE_NAME="$(CLI_ARCHIVE_NAME)" \
		SKIP_GITHUB="false" \
		SKIP_DOCKER="false" \
		goreleaser release \
		  --skip=validate \
			--verbose \
			--clean

.PHONY: docker
docker: release.local
	@IMG="$$(jq -erc --arg arch "$$(go env GOARCH)" '.[] | select(.type | contains("Docker Image")) | select(.name | contains($$arch)) | .name' ./dist/artifacts.json)" && \
	  docker run --rm -it --entrypoint /bin/bash "$$IMG"
