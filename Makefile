SHELL := bash
MODULES = collections iterator option result

generate-docs:
	for module in $(MODULES); do \
		echo "Generating docs for $$module"; \
		gomarkdoc ./$$module > docs/$$module.md; \
		pandoc docs/$$module.md --toc --metadata title="Go Rust Std - $$module Docs" -c https://unpkg.com/sakura.css/css/sakura.css --self-contained -o docs/$$module.html; \
		rm docs/$$module.md; \
		echo ""; \
	done
