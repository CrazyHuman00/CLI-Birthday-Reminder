package cmd

import (
	"birthday-reminder/db"
	"birthday-reminder/model"
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

/** サブコマンドの変数群 **/
var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Check today's birthdays",
	Run: func(cmd *cobra.Command, args []string) {
		
	},
}

var AddCmd = &cobra.Command{
	Use:   "add [name] [date]",
	Short: "Add a new birthday",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error{
		return AddCommand(args[0], args[1])
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

func init() {
	dbConn := db.ConnectDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.UserBirthday{})

	rootCmd.AddCommand(execCmd)
	rootCmd.AddCommand(AddCmd)
	rootCmd.AddCommand(ListCmd)
	rootCmd.AddCommand(RemoveCmd)
	rootCmd.AddCommand(UpdateCmd)
}

/** コマンドのメソッド群 **/
func CheckTodayBirthdays(db *gorm.DB) {
	today := time.Now().Format("01-02")
	var birthdays []model.UserBirthday

	result := db.Where("strftime('%m-%d', birthday) = ?", today).Find(&birthdays)
	if result.Error != nil {
		log.Fatalln(result.Error)
	}

	fmt.Printf("Today's birthdays:\n")
	for _, b := range birthdays {
		fmt.Printf("🎂Happy Birthday to %s!\n", b.Name)
	}
}

// 誕生日を追加する
func AddCommand(username string, birthday string) error {
	birthdayTime, err := time.Parse("01/02", birthday)
	if err != nil {
		log.Fatalln(err)
	}

	result := dbConn.Create(&model.UserBirthday{Name: username, Birthday: birthdayTime})
	if result.Error != nil {
		return result.Error
	}
	fmt.Printf("Birthday added successfully! %s's birthday is on %s\n", username, birthday)
	return nil
}

// 誕生日をリストする
func ListCommand() {
	rootCmd.AddCommand(ListCmd)
}

// 誕生日を削除する
func RemoveCommand() {
	rootCmd.AddCommand(RemoveCmd)
}

// 誕生日を更新する
func UpdateCommand() {
	rootCmd.AddCommand(UpdateCmd)
}
