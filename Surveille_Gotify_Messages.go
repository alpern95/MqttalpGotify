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
#             : 
# v0.3        : 16/02/2023 Ajouter des fonctions pour la lisibilité
#             : 16/03/2023 creat fonction readappmess & delappmess
# v0.4        : 01/03/2023
#             : 01/03/2023 Nommage du programme main.go vers Surveille_Gotify_Messages.go
# v0.5        : 09/01/2024 Fichier de configuration dans /etc/Surveille_Gotify_Messages/
#             : 09/01/2024 Fichier de log dans /var/log/Surveille_Gotify_Messages/
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
        "encoding/json"
	//"io/ioutil"
        "os"

	//"github.com/gotify/go-api-client/v2/auth"
	"github.com/gotify/go-api-client/v2/client/message"
	"github.com/gotify/go-api-client/v2/gotify"
	"github.com/gotify/go-api-client/v2/models"
        "github.com/gotify/go-api-client/v2/auth"
        "github.com/gotify/go-api-client/v2/client/application" // ajout le 5/01/2024

        // puplish mqtt
        //https://github.com/lucacasonato/mqtt"
        "github.com/lucacasonato/mqtt"
        "time"
        "context"
)

type Config struct {
    Loglevel string `json:"level"`
}

type GetAppMessagesOK struct {
	Payload *models.PagedMessages
}

var config Config

var mois string
var jour string
var annee string
var maint time.Time
var heure string
var minute string
var seconde string

var pathlog string
var pathetc string
var filename string
var loglevel string

var messagesResponse GetAppMessagesOK

const (
	gotifyURL        = "http://localhost:80"
	applicationToken = "AcENJxTa-T-5c68"  //APP 0 AcENJxTa-T-5c68
        clientsToken = "CtIjH3i4GmiCwAP"      //Client  Linaro CtIjH3i4GmiCwAP Chrome CPIqN6zqPjPWEpI
)

func main() {
    start := time.Now()
    pathlog := "/var/log/Surveille_Gotify_Messages/"
    pathetc := "/etc/Surveille_Gotify_Messages/"
    config, _ := LoadConfiguration(pathetc+"config.json")
    InitTime()
    filename = pathlog+"Surveille_Gotify_Messages-"+jour +mois +annee+".log"
    loglevel = config.Loglevel
    InitLogs()
    log.Info("Loglevel : ",loglevel)
    log.Debug("========================= Début à : ",start)

    // Boucle infinie
    for true { 

      // Ajout readApp Lire les appli disponibles
      err := getapp()
      if    err != nil {
           log.Fatal ("FATAL appel readappmess",err)
       }

       // Lire les messages dans l'appli 
       err = readappmess()
       if    err != nil {
          log.Fatal ("FATAL appel readappmess",err)
       }

       // supprimer les messages lues (faire une fonction)
       err = delappmess()
       if err != nil { 
          log.Fatal ("FATAL appel delappmess",err) 
       }
       time.Sleep(2 * time.Second)
       //envoie notification
       end := time.Now()
       log.Debug("========================= Fin à : ",end)
     } // fin boucle infinie
}

// Ajout d'une fonction pour afficher les applications.
func getapp()(error){
    myURL, _ := url.Parse(gotifyURL)
    client := gotify.NewClient(myURL, &http.Client{}) // ok
    paramapps := application.NewGetAppsParams()
    listeapplication, err := client.Application.GetApps(paramapps,auth.TokenAuth(clientsToken)) // OK
    if err != nil {
        log.Fatalf("func getapp : Ne peut afficher les applications %v" , err)
        return err}
    log.Debug("func getapp : ",listeapplication)
    return err
    }

// Suppressions des messages
func delappmess()(error){
    myURL, _ := url.Parse(gotifyURL)
    client := gotify.NewClient(myURL, &http.Client{})
    paramdels := message.NewDeleteAppMessagesParams()
    paramdels.ID = 8
    messagesDelResponse, err := client.Message.DeleteAppMessages(paramdels,auth.TokenAuth(clientsToken)) // OK
    if err != nil {
        log.Fatalf("Could not get messages %v", err)
        return err
    }
    log.Debug("func delappmess: All Messages Deleted ",messagesDelResponse)
    return err
}

func ctx() context.Context {
	cntx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	return cntx
}

