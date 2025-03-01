Name: {{.serviceName}} # service name
Host: {{.host}} # service host
Port: {{.port}} # service port
Mode: dev  # service mode (dev/test/prod)
Timeout: 60s # service timeout duration

Orm:
    DSN: "mysql://root:123456@127.0.0.1:3306/test?charset=utf8mb4" # database source name
    Debug: false # open or close debug log
    Max: 150 # max database connections
    Idle: 5 # max idle database connections
