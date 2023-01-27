flowchart LR
    A[entrée] --> B{Porte ouverte?}
    B -->|non| E[Sortie]
    B -->|oui| D{Alarme armée?}
    D -->|oui| F
    D -->|non| E[sortie] 
    F[Mqtt topic, alarme passe à 1]  
    F --> |go| G[envoi notification gotify]
    G --> E[sortie]
