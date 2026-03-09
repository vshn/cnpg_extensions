# Pg Repack

The `pg_repack` PostgreSQL extension lets you reclaim storage and reduce table/index bloat
without holding an exclusive lock. It is a drop-in replacement for `CLUSTER` and `VACUUM FULL`
that works online. For more information, see the [official documentation](https://reorg.github.io/pg_repack/).

## Usage

### 1. Add the Pg Repack extension image to your Cluster

Define the `pg_repack` extension under the `postgresql.extensions` section of
your `Cluster` resource. For example:

```yaml
apiVersion: postgresql.cnpg.io/v1
kind: Cluster
metadata:
  name: cluster-pg-repack
spec:
  imageName: ghcr.io/cloudnative-pg/postgresql:18-minimal-bookworm
  instances: 1

  storage:
    size: 1Gi

  postgresql:
    extensions:
    - name: pg_repack
      image:
        # renovate: suite=bookworm-pgdg depName=postgresql-18-repack
        reference: ghcr.io/vshn/pg_repack:1.5.3-18-bookworm
```

### 2. Enable the extension in a database

You can install `pg_repack` in a specific database by creating or updating a
`Database` resource. For example, to enable it in the `app` database:

```yaml
apiVersion: postgresql.cnpg.io/v1
kind: Database
metadata:
  name: cluster-pg-repack-app
spec:
  name: app
  owner: app
  cluster:
    name: cluster-pg-repack
  extensions:
  - name: pg_repack
    # renovate: suite=bookworm-pgdg depName=postgresql-18-repack
    version: '1.5.3'
```

### 3. Verify installation

Once the database is ready, connect to it with `psql` and run:

```sql
\dx
```

You should see `pg_repack` listed among the installed extensions.
