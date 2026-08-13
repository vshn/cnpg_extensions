# Plpython3u

The `plpython3u` PostgreSQL extension allows writing functions and stored procedures in Python 3.
It is the untrusted variant of PL/Python, meaning it can access the file system and network from within the database.
For more information, see the [official documentation](https://www.postgresql.org/docs/current/plpython.html).

## Usage

### 1. Add the Plpython3u extension image to your Cluster

Define the `plpython3u` extension under the `postgresql.extensions` section of
your `Cluster` resource. For example:

```yaml
apiVersion: postgresql.cnpg.io/v1
kind: Cluster
metadata:
  name: cluster-plpython3u
spec:
  imageName: ghcr.io/cloudnative-pg/postgresql:18-minimal-bookworm
  instances: 1

  storage:
    size: 1Gi

  postgresql:
    extensions:
    - name: plpython3u
      image:
        # renovate: suite=bookworm-pgdg depName=postgresql-plpython3-18
        reference: ghcr.io/vshn/plpython3u:18.6-18-bookworm
      ld_library_path:
        - system
```

### 2. Enable the extension in a database

You can install `plpython3u` in a specific database by creating or updating a
`Database` resource. For example, to enable it in the `app` database:

```yaml
apiVersion: postgresql.cnpg.io/v1
kind: Database
metadata:
  name: cluster-plpython3u-app
spec:
  name: app
  owner: app
  cluster:
    name: cluster-plpython3u
  extensions:
  - name: plpython3u
    # renovate: suite=bookworm-pgdg depName=postgresql-plpython3-18
    version: '18.6'
```

### 3. Verify installation

Once the database is ready, connect to it with `psql` and run:

```sql
\dx
```

You should see `plpython3u` listed among the installed extensions.
