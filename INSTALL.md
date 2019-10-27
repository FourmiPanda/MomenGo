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
2. Start REDIS Server : TODO
3. Start CAPTORS Publiser : : myproject/cmd/capteur/main.go
4. Start the REST API : : myproject/cmd/api/api.go