package cli

import (
	"time"

	"github.com/bwoff11/frens/internal/db"
	"github.com/bwoff11/frens/internal/models"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(userCmd)

	userCmd.AddCommand(addCmd)

	userCmd.AddCommand(delCmd)
}

var userCmd = &cobra.Command{
	Use:   "users",
	Short: "Manage users in the database",
	Long:  `Directly modify users in the database.`,
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a user",
	Long:  `Add a user.`,
	Run: func(cmd *cobra.Command, args []string) {
		db.Connect()

		newAccount := models.Account{
			Username:       "username2",
			Acct:           "accountname2",
			URL:            "https://example.com",
			DisplayName:    "Display Name",
			Note:           "Note",
			Avatar:         "https://example.com/avatar.png",
			AvatarStatic:   "https://example.com/avatar.png",
			Header:         "https://example.com/header.png",
			HeaderStatic:   "https://example.com/header.png",
			Locked:         false,
			Discoverable:   true,
			CreatedAt:      time.Now(),
			LastStatusAt:   time.Now(),
			StatusCount:    0,
			FollowersCount: 0,
			FollowingCount: 0,
			// Moved:          false,
			// Fields:         []models.Field{},
			Bot: false,
			// Source:        models.Source{},
			Suspended:     false,
			MuteExpiresAt: time.Now(),
		}
		db.AddAccount(&newAccount)
	},
}

var delCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a user",
	Long:  `Delete a user.`,
	Run: func(cmd *cobra.Command, args []string) {
		db.Connect()

		accoundToDelete := models.Account{
			ID: 12,
		}

		db.DeleteAccount(&accoundToDelete)
	},
}
