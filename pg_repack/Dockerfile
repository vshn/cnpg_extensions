FROM ghcr.io/cloudnative-pg/postgresql:18 AS source

USER root

RUN apt-get update && apt-get install postgresql-18-repack

FROM scratch

# https://packages.debian.org/sid/amd64/postgresql-18-repack/filelist
COPY --from=source /usr/lib/postgresql/18/lib/bitcode/pg_repack.index.bc /lib/pg_repack/
COPY --from=source /usr/lib/postgresql/18/lib/bitcode/pg_repack/pgut/pgut-spi.bc /lib/pg_repack/
COPY --from=source /usr/lib/postgresql/18/lib/bitcode/pg_repack/repack.bc /lib/pg_repack/
COPY --from=source /usr/lib/postgresql/18/lib/pg_repack.so /lib/

COPY --from=source /usr/share/postgresql/18/extension/pg_repack--1.5.3.sql /share/extension/
COPY --from=source /usr/share/postgresql/18/extension/pg_repack.control /share/extension/

COPY --from=source /usr/lib/postgresql/18/bin/pg_repack /bin/
