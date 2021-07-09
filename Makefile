include .bingo/Variables.mk

.PHONY: deps
deps: $(BINGO)
	$(BINGO) get
