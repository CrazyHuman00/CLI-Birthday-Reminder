package cmd

import (

	"cli-birthday-reminder/model"
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

// サブコマンドの変数群
var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Check today's birthdays",
	Run: func(cmd *cobra.Command, args []string) {
		
	},
}

var AddCmd = &cobra.Command{
	Use:   "add [name] [date]",
	Short: "Add a new birthday",
	Run: func(cmd *cobra.Command, args []string) {
		
	},
}

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all birthdays",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var RemoveCmd = &cobra.Command{
	Use:  "remove [name]",
	Short: "Remove a birthday",
	Run: func(cmd *cobra.Command, args []string) {
		
	},
}

var UpdateCmd = &cobra.Command{
	Use:  "update [name] [date]",
	Short: "Update a birthday",
	Run: func(cmd *cobra.Command, args []string) {
		
	},
}

// サブコマンドの初期化
func init() {
	rootCmd.AddCommand(execCmd)
}

// コマンドのメソッド群
// 今日の誕生日をチェックする
func CheckTodayBirthdays(db *gorm.DB) {
	today := time.Now().Format("01-02")
	var birthdays []model.Birthday

	result := db.Where("strftime('%m-%d', birthday) = ?", today).Find(&birthdays)
	if result.Error != nil {
		log.Fatalln(result.Error)
	}

	fmt.Println("Today's birthdays:")
	for _, b := range birthdays {
		fmt.Println("🎂Happy Birthday to %s!\n", b.Name)
	}
}

// 誕生日を追加する
func AddCommand(username string, birthday string) error {
	
	return nil
}

func ListCommand() {
	rootCmd.AddCommand(ListCmd)
}

func RemoveCommand() {
	rootCmd.AddCommand(RemoveCmd)
}

func UpdateCommand() {
	rootCmd.AddCommand(UpdateCmd)
}
