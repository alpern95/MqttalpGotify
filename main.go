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
        "fmt"
	//"log"
        log "github.com/sirupsen/logrus"
	"net/http"
	"net/url"

	//"github.com/gotify/go-api-client/v2/auth"
	//"github.com/gotify/go-api-client/v2/client/message"
	//"github.com/gotify/go-api-client/v2/gotify"
	//"github.com/gotify/go-api-client/v2/models"
	"github.com/gotify/client-go/v2"
)

const (
	gotifyURL        = "http://localhost:80"
	applicationToken = "A9fTnaUlyZVyDO0"
)

func main() {
        log.Info("Début")
	myURL, _ := url.Parse(gotifyURL)

        // lire les messages reçus sur mon serveur gotify 
	client := gotify.NewClient("http://localhost:8080", "A9fTnaUlyZVyDO0")
	msgs, err := client.GetMessages(0, 100)
	if err != nil {
		log.Fatalf("failed to get messages: %v", err)
	}
	for _, msg := range msgs {
		fmt.Println("Message:", msg.Message, "Priority:", msg.Priority, "Title:", msg.Title)
	}

        log.Info("print message ",resp)

    //faire getmessage.  func (*Client) GetMessages  
    // func (a *Client) GetMessages(params *GetMessagesParams, authInfo runtime.ClientAuthInfoWriter) (*GetMessagesOK, error)
    log.Info("print params Body ",params.Body)
    //log.Info("print message ",Message.params)
    //message := client.NewCreateMessageParams()
    //resp := client.Messages.GetMessages()
    //message := client.NewCreateMessageParams()  //
    //versionResponse, err := client.Messages.GetAppMessagesParameters(nil) // essaie de creer les messages parameter avant
    //	if err != nil {
    //		log.Fatal("Could not request version ", err)
    // 		return
    //	}

    //version := versionResponse.Payload
    //log.Println("Found version", *version)
    //log.Info("print versionResponse ",versionResponse)
//fmt.Println(res)

////////////////////////////////////////////////
        // envoie d'un message (testé marche)
        //  à utiliser dans le cas de réception d'un message Alarme ON ou OFF
        //params := message.NewCreateMessageParams()
        //params.Body = &models.MessageExternal{
        //        Title:    "Alarm",
        //        Message:  "ON",
        //        Priority: 5,
        //}
        //_, err := client.Message.CreateMessage(params, auth.TokenAuth(applicationToken))

        //if err != nil {
        //        log.Fatalf("Could not send message %v", err)
        //        return
        //}
        //log.Println("Message Sent!")

        // supprimer les messages lues
}
