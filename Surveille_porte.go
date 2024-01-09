/*
# ===============================================================================================
# Auteur : Alain Pernelle pour Projet "Home-Security-Project"
# Date creation : 09/01/2024
#
# Description : Surveiller le topic mqtt porte_ouverte
#             :  Si le topic porte_ouverte 0 , et si topic alarme_armee  0 Sortie
#             :  Si le topic porte_ouverte 0 , et si topic alarme_armee  1 Sortie 
#             :  Si le topic porte_ouverte 1 , et si topic alarme_armee	 0 Sortie
#             :  Si le topic porte_ouverte 1 , et si topic alarme_armee  1 Pub 1 dan topic alame  Sortie
#
# Parametres  :  dans le fichier de configuration
#             :  paramètres par defaut pour le projet
#
#             : Documentation API  https://pkg.go.dev/github.com/gotify/go-api-client/v2@v2.0.4#section-readme
#             : Versions
#             : ========
# v0.0        : 10/01/2024
# v0.1        :
# v0.2        :
#             :
# v0.3        :
#             :
# v0.4        :
#             :
# v0.5        :
#             :
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
        //"net/http"
        //"net/url"
        "encoding/json"
        //"io/ioutil"
        "os"
        // puplish mqtt
        //https://github.com/lucacasonato/mqtt"
        "github.com/lucacasonato/mqtt"
        "time"
        "context"
)

type Config struct {
    Loglevel string `json:"level"`
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

//var messagesResponse GetAppMessagesOK

const (
        gotifyURL        = "http://localhost:80"
        applicationToken = "AcENJxTa-T-5c68"  //APP 0 AcENJxTa-T-5c68
        clientsToken = "CtIjH3i4GmiCwAP"      //Client  Linaro CtIjH3i4GmiCwAP Chrome CPIqN6zqPjPWEpI
)

func main() {
    start := time.Now()
    pathlog := "/var/log/Surveille_porte/"
    pathetc := "/etc/Surveille_porte/"
    config, _ := LoadConfiguration(pathetc+"config.json")
    InitTime()
    filename = pathlog+"Surveille_Gotify_Messages-"+jour +mois +annee+".log"
    loglevel = config.Loglevel
    InitLogs()
    log.Info("Loglevel : ",loglevel)
    log.Debug("========================= Début à : ",start)

    // Boucle infinie
    for true {

      //souscrire au topic porte_ouverte
       err := subtopicporte()
       if    err != nil {
          log.Fatal ("FATAL appel fonction subtopicporte",err)
       }

       //envoie notification
       end := time.Now()
       log.Debug("========================= Fin à : ",end)
    } // fin boucle infinie
}

// Fonction subscribe topic
    func subtopicporte()(error) {
        // création du client (New)
        mqttclient, err := mqtt.NewClient(mqtt.ClientOptions{
                Servers: []string{"tcp://localhost:1883"},
        })
        if err != nil {
                log.Fatalf("func : subtopicporte : failed to create mqtt client: %v\n", err)
        }
        // connexion au serveur mqtt
        err = mqttclient.Connect(ctx())
        if err != nil {
            log.Fatalf("func : subtopicporte : failed to connect to mqtt server: %v\n", err)
        }else {log.Debug("func : subtopicporte : connect to mqtt server: OK %v\n", err)}

	err = mqttclient.Subscribe(ctx(), "alarme_armee/#", mqtt.AtMostOnce)
	if err != nil {
		log.Fatalf("func : subtopicporte : failed to subscribe to config service: %v\n", err)
	}
        return err
}

// fonction publish alarme variable 1 ou 0
    func pubalarme(valeur int)(error) {
        // Publier 1 ou 0 dans topic alarme
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
        err = mqttclient.PublishJSON(ctx(), "alarme/",valeur, mqtt.AtLeastOnce)
        if err != nil {
            panic(err)
        }else {log.Debug("func pubalarme: Publish Value  %v\n", err)}
        return err
}

func ctx() context.Context {
   cntx, _ := context.WithTimeout(context.Background(), 1*time.Second)
   return cntx
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
