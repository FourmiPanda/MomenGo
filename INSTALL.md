# Installation

### 1. Dependecies
Make sure you got Paho and Redigo installed :


##### Paho :
```shell
    go get github.com/eclipse/paho.mqtt.golang
```
##### Redigo :

```shell
    go get github.com/gomodule/redigo/redis
```

### 2. Environement Variables
Set GOROOT and GOPATH :

```shell
    GOPATH=YOUR_PROJECT_FOLDER
```
```shell
    GOROOT=YOUR_GO_INSTALLATION_FOLDER
```

### 3. Running
1. Start BROKER MQTT : myproject/cmd/broker/initBroker.go
2. Start REDIS Server : run redis_server.exe
3. Start Redis captors subscriber : myproject/cmd/mqttToRedis/mqttToRedis.go
4. Start CSV captors subscriber : myproject/cmd/csv/csv.go
5. Start CAPTORS Publisers : myproject/cmd/capteur/main.go
6. Start the REST API : myproject/cmd/api_swagger/main.go
7. Start the web view : myproject/cmd/http_server/main.go
