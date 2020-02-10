
PROJECT_ROOT=$(CURDIR)
PROJECT=$(shell basename $(PROJECT_ROOT))

COVERDIR=$(PROJECT_ROOT)/.cover
COVERAGEFILE=$(COVERDIR)/cover.out
COVERAGEREPORT=$(COVERDIR)/report.html

GINKGO=go run github.com/onsi/ginkgo/ginkgo

coverage-run:
	@mkdir -p $(COVERDIR)
	@$(GINKGO) -r -covermode=count --cover --trace ./
	@echo "mode: count" > "${COVERAGEFILE}"
	@find ./* -type f -name *.coverprofile -exec grep -h -v "^mode:" {} >> "${COVERAGEFILE}" \; -exec rm -f {} \;

coverage: coverage-run
	@sed -i -e "s|_$(PROJECT_ROOT)/|./|g" "${COVERAGEFILE}"
	@cp "${COVERAGEFILE}" coverage.txt

coverage-gen-html:
	@go tool cover -html="${COVERAGEFILE}" -o $(COVERAGEREPORT)

coverage-html: coverage coverage-gen-html
	@xdg-open $(COVERAGEREPORT) 2> /dev/null > /dev/null

test:
	@$(GINKGO) -r --randomizeAllSpecs --randomizeSuites --failOnPending --cover --trace --race --progress

.PHONY: coverage coverage-run coverage-html