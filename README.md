# Home-Security-Project 
# ![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
This projetc is not ready to deploiement ![Home-Security-Project]

Projet de sécurisation de domicile

Ce projet utilise un esp32c3 pour le capteur de porte.

Le capteur de porte utilise le topic (mqtt) Porte_ouverte pour avertir d'une intrusion.

![LA stack des programmes, ](stack.png)

## Topics Mqtt
![Topic_Mqtt, ](mermaid-mqtt.svg)
Les topics sont accédés par les 2 sondes esp32, (porte et sirène), et par les 2 programmes GO

## Surveille_porte
![Premier programme, ](mermaid-diagram-pg1.svg)

On surveille le topic mqtt porte
Si la porte est ouverte et l'alarme armée, on positionne le topic mqtt alarme à 1.

## Surveille_Gotify_Messages
![Second programme, ](mermaid-diagram-pg2.svg)

On survelle les message reçu par Gotify, si un message "Alarme ON" est reçu, on positionne le topic mqtt alarm armée à 1
si un message "Alarme OFF" est reçu, onposition le topic mqtt alarm armée à 0

