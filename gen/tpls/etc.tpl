Name: {{.serviceName}} # service name
Host: {{.host}} # service host
Port: {{.port}} # service port
Mode: dev  # service mode (dev/test/prod)
Timeout: 60s # service timeout duration

Orm:
    DSN: "mysql://root:12345678@127.0.0.1:3306/test?charset=utf8mb4" # database source name
    Debug: false # open or close debug log
    MaxConns: 150 # max database connections
    IdleConns: 5 # max idle database connections
    NodeId: 1 # snowflake node id
