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

var periodCmd = &cobra.Command{
    Use:   "period",
    Short: "Print birthdays in the next 10 days",
    RunE: func(cmd *cobra.Command, args []string) error {
        return PeriodCommand()
    },
}

func init() {
	rootCmd.AddCommand(todayCmd)
	rootCmd.AddCommand(periodCmd)
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

func PeriodCommand() error {
	// DBに接続
	dbConn := db.ConnectDB()
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.UserBirthday{})

	var users []model.UserBirthday
	period := 10
	startDate := time.Now().AddDate(0, 0, -period).Format("01/02")
	endDate := time.Now().AddDate(0, 0, period).Format("01/02")

	result := dbConn.Where("strftime('%m/%d', birthday) BETWEEN ? AND ?", startDate, endDate).Find(&users)
	if result.Error != nil {
		return result.Error
	}

	if len(users) == 0 {
        fmt.Println("No birthdays in the next 10 days!")
        return nil
    }

	fmt.Println("Upcoming birthdays:")
    for _, user := range users {
        fmt.Printf("%12s %s\n", user.Name, user.Birthday.Format("01/02"))
    }
	return nil
}