CREATE TABLE "readings"(
     reading_time TIMESTAMP WITH TIME ZONE NOT NULL
    , metric_id INTEGER
    , machine_id INTEGER
    , value NUMERIC
);
SELECT create_hypertable('readings', 'reading_time',);

CREATE TABLE "metrics"(
    metric_id   INTEGER,
    name        TEXT
);
INSERT INTO metrics(metric_id, name) VALUES
(1, 'engine_temperature'),
(2, 'oil_temperature'),
(3, 'oil_presssure'),
(4, 'running_hours'),
(5, 'engine_load');

