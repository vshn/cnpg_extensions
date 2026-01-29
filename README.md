# CNPG extensions repository

See the official documentation how to package extensions: https://cloudnative-pg.io/docs/1.28/imagevolume_extensions/#image-specifications

If the extension is available in the Debian repositories, it is straight forward to package the extensions.
See the `pg_repack` extension for an example.

For packages that need to be built from source, please refer to the extensions docs.

> [!warning] 
> Always use the same CNPG image as the base that you plan to deploy.
> That ensures that it will be compatible.
