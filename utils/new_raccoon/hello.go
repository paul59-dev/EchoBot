package newraccoon

import (
	"log"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
)

func Hello(s *discordgo.Session) {
	// Lire le contenu du fichier bienvenue.txt
	welcomeMessage, err := os.ReadFile(fileHelloMD)
	if err != nil {
		log.Fatal("Erreur lors de la lecture du fichier bienvenue.md:", err)
	}

	// Envoyer le contenu du fichier bienvenue.txt sur Discord
	welcomeChannelID := "1146958639844823095" // salon bienvenue
	// Supprimer tous les messages du salon avant d'envoyer le nouveau message
	messages, err := s.ChannelMessages(welcomeChannelID, 100, "", "", "")
	if err != nil {
		log.Println("Erreur lors de la récupération des messages:", err)
	} else {
		for _, msg := range messages {
			err := s.ChannelMessageDelete(welcomeChannelID, msg.ID)
			if err != nil {
				log.Println("Erreur lors de la suppression du message:", err)
			}
			time.Sleep(1 * time.Second) // Attendre 1 seconde entre chaque suppression
		}
	}
	_, err = s.ChannelMessageSend(welcomeChannelID, string(welcomeMessage))
	if err != nil {
		log.Println("Erreur lors de l'envoi du message de bienvenue:", err)
	}
}
