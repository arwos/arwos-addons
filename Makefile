
.PHONY: install
install:
	go install go.osspkg.com/goppy/v2/cmd/goppy@latest
	goppy setup-lib

.PHONY: lint
lint:
	goppy lint

.PHONY: license
license:
	goppy license

.PHONY: build
build:
	goppy build --arch=amd64,arm64 --mode=plugin

.PHONY: tests
tests:
	goppy test

.PHONY: pre-commit
pre-commit: install lint tests build

.PHONY: ci
ci: pre-commit

