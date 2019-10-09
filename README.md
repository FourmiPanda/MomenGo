# MomenGo
 Système de collecte et de restitution de données météo des aéroports ( température, vitesse du vent, pression atmosphérique) 

## Structure des messages MQTT
```json
{
    "temp" : [{
        "id" : "captor1",
        "tempValue" : 1
    }],
    "wind" : [{
        "id" :  "captor2",
        "windValue" : 1
    }],
    "pressure" : [{
        "id" :  "captor3",
        "pressureValue" :  1
    }]
}    
```
    