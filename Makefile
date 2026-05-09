.PHONY: release

cli:
	go build -o rex-cli cmd/main.go

release:
	@git fetch --tags
	@version=$$(git describe --tags $$(git rev-list --tags --max-count=1)); \
	echo "当前 tag: $$version"; \
	version_num=$${version#v}; \
	major=$$(echo $$version_num | cut -d. -f1); \
	minor=$$(echo $$version_num | cut -d. -f2); \
	patch=$$(echo $$version_num | cut -d. -f3); \
	int_version=$$(printf "%02d%02d%02d" $$major $$minor $$patch); \
	echo "六位整数版本号: $$int_version"; \
	new_int_version=$$((10#$$int_version + 1)); \
	echo "递增后整数版本号: $$new_int_version"; \
	new_int_version_fmt=$$(printf "%06d" $$new_int_version); \
	new_major=$$(echo $$new_int_version_fmt | cut -c1-2 | sed 's/^0*//'); \
	new_minor=$$(echo $$new_int_version_fmt | cut -c3-4 | sed 's/^0*//'); \
	new_patch=$$(echo $$new_int_version_fmt | cut -c5-6 | sed 's/^0*//'); \
	[ -z "$$new_major" ] && new_major=0; \
	[ -z "$$new_minor" ] && new_minor=0; \
	[ -z "$$new_patch" ] && new_patch=0; \
	new_version="v$${new_major:-0}.$${new_minor:-0}.$${new_patch:-0}"; \
	echo "🎉 新版本号: $$new_version"; \
	echo "$$new_version" > rex/version; \
	git tag "$$new_version"; \
	git push origin tag "$$new_version"; \