# timeseries-playground

Test environment to play around with timeseries databases

## go-timeseries

Playing around with a CLI in go. Created a script to generate random data. Usage:

```bash
cd go-timeseries
go build

./timeseries --help
```

## docker-compose

Created a docker-compose with some databases connected to Grafana. Usage (already populated):

```bash
# grafana credentials: admin - admin123
docker-compose up -d grafana
```

## databases

### postgres

Postgres was meant as a reference against "real" timeseries databases.

#### population

```
$(timeseries connect postgres)
# psql (12.2 (Ubuntu 12.2-2.pgdg19.10+1))
# Type "help" for help.

# zlab=#
\copy readings from ./data/sth.csv CSV HEADER;
```

### timescale

Due to familiarity with SQL timescale was chosen as first test subject.

#### population

```
$(timeseries connect timescale)
# psql (12.2 (Ubuntu 12.2-2.pgdg19.10+1))
# Type "help" for help.

# zlab=#
\copy readings from ./data/sth.csv CSV HEADER;
```

### InfluxDB

Influx was chosen due to overall popularity.

#### population

The population takes forever, but that is what I got in the end...

Using telegraf:

The plugin tail was used, but a new feature was required. Due to this a custom Dockerfile
with the nightly build was created.

```bash
docker-compose up -d influx
# wait a bit for influx to start
docker-compose up telegraf
# watch logs until all metrics have been written
```
