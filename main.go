package main

import (
	"echo-discord-bot/partials"
	"echo-discord-bot/utils"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

// Déclare la variable sess en tant que variable globale
var Sess *discordgo.Session
var Err error

var Entrees_id = "1147588275570233344"
var Sorties_id = "1147588295031795822"

func main() {
	// Chargez les variables d'environnement à partir du fichier .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erreur lors du chargement du fichier .env")
	}

	// Récupérez la clé Discord à partir des variables d'environnement
	botToken := os.Getenv("ECHO_BOT_TOKEN")

	// Initialisez votre session Discord avec la clé
	Sess, Err := discordgo.New("Bot " + botToken)
	if err != nil {
		log.Fatal(Err)
	}

	Sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	partials.ConnectBot(Sess)

	Err = Sess.Open()
	if Err != nil {
		log.Fatal(Err)
	}
	defer Sess.Close()

	fmt.Println("the bot is online !")

	utils.Utils(Sess)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