// fonction publish arme armée ajouté varible 1 ou 0
    func pubalarmearmee(valeur int)(error) {
        // Publier 1 dans topic alarme_armee  si le mesage est alarme on
        // Publier 0 dans topic alarme_armee  si le mesage est alarme off
        mqttclient, err := mqtt.NewClient(mqtt.ClientOptions{
                Servers: []string{"tcp://localhost:1883"},
        })
        if err != nil {
                log.Fatalf("failed to create mqtt client: %v\n", err)
        }

    log.Debug("func pubalarmearmee: NewMqttCient",mqttclient)
    log.Debug("print typeof mqtt client ",reflect.TypeOf(mqttclient))
        // connect to mqtt server
        err = mqttclient.Connect(ctx())
        if err != nil {
            log.Fatalf("failed to connect to mqtt server: %v\n", err)
        }else {log.Debug("connect to mqtt server: OK %v\n", err)}

    // ajout valeur variable 0 ou 1
    err = mqttclient.PublishJSON(ctx(), "alarme_armee/",valeur, mqtt.AtLeastOnce)
    if err != nil {
        panic(err)
    }else {log.Debug("func pubalarmearmee: Publish Value  %v\n", err)}
    return err
}

func readappmess()(error){
    myURL, _ := url.Parse(gotifyURL)
    // lire les messages recue
    client := gotify.NewClient(myURL, &http.Client{}) // ok
    params := message.NewGetAppMessagesParams() //Ajout App
    params.ID = 8                               // Id de l'application
    messagesResponse, err := client.Message.GetAppMessages(params,auth.TokenAuth(clientsToken)) // OK
    if err != nil {
        log.Info("func readappmess: Les parametres; ",params.ID)
        log.Fatalf("func readappmess: Could not get messages %v" , err)
        return err
    }
    // extraire les messages
    messages := messagesResponse.Payload
    log.Debug("func readappmess: Les messages ",messages)
    log.Debug("func readappmess: Les messages all ",messages.Messages)
    // calcul Taille
    log.Debug("func readappmess: Calcul du message Paging Size ",len(messages.Messages))
    mess := messages.Messages
    for _, Messages := range mess {      // Boucle for pour chaque messages dans l'appli ID 
        // traitez chaque message ici
        log.Info("func readappmess: Message ApplicationID:  ",Messages.ApplicationID)
        //log.Info("Message: ",Messages.Message)
        // traitez chaque message ici Si message alarme on alors pubalarmearmee()
        if Messages.Title=="alarme"{
            //traitement des messages alarme
            log.Debug("func readappmess: Titre du message: ",Messages.Title)
            switch Messages.Message{
            case "off":
                log.Debug("func readappmess: Message case off: ",Messages.Message)
            case "Off":
                err = pubalarmearmee(0)
                if err != nil {
                  panic(err)
                }else {log.Debug("func readappmess: Publish mqtt message Case Off: OK %v\n", err)} 
                log.Debug("func readappmess: Publish mqtt message Case Off: ",Messages.Message)
            case "on":
                pubalarmearmee(1)
                log.Debug("func readappmess: Publish mqtt Message case on: ",Messages.Message)
            case "On":
                pubalarmearmee(1)
                log.Debug("func readappmess: Publish mqtt Message case on: ",Messages.Message)
            default:
                log.Info("func readappmess: Publish mqtt Message case default: ",Messages.Message)
            }
        }
        log.Debug("func readappmess: Message à supprimer Apllication ID : ",Messages.ApplicationID)
        log.Debug("func readappmess: Message à supprimer Message ID: ",Messages.ID)
    }
	return err
}

func LoadConfiguration(filename string) (Config,error) {
    var config Config
    log.Debug("func LoadConfiguration: Le nom du fichier de configuration : ",filename)
    configFile, err := os.Open(filename)
    defer configFile.Close()
    if err != nil {
        log.Debug("func LoadConfiguration: Vérifier la presence d'un fichier de config",config,err)
    }
    jsonParser := json.NewDecoder(configFile)
    err = jsonParser.Decode(&config)
    if err != nil {
        log.Debug("func LoadConfiguration: Json Paser failed to read le fichier de configuration ",config,err)
    }
    return config, err
}

func InitTime() {
    maint = time.Now()
    annee = maint.Format("06")
    mois = maint.Format("01")
    jour = maint.Format("02")
    heure = maint.Format("15")
    minute = maint.Format("04")
    seconde = maint.Format("05")
}

func InitLogs() {
    //Initialisation des Logs
    Formatter := new(log.TextFormatter)
    Formatter.TimestampFormat = "02-01-2006 15:04:05"
    Formatter.FullTimestamp = true
    log.SetFormatter(Formatter)
    log.SetLevel(log.DebugLevel)  // reste en debug si il n'est pas change par la configuration

    // Lecture du fichier de log
    f, err := os.OpenFile(filename, os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0644)
    if err != nil {
        // Cannot open log file. Logging to stderr
        log.Fatal("Error Oppening Log File:", err)
    }else{
        log.SetOutput(f)
        log.Info("ouverture du fichier de log OK")
    }
    //defer f.Close()
    // Set Loglevel
    switch loglevel {
        case "debug":
                log.SetLevel(log.DebugLevel)
        case "info":
                 log.SetLevel(log.InfoLevel)
        case "warn":
                log.SetLevel(log.WarnLevel)
        case "error":
                log.SetLevel(log.ErrorLevel)
        }
}
