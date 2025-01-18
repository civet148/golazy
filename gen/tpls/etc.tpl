Name: {{.serviceName}}
Host: {{.host}}
Port: {{.port}}
Mode: dev # dev or prod
Timeout: 60s

Orm:
    DSN: "mysql://root:123456@127.0.0.1:3306/test?charset=utf8mb4"
    Debug: false
    Max: 150
    Idle: 5
