package cmd

import (
	"birthday-reminder/db"
	"birthday-reminder/model"
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
)

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
	RunE: func(cmd *cobra.Command, args []string) error {
		return AddCommand(args[0], args[1])
	},
}

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all birthdays",
	RunE: func(cmd *cobra.Command, args []string) error {
		return ListCommand()
	},
}

var RemoveCmd = &cobra.Command{
	Use:  "remove [name]",
	Short: "Remove a birthday",
	RunE: func(cmd *cobra.Command, args []string) error {
		return RemoveCommand()
	},
}

var UpdateCmd = &cobra.Command{
	Use:  "update [name] [date]",
	Short: "Update a birthday",
	Run: func(cmd *cobra.Command, args []string) {
		
	},
}

func init() {
	rootCmd.AddCommand(execCmd)
	rootCmd.AddCommand(AddCmd)
	rootCmd.AddCommand(ListCmd)
	rootCmd.AddCommand(ListCmd)
	rootCmd.AddCommand(RemoveCmd)
	rootCmd.AddCommand(UpdateCmd)
}

/** コマンドのメソッド群 **/

// 誕生日を追加する
func AddCommand(username string, birthday string) error {
	// DBに接続
	dbConn := db.ConnectDB()
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.UserBirthday{})

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
func ListCommand() error {
	// DBに接続
	dbConn := db.ConnectDB()
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.UserBirthday{})

	var users []model.UserBirthday
	result := dbConn.Find(&users)
	if result.Error != nil {
		return result.Error
	}

	// 誕生日が見つからない場合
	if len(users) == 0 {
		fmt.Println("No birthdays found")
		return nil
	}

	// 誕生日が見つかった場合
	for _, user := range users {
		fmt.Printf("%s's birthday is on %s\n", user.Name, user.Birthday.Format("01/02"))
	}
	return nil
}


// 誕生日を削除する
func RemoveCommand() error {
	// DBに接続
	dbConn := db.ConnectDB()
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.UserBirthday{})

	return nil
}

// 誕生日を更新する
func UpdateCommand() error {
	// DBに接続
	dbConn := db.ConnectDB()
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.UserBirthday{})

	return nil
}
