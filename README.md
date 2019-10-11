# MomenGo
 Système de collecte et de restitution de données météo des aéroports ( température, vitesse du vent, pression atmosphérique) 

## Structure des messages MQTT
```json
{
 "idCaptor":  integer,
 "idAirport": string,
 "measure":   string,
 "value":     number,
 "timestamp": timestamp (https://en.wikipedia.org/wiki/ISO_8601)
}
```

Exemple :
```json
{
 "idCaptor":  1,
 "idAirport": "BIA",
 "measure":   "Temperature",
 "value":     27.6,
 "timestamp": "2007-03-01T13:00:00Z"
}
```
```
    