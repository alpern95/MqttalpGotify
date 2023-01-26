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
	"log"
	"net/http"
	"net/url"

	"github.com/gotify/go-api-client/v2/auth"
	"github.com/gotify/go-api-client/v2/client/message"
	"github.com/gotify/go-api-client/v2/gotify"
	"github.com/gotify/go-api-client/v2/models"
)

const (
	gotifyURL        = "http://localhost:80"
	applicationToken = "A9fTnaUlyZVyDO0"
)

func main() {
	myURL, _ := url.Parse(gotifyURL)

        // Get version de gotify
	client := gotify.NewClient(myURL, &http.Client{})
	versionResponse, err := client.Version.GetVersion(nil)

	if err != nil {
		log.Fatal("Could not request version ", err)
		return
	}
	version := versionResponse.Payload
	log.Println("Found version", *version)

        // envoie d'un message
	params := message.NewCreateMessageParams()
	params.Body = &models.MessageExternal{
		Title:    "Alarm",
		Message:  "ON",
		Priority: 5,
	}
	_, err = client.Message.CreateMessage(params, auth.TokenAuth(applicationToken))

	if err != nil {
		log.Fatalf("Could not send message %v", err)
		return
	}
	log.Println("Message Sent!")

        // lire les messages recue

        // supprimer les messages lues
}
