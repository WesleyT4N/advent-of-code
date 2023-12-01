.PHONY: new

new:
	@count=$$(find ./ -maxdepth 1 -type d -name '[0-9][0-9]' | wc -l); \
	count=$$(($$count + 1)); \
	printf -v folder "%02d" $$count; \
	mkdir -p $$folder; \
	cp -r templates/* $$folder; \
