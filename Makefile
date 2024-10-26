SHELL := /bin/bash
VERSION ?= $(shell git describe --match 'v[0-9]*' --tags --always)

changelog_args = -o CHANGELOG.md -tag-filter-pattern '^v'

build_loc = ./bin/app
build_args = -ldflags "-s -w -X 'github.com/bagastri07/gin-boilerplate-service/internal/app/bootstrap.Version=$(VERSION)'" -o $(build_loc) main.go
build=go build $(build_args)
check_build = \
	if [ ! -f $(build_loc) ]; then \
		echo "Binary not found, building..."; \
		$(build); \
	fi

migrate_up = $(build_loc) migration --action=up
migrate_down = $(build_loc) migration --action=down

.PHONY: build
build:
	$(build)

.PHONY: run
run:
	$(build_loc) server

.PHONY: run_worker
run_worker:
	$(build_loc) worker

.PHONY: run_dev
run_dev:
	$(build) && $(build_loc) server

.PHONY: run_worker_dev
run_worker_dev:
	$(build) && $(build_loc) worker

.PHONY: migrate_up
migrate_up:
	@$(check_build); \
	if [ "$(version)" = "" ]; then \
		if $(migrate_up); then \
			echo "Migration applied successfully"; \
		else \
			echo "Migration failed."; \
		fi; \
	else \
		if $(build_loc) migration --action=up-to --version=$(version); then \
			echo "Migration applied successfully"; \
		else \
			echo "Migration failed."; \
		fi; \
	fi

.PHONY: migrate_down
migrate_down:
	@$(check_build); \
	if [ "$(version)" = "" ]; then \
		if $(migrate_down); then \
			echo "Migration applied successfully"; \
		else \
			echo "Migration failed."; \
		fi; \
	else \
		if $(build_loc) migration --action=down-to --version=$(version); then \
			echo "Migration applied successfully"; \
		else \
			echo "Migration failed."; \
		fi; \
	fi

.PHONY: migrate_create
migrate_create:
	@$(check_build); \
	if [ "$(name)" = "" ]; then \
		echo 'Migration file needs a name'; \
	else \
		if $(build_loc) migration --action=create --name=$(name); then \
			echo "Migration created successfully"; \
		else \
			echo "Migration creation failed."; \
		fi; \
	fi

.PHONY: clean_mock
clean_mock:
	find . -type d -name "mock" -exec rm -rf {} +

.PHONY: mock_only
mock_only: ; $(info $(M) generating mock...) @
	@./script/mockgen.sh

.PHONY: mock
mock: clean_mock mock_only

.PHONY: changelog
changelog:
ifdef version
	$(eval changelog_args=--next-tag $(version) $(changelog_args))
endif
	git-chglog $(changelog_args)

.PHONY: lint
lint:
	golangci-lint run --print-issued-lines=false --exclude-use-default=false --fix --timeout=3m

.PHONY: test_only
test_only: ; $(info $(M) start unit testing...) @
	@go test $$(go list ./... | grep -v /mocks/) -gcflags=all=-l --race -v -short -coverprofile=coverage.out
	@echo "\n*****************************"
	@totalCoverage=$$(go tool cover -func coverage.out | awk '/total:/ {print substr($$3, 1, length($$3)-1)}'); \
	echo "**  TOTAL COVERAGE: $$totalCoverage%  **"; \
	echo "*****************************\n"; \

.PHONY: test
test: lint test_only

.PHONY: coverage
coverage:
	go test ./... -gcflags=all=-l -coverprofile=coverage.out && go tool cover -html=coverage.out
