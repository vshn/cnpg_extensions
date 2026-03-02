local_registry := "registry.127.0.0.1.nip.io:8443"

# List available recipes
help:
    @just --list

# Build an extension image locally
# Usage: just build pg_repack 18 1.5.3
build ext pg_version ext_version:
    docker build \
      --build-arg PG_VERSION={{ pg_version }} \
      --build-arg EXT_VERSION={{ ext_version }} \
      -t {{ ext }}:pg{{ pg_version }}-{{ ext_version }} \
      -f {{ ext }}/Dockerfile \
      .

# Build and push to local OCI registry
# Usage: just push-local pg_repack 18 1.5.3
push-local ext pg_version ext_version: (build ext pg_version ext_version)
    docker tag {{ ext }}:pg{{ pg_version }}-{{ ext_version }} {{ local_registry }}/{{ ext }}:pg{{ pg_version }}-{{ ext_version }}
    docker push {{ local_registry }}/{{ ext }}:pg{{ pg_version }}-{{ ext_version }}
