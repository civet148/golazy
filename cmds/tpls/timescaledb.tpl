docker run -d \
--name timescaledb \
-p 5432:5432 \
-e POSTGRES_PASSWORD=12345678 \
-v ./timescale_data:/home/postgres/pgdata/data \
timescale/timescaledb-ha:pg17