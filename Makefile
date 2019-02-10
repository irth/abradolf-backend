yaml_files := $(shell find src -name '*.yaml')

index.html: ${yaml_files}
	redoc-cli bundle src/index.yaml -o index.html

.PHONY: serve
serve:
	redoc-cli serve --watch src/index.yaml
