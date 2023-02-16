# Home-Security-Project 
# ![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
This projetc is not ready to deploiement ![Home-Security-Project]

## La stack 
![LA stack des programmes, ](stack.png)

##Projet de sécurisation de domicile

Ce projet utilise un esp32c3 pour le capteur de porte et un autre pour la Sirène.

Le capteur de porte utilise le topic (mqtt) Porte_ouverte pour avertir que la porte est ouverte.

La sirène utilise le le topic (mqtt) alarme, si alarme alors la sirène s'active.

Deux programme Go 
 - On surveille le topic mqtt porte
Si la porte est ouverte et l'alarme armée, on positionne le topic mqtt alarme à 1.

 - On surveille les messages reçu par Gotify, si un message "Alarme ON" est reçu, on positionne le topic mq$
si un message "Alarme OFF" est reçu, onposition le topic mqtt alarm armée à 0

  - Un autre programme montor basé sur Janitor 
surveille la disponibilité des sondes esp32
  [Lien vers Janitor ](https://github.com/a-bali/janitor)

## Topics Mqtt

| Nom du Topic  | Propos        | Valeur  | Signification    | 
| ------------- |:-------------:| -------:| ----------------:|
| porte         | monitoring    |    1    | 1 up and working |
| sirene        | monitoring    |    1    | 1 up and working |
| porte_ouverte | sensor        |  0 ou 1 | 1 porte ouverte  |
| alarme        | sensor        |  0 ou 1 | 1 alarme activé  |
| alarme_armee  | sensesor      |  0 ou 1 | 1 alarme armée   |

![Topic Mqtt, ](mermaid-mqtt.svg)


Les topics sont accédés par les 2 sondes esp32, (porte et sirène), et par les 2 programmes GO

## Sonde esp32c3 porte
![ESP32C3_Porte, ](sonde_porte.svg)

## Sonde esp32c3 Sirène
![ESP32C3_Sirene, ](sonde_mqtt_sirene.svg)

## Surveille_porte
![Surveille_porte, ](mermaid-diagram-pg1.svg)

## Surveille_Gotify_Messages
![Second programme, ](mermaid-diagram-pg2.svg)


## Test Surveille_Gotify_Messages

Sur un l'envoie de message alarme On ou alarme Off via l'appli sur mon téléphone.
Le programme lit bien les messages et les suprime ensuite et selon le message, 
il positionne bien le topic alarme_armee à 0 ou 1
