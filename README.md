# Home-Security-Project 
# ![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
This projetc is not ready to deploiement ![Home-Security-Project]

Projet de sécurisation de domicile

Ce projet utilise un esp32c3 pour le capteur de porte.

Le capteur de porte utilise le topic (mqtt) Porte_ouverte pour avertir d'une intrusion.

![LA stack des programmes, ](stack.png)


## Programme 1

![Premier programme, ](mermaid-diagram-pg1.svg)

Surveille le topic mqtt porte
Si la porte est ouverte et l'alarme armée, on positionne le topic mqtt alarme à 1.


## Programme 2
![Second programme, ](mermaid-diagram-pg2.svg)

On survelle les message reçu par Gotify, si un message "Alarm ON" est reçu, on positionne le topic mqtt alarm armée à 1
si un message "Alarm OFF" est reçu, onposition le topic mqtt alarm armée à 0

