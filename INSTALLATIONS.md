
### Go

You can download Go toolchain at
[https://go.dev/doc/install](https://go.dev/doc/install)

You can check the installation was successful by running:
```
 go version
 ```
 it should print the version.


 ### Postgres

For Mac

```bash
brew install postgresql@15
```

Ensure the installation worked. The psql command-line utility is the default client for Postgres. Use it to make sure you're on version 15+ of Postgres:

```bash
psql --version
```

Start the Postgres server in the background:

```bash
brew services start postgresql@15
```

You can use `psql` it's the "default" client for Postgres.

```bash
psql postgres
```

