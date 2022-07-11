SHELL := bash
MODULES = collections option result

generate-docs:
	for module in $(MODULES); do \
		echo "Generating docs for $$module"; \
		gomarkdoc ./$$module > $$module/README.md; \
		gomarkdoc ./$$module > docs/$$module.md; \
		echo ""; \
	done
