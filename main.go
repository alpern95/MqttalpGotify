/*
# ===============================================================================================
# Auteur : Alain Pernelle pour Projet "Home-Security-Project"
# Date creation : 25/01/2023
#
# Description : Gerer les messages de gotify pour activé l'alarme
#             : Objectif lire les messages avec titre Alarme et ON ou OFF
#
# Parametres  :  dans le fichier de configuration
#             :  paramètres par defaut pour le projet
#
#
#             : Versions
#             : ========
# v0.0        : 25/01/2023
# v0.1        :
# v0.2        :
# v0.3        :
#             : Bugs Issues
#             : ===========
#             :
*/

package main

import (
        "net/http"
        //"net/url"
        log "github.com/sirupsen/logrus"
)


func main() {
    log.Info("Info : demarrage ")
    /* http.PostForm("http://localhost:80/message?token=A9fTnaUlyZVyDO0",
        url.Values{"message": {"Alarme ON"}, "title": {"Alarme"}})
    */
    resp, err := http.Get("http://localhost:80/message?token=CbcCqwh5RMQEsOR")

    log.Info("Loglevel : ",resp) 
    if err != nil {
        log.Fatal(err)
    }else {log.Info("Sortie texte  : ",resp) }

}
