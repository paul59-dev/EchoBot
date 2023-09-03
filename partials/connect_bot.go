package partials

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func ConnectBot(s *discordgo.Session) {

	// Configurez un gestionnaire d'événements pour le serveur (guild)MemberAdd
	s.AddHandler(func(s *discordgo.Session, e *discordgo.GuildMemberAdd) {
		// Récupérez le membre qui a rejoint le serveur
		membre, err := s.GuildMember(e.GuildID, e.User.ID)
		if err != nil {
			log.Println("Erreur lors de la récupération du membre:", err)
			return
		}

		// Envoyez un message de bienvenue au salon de bienvenue
		s.ChannelMessageSend("ENTREES_ID", fmt.Sprintf("Bienvenue sur le serveur, %s !", membre.Mention()))
	})

	// Configurez un gestionnaire d'événements pour le serveur (guild)MemberRemove
	s.AddHandler(func(s *discordgo.Session, e *discordgo.GuildMemberRemove) {
		// Récupérez le membre qui a quitté le serveur
		membre, err := s.GuildMember(e.GuildID, e.User.ID)
		if err != nil {
			log.Println("Erreur lors de la récupération du membre:", err)
			return
		}

		// Envoyez un message d'adieu au salon de départ
		s.ChannelMessageSend("SORTIES_ID", fmt.Sprintf("Au revoir, %s !", membre.Mention()))
	})
}
