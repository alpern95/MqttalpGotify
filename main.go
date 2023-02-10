/*
# ===============================================================================================
# Auteur : Alain Pernelle pour Projet "Home-Security-Project"
# Date creation : 25/01/2023
#
# Description : Gerer les messages de gotify pour activer ou désactiver l'alarme
#             : Objectif lire les messages avec un titre Alarme et ON ou OFF
#             :  Si le alarm on, on positionne le topic MQTT "alarme_armée" 0 'non armée) ou 1 (armée)
#
# Parametres  :  dans le fichier de configuration
#             :  paramètres par defaut pour le projet
#
#             : Documentation API  https://pkg.go.dev/github.com/gotify/go-api-client/v2@v2.0.4#section-readme
#             : Versions
#             : ========
# v0.0        : 25/01/2023
# v0.1        : 02/02/2023 test lecture message Gotify en utilisant /gotify/go-api-client
# v0.2        : 07/02/2023 test publish mqtt topic alarme_armee
#             : utilisation de /https://github.com/lucacasonato/mqtt" 
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
	"github.com/gotify/go-api-client/v2/models"
        "github.com/gotify/go-api-client/v2/auth"

        // puplish mqtt
        //https://github.com/lucacasonato/mqtt"
        "github.com/lucacasonato/mqtt"
        "time"
        "context"
)

type GetAppMessagesOK struct {
	Payload *models.PagedMessages
}

var messagesResponse GetAppMessagesOK

const (
	gotifyURL        = "http://localhost:80"
	applicationToken = "CbcCqwh5RMQEsOR"  //applitoken A9fTnaUlyZVyDO0  clenttoken  CbcCqwh5RMQEsOR
)

func main() {
    log.Info("Début")
    myURL, _ := url.Parse(gotifyURL)


    // lire les messages recue
    client := gotify.NewClient(myURL, &http.Client{}) // ok 

    params := message.NewGetAppMessagesParams() //Ajout App
    params.ID = 1

    messagesResponse, err := client.Message.GetAppMessages(params,auth.TokenAuth(applicationToken)) // OK
    if err != nil {
        log.Fatalf("Could not get messages %v", err)
        return
    }
    log.Info("print version en real message ",messagesResponse)
    log.Info("print typeof version en real message ",reflect.TypeOf(messagesResponse))

    // Faire une fonction  avec les messages gotify en retour.

    // extraire les messages
    messages := messagesResponse.Payload
    log.Info("Les messages ",messages)

    log.Info("Les messages all ",messages.Messages)
    // calcul Taille
    log.Info("print Calcul du message Paging Size ",len(messages.Messages))

    mess := messages.Messages
    for _, Messages := range mess {
        // traitez chaque message ici
        log.Info("Message ApplicationID:  ",Messages.ApplicationID)
        //log.Info("Message: ",Messages.Message)
        // traitez chaque message ici Si message alarme on alors pubalarmearmee()
        if Messages.Title=="alarme"{
            //traitement des messages alarme 
            log.Info("Titre du message: ",Messages.Title)
            switch Messages.Message{
            case "off":
                log.Info("Message case off: ",Messages.Message)
            case "Off":
                 pubalarmearmee()
                log.Info("Message case off: ",Messages.Message)
                pubalarmearmee()
            case "on":
                pubalarmearmee()
                 log.Info("Message case on: ",Messages.Message)
            case "On":
                pubalarmearmee()
                log.Info("Message case on: ",Messages.Message)
            default:
                log.Info("Message case default: ",Messages.Message)
            }
        }
        log.Info("Message à supprimer Apllication ID : ",Messages.ApplicationID)
        log.Info("Message à supprimer Message ID: ",Messages.ID)
    }

    // supprimer les messages lues
    //faire new NewDeleteAppMessagesParams() 
    //params := message.NewGetAppMessagesParams() //Ajout App
    //params.ID = 1

    // DeleteAppMessages(params *DeleteAppMessagesParams, authInfo runtime.ClientAuthInfoWriter)
    paramdels := message.NewDeleteAppMessagesParams()
    paramdels.ID = 1
    messagesDelResponse, err := client.Message.DeleteAppMessages(paramdels,auth.TokenAuth(applicationToken)) // OK
    if err != nil {
        log.Fatalf("Could not get messages %v", err)
        return
    }
    // cannot use params (type *message.GetAppMessagesParams) as type *message.DeleteAppMessagesParams
    log.Info("Message Del ",messagesDelResponse)
    //pubalarmearmee()


    //envoie notification

}

// ajout pour mqtt
func ctx() context.Context {
	cntx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	return cntx
}

// fonction publish arme armée ajouté varible 1 ou 0
    func pubalarmearmee() {
        // Publier 1 dans topic alarme_armee  si le mesage est alarme on
        // Publier 0 dans topic alarme_armee  si le mesage est alarme off
        mqttclient, err := mqtt.NewClient(mqtt.ClientOptions{
                Servers: []string{"tcp://localhost:1883"},
        })
        if err != nil {
                log.Fatalf("failed to create mqtt client: %v\n", err)
        }

    log.Info("MQTT NewMqttCient",mqttclient)
    log.Info("print typeof mqtt client ",reflect.TypeOf(mqttclient))
        // connect to mqtt server
        err = mqttclient.Connect(ctx())
        if err != nil {
            log.Fatalf("failed to connect to mqtt server: %v\n", err)
        }else {log.Info("connect to mqtt server: OK %v\n", err)}

    err = mqttclient.PublishJSON(ctx(), "alarme_armee/",1, mqtt.AtLeastOnce)
    if err != nil {
        panic(err)
    }else {log.Info("Publish mqtt message: OK %v\n", err)}}
