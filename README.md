# gator-rss-go

## Run locally

You will need `Postgres` and `Go` installed in your machine.

See these [INSTALLATION DOCS](./INSTALLATIONS.md)

### Install Gator

You will install the cli program with Go.
Run:

```bash
go install github.com/agustin-carnevale/gator-rss-go@latest
```

### Create Gator DB
Once you have postgres installed, you need to create the DB for our gator program.

Enter the psql shell with:
```bash
psql postgres
```
Then 
```bash
postgres=# CREATE DATABASE gator;
```

### Create config file

Now, you will need to create a config json file.

```bash
touch ~/.gatorconfig.json
```

The file will contain the database connection string, should look like this:

```json
{
  "db_url": "protocol://username:password@host:port/database"
}
```
Fill with your own data.
Is probably going to be something like this:
`postgres://your-username:@localhost:5432/gator`

Check that you have the correct connection string:
```bash
psql "postgres://your-username:@localhost:5432/gator"
```
It should connect you to the gator database directly.


## Using Gator

Start by registering a user
```bash
➜ gator-rss-go register your-name
```

Then you can add a feed for gator to aggregate:
```bash
➜ gator-rss-go addfeed website-rss-url
```

Tell gator how often to aggregate/update data from rss sources (feeds)
```bash
➜ gator-rss-go agg [5s|1m|10m|etc]
```


You can now start using the browse cmd to see updated posts from your feeds (and other feeds you are following):
```bash
➜ gator-rss-go browse
```
