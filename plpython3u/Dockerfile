ARG BASE=ghcr.io/cloudnative-pg/postgresql:18-minimal-bookworm
FROM $BASE AS builder

ARG PG_MAJOR
ARG EXT_VERSION

USER 0

RUN set -eux && \
	# Initial system libraries
	ldconfig -p | awk '{print $NF}' | grep '^/' | sort | uniq > /tmp/base-image-libs.out && \
  # Install PL/Python3
  apt-get update && apt-get install -y --no-install-recommends \
      "postgresql-plpython3-${PG_MAJOR}=${EXT_VERSION}" && \
    rm -rf /var/lib/apt/lists/*

# Gather PL/Python3 system libraries and their licenses
RUN mkdir -p /system /licenses && \
	# Get libraries
	ldd /usr/lib/postgresql/${PG_MAJOR}/lib/plpython3.so \
    | awk '{print $3}' | grep '^/' | sort | uniq > /tmp/all-deps.out && \
  # Extract all the libs that aren't already part of the base image
	comm -13 /tmp/base-image-libs.out /tmp/all-deps.out > /tmp/libraries.out && \
	while read -r lib; do \
		resolved=$(readlink -f "$lib"); \
		dir=$(dirname "$lib"); \
		base=$(basename "$lib"); \
		# Copy the real file
		cp -a "$resolved" /system/; \
		# Reconstruct all its symlinks
		for file in "$dir"/"${base%.so*}.so"*; do \
			[ -e "$file" ] || continue; \
			# If it's a symlink and it resolves to the same real file, we reconstruct it
			if [ -L "$file" ] && [ "$(readlink -f "$file")" = "$resolved" ]; then \
				ln -sf "$(basename "$resolved")" "/system/$(basename "$file")"; \
			fi; \
		done; \
	done < /tmp/libraries.out && \
	# Get licenses
	for lib in $(find /system -maxdepth 1 -type f -name '*.so*'); do \
		# Get the name of the pkg that installed the library
		pkg=$(dpkg -S "$(basename "$lib")" | grep -v "diversion by" | awk -F: '/:/{print $1; exit}'); \
		[ -z "$pkg" ] && continue; \
		mkdir -p "/licenses/$pkg" && cp -a "/usr/share/doc/$pkg/copyright" "/licenses/$pkg/copyright"; \
	done

FROM scratch
ARG PG_MAJOR

# Licenses
COPY --from=builder /licenses /licences/
COPY --from=builder /usr/share/doc/postgresql-plpython3-${PG_MAJOR}/copyright /licenses/postgresql-plpython3-${PG_MAJOR}/

# Libraries
COPY --from=builder /usr/lib/postgresql/${PG_MAJOR}/lib/plpython3* /lib/

# Share
COPY --from=builder /usr/share/postgresql/${PG_MAJOR}/extension/plpython3u* /share/extension/

# System libs
COPY --from=builder /system /system

USER 65532:65532
