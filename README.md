Hello world application
============================
[![Build Status](https://travis-ci.org/ftob/ta.svg?branch=master)](https://travis-ci.org/ftob/ta)

Getting starting
-------------------
1. Install docker and docker-compose <br>
    1a. Docker CE: <br>
        CentOS - https://docs.docker.com/install/linux/docker-ce/centos/ <br>
        Ubuntu - https://docs.docker.com/install/linux/docker-ce/ubuntu/ <br>
        MacOS  - https://docs.docker.com/docker-for-mac/install/ <br>
    1b. docker-compose: https://docs.docker.com/compose/install/ <br>
    
2. Git clone project 
3. Build container - ```$ docker-compose build```  
4. Start containers - ```$ docker-compose up -d```

Details service
--------------------
**Application "Hello world"**

Application listening port 8080 on localhost http://localhost:8080. <br>

Endpoints: 
   - http://localhost:8080/ - root ("Hello world")
   - http://localhost:8080/metrics - Prometheus metrics
   - http://localhost:8080/service/v1/health - inner health-check application
   
**Goss server (https://github.com/aelsabbahy/goss)**

External server validation.

Endpoint:
   - http://localhost:8081/healthz


**Prometheus**

Monitoring server.

Endpoint: 
   - http://localhost:9090
   
   
**Grafana**

Grafana server - metrics dashboard.

Endpoint:
   - http://localhost:3000

Default login and password - *admin* and *admin*

Dashboard - application

Structure project
---------------------

```
   ├── .docker - Docker files directory
   │   ├── grafana 
   │   │   ├── Dockerfile
   │   │   └── etc
   │   │       └── grafana
   │   │           ├── dashboards
   │   │           │   ├── application.json - Dashboards
   │   │           │   └── dashboard.yaml - Dasshboard config
   │   │           └── datasources - Datasources
   │   │               └── datasource.yaml -Datasource config
   │   ├── healthcheck - Goss server
   │   │   ├── Dockerfile
   │   │   └── goss
   │   │       └── goss.yaml - Tests
   │   └── prometheus 
   │       ├── Dockerfile
   │       └── etc
   │           └── prometheus
   │               ├── alert.rules
   │               └── prometheus.yml - cconfig
   ├── docker-compose.yml - docker-compose config
   ├── Dockerfile - dockerfile application
   ├── health - Inner health-check 
   │   └── service.go 
   ├── index - Hello world service
   │   ├── instrumeting.go
   │   ├── logging.go
   │   ├── service.go
   │   └── service_test.go
   ├── main.go 
   ├── notallow - MethodNotAllow service
   │   ├── instrumeting.go
   │   ├── logging.go
   │   ├── service.go
   │   └── service_test.go
   ├── notfound - Not found service
   │   ├── instrumeting.go
   │   ├── logging.go
   │   ├── service.go
   │   └── service_test.go
   ├── README.md 
   └── server - HTTP server and endpoints
       ├── health.go
       ├── index.go
       ├── notallow.go
       ├── notfound.go
       └── server.go
```

