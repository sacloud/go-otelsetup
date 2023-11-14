#====================
AUTHOR         ?= The sacloud/go-otelsetup Authors
COPYRIGHT_YEAR ?= 2022-2023

BIN            ?= go-otelsetup
GO_FILES       ?= $(shell find . -name '*.go')

include includes/go/common.mk
#====================

default: $(DEFAULT_GOALS)
tools: dev-tools
