SHELL := bash
MODULES = collections iterator option result

generate-docs:
	for module in $(MODULES); do \
		echo "Generating docs for $$module"; \
		gomarkdoc ./$$module > docs/$$module.md; \
		echo ""; \
	done
