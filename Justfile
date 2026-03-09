local_registry := "registry.127.0.0.1.nip.io:8443"
arch := `uname -m | sed 's/x86_64/amd64/;s/aarch64/arm64/'`

# List available recipes
help:
    @just --list

# Build an extension image locally (native platform, loads into docker daemon)
# Usage: just build pg_repack
build ext:
    docker buildx bake \
      -f {{ ext }}/metadata.hcl \
      -f docker-bake.hcl \
      --set "*.platforms=linux/{{ arch }}" \
      --load

# Build and push to local OCI registry (multi-platform)
# Usage: just push-local pg_repack
push-local ext:
    registry={{ local_registry }} docker buildx bake \
      -f {{ ext }}/metadata.hcl \
      -f docker-bake.hcl \
      --push

# Scaffold a new extension directory from templates
# Usage: just scaffold pg_repack
# Use PACKAGE to override the Debian package pattern (default: postgresql-%version%-<name>)
# Use DISTROS/VERSIONS to override defaults
scaffold ext:
    go run ./generate \
      -name {{ ext }} \
      ${PACKAGE:+-package "$PACKAGE"} \
      ${SQL_NAME:+-sql-name "$SQL_NAME"} \
      ${DISTROS:+-distros "$DISTROS"} \
      ${VERSIONS:+-versions "$VERSIONS"}
