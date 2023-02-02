/*
# ===============================================================================================
# Auteur : Alain Pernelle pour Projet "Home-Security-Project"
# Date creation : 25/01/2023
#
# Description : Gerer les messages de gotify pour activer ou désactiver l'alarme
#             : Objectif lire les messages avec titre Alarme et ON ou OFF
#
# Parametres  :  dans le fichier de configuration
#             :  paramètres par defaut pour le projet
#
#             : Documentation API  https://pkg.go.dev/github.com/gotify/go-api-client/v2@v2.0.4#section-readme
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
        "reflect"
        //"fmt"
	//"log"
        log "github.com/sirupsen/logrus"
	"net/http"
	"net/url"

	//"github.com/gotify/go-api-client/v2/auth"
	"github.com/gotify/go-api-client/v2/client/message"
	"github.com/gotify/go-api-client/v2/gotify"
	//"github.com/gotify/go-api-client/v2/models"
        "github.com/gotify/go-api-client/v2/auth"
)

const (
	gotifyURL        = "http://localhost:80"
	applicationToken = "CbcCqwh5RMQEsOR"  //applitoken A9fTnaUlyZVyDO0  clenttoken  CbcCqwh5RMQEsOR
)

func main() {
        log.Info("Début")
	myURL, _ := url.Parse(gotifyURL)

        // lire les messages recue
    client := gotify.NewClient(myURL, &http.Client{}) // ok 
    log.Info("print client ",client)
    log.Info("print typeof client ",reflect.TypeOf(client))

    params := message.NewGetAppMessagesParams() //Ajout App
    params.ID = 1
    log.Info("print params NewGetAppMessagesParams ",params)
    log.Info("print typeof params ",reflect.TypeOf(params))

messagesResponse, err := client.Message.GetAppMessages(params,auth.TokenAuth(applicationToken)) // OK mais 401 GetApps
//versionResponse, err := client.Message.GetAppsMessages(params,auth.TokenAuth(applicationToken))
if err != nil {
    log.Fatalf("Could not get messages %v", err)
    return
}
log.Info("print version en real message ",messagesResponse)
log.Info("print typeof version en real message ",reflect.TypeOf(messagesResponse))


//log.Info("print message payload ",messagesResponse.Payload)
//resp, err := client.Messages.(params, nil) //get_app_messages_responses

//reader :=  client.CreateClientReader()
//resp, err := client.CreateClientReader(params, nil)
//if err != nil {
//    log.Fatalf("Could not get messages %v", err)
//    return
//}
//messages := resp.GetPayload()
//for _, message := range messages {
    // traitez chaque message ici
//}


        // supprimer les messages lues
        //envoie notification
}
