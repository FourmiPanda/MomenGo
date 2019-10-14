# MomenGo
 Système de collecte et de restitution de données météo des aéroports ( température, vitesse du vent, pression atmosphérique) 

## Structure des messages MQTT
```json
{
 "idCaptor":  "integer",
 "idAirport": "string",
 "measure":   "string",
 "values":     [
      {
       "value": "number",
       "time":  "timestamp (https://en.wikipedia.org/wiki/ISO_8601)"
      } 
    ]
}
```

Exemple :
```json
{
 "idCaptor":  1,
 "idAirport": "BIA",
 "measure":   "Temperature",
 "values":     [
       {
        "value": 27.8,
        "time":  "2007-03-01T13:00:00Z"
       },
       {
        "value": 21.9,
        "time":  "2008-03-01T13:00:00Z"
       }
     ]
}
```
```
    