# MNC Test Stage 2
## Prerequisites
- Posgresql database
- Go Progamming Langgauge v1.23
- Redis
## How To Run
- Copy `config.example` and rename into `config.yaml` and fill the config
- run migration the set up the database or you can use `./database_dump/mnc-test-dump`
```bash
make migrate_up
```
- Run The Server (REST API) using this command
```bash
make run_dev
```
- Run The Worker (Async Process) using this command
```bash
make run_worker_dev
```
- And Finally play with the API using your API Client like postman, you can load the postman doc here `./database_dump/MNC-TEST.postman_collection.json`
## Unit Test
To Run the Unit Test you can use this command
```bash
make test
```
## Asynq Monitong
After you run the server you can access `http://localhost:4050/monitoring`

## TODO
- Increase code coverage by add more unit test.
- Implement Asynq Scheduler for handle transaction with `PENDING` status.
