# Graph database for public IPs

### Initialization of the database

```bash
docker run \
    --name testneo4j \
    -p7474:7474 -p7687:7687 \
    -d \
    -v $HOME/neo4j/data:/data \
    -v $HOME/neo4j/logs:/logs \
    -v $HOME/neo4j/import:/var/lib/neo4j/import \
    -v $HOME/neo4j/plugins:/plugins \
    --env NEO4J_AUTH=<admin>/<passwd> \
    neo4j:latest
```

### Setup environment variables

```bash
export DBADMIN=<admin>
export DBPASSW=<passwd>
export DBURI=neo4j://<dburi>
```

### Execute API

```bash
go run cmd/api/main.go
```
