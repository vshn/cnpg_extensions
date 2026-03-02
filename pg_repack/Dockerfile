ARG PG_VERSION=18
ARG EXT_VERSION=1.5.3

FROM ghcr.io/cloudnative-pg/postgresql:${PG_VERSION} AS source

ARG PG_VERSION=18
ARG EXT_VERSION=1.5.3

USER root

RUN apt-get update && apt-get install -y --no-install-recommends \
    postgresql-${PG_VERSION}-repack \
  && rm -rf /var/lib/apt/lists/*

FROM scratch

ARG PG_VERSION=18
ARG EXT_VERSION=1.5.3

# https://packages.debian.org/sid/amd64/postgresql-18-repack/filelist
COPY --from=source /usr/lib/postgresql/${PG_VERSION}/lib/bitcode/pg_repack.index.bc /lib/pg_repack/
COPY --from=source /usr/lib/postgresql/${PG_VERSION}/lib/bitcode/pg_repack/pgut/pgut-spi.bc /lib/pg_repack/
COPY --from=source /usr/lib/postgresql/${PG_VERSION}/lib/bitcode/pg_repack/repack.bc /lib/pg_repack/
COPY --from=source /usr/lib/postgresql/${PG_VERSION}/lib/pg_repack.so /lib/

COPY --from=source /usr/share/postgresql/${PG_VERSION}/extension/pg_repack--${EXT_VERSION}.sql /share/extension/
COPY --from=source /usr/share/postgresql/${PG_VERSION}/extension/pg_repack.control /share/extension/

COPY --from=source /usr/lib/postgresql/${PG_VERSION}/bin/pg_repack /bin/
