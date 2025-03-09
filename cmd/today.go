package cmd

import (
	"github.com/CrazyHuman00/CLI-Birthday-Reminder/db"
	"github.com/CrazyHuman00/CLI-Birthday-Reminder/model"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "Print today's birthday",
	RunE: func(cmd *cobra.Command, args []string) error {
		return TodayCommand()
	},
}

func init() {
	rootCmd.AddCommand(todayCmd)
}

// 今日の誕生日を表示する
func TodayCommand() error{
	// DBに接続
	dbConn := db.ConnectDB()
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.UserBirthday{})

	var users []model.UserBirthday
	today := time.Now().Format("01/02")

	result := dbConn.Where("strftime('%m/%d', birthday) = ?", today).Find(&users)
	if result.Error != nil {
		return result.Error
	}

	if len(users) == 0 {
		fmt.Println("No birthdays today!")
		return nil
	}

	fmt.Println("Today's birthdays:")
	for _, user := range users {
		fmt.Printf("%12s %s\n", user.Name, user.Birthday.Format("01/02"))
	}
	return nil
}