package cmd

import (
	"fmt"
	"os"

	crafter "github.com/kn-lim/xivcrafter/pkg/crafter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Prints and validates XIVCrafter's configuration",
	Long:  `Prints and validates XIVCrafter's configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		food := viper.GetString("food")
		s := fmt.Sprintf("food: %s", food)
		fmt.Println(s)

		foodDuration := viper.GetInt("foodDuration")
		s = fmt.Sprintf("foodDuration: %d", foodDuration)
		fmt.Println(s)

		potion := viper.GetString("potion")
		s = fmt.Sprintf("potion: %s", potion)
		fmt.Println(s)

		amount := viper.GetInt("amount")
		s = fmt.Sprintf("amount: %d", amount)
		fmt.Println(s)

		macro1 := viper.GetString("macro1")
		s = fmt.Sprintf("macro1: %s", macro1)
		fmt.Println(s)

		macro1Duration := viper.GetInt("macro1Duration")
		s = fmt.Sprintf("macro1Duration: %d", macro1Duration)
		fmt.Println(s)

		macro2 := viper.GetString("macro2")
		s = fmt.Sprintf("macro2: %s", macro2)
		fmt.Println(s)

		macro2Duration := viper.GetInt("macro2Duration")
		s = fmt.Sprintf("macro2Duration: %d", macro2Duration)
		fmt.Println(s)

		confirm := viper.GetString("confirm")
		s = fmt.Sprintf("confirm: %s", confirm)
		fmt.Println(s)

		cancel := viper.GetString("cancel")
		s = fmt.Sprintf("cancel: %s", cancel)
		fmt.Println(s)

		startPause := viper.GetString("startPause")
		s = fmt.Sprintf("startPause: %s", startPause)
		fmt.Println(s)

		stop := viper.GetString("stop")
		s = fmt.Sprintf("stop: %s", stop)
		fmt.Println(s)

		// Check if all keys and flags are valid
		fmt.Println()
		xiv := new(crafter.XIVCrafter)
		xiv.Init(food, foodDuration, potion, amount, macro1, macro1Duration, macro2, macro2Duration, confirm, cancel, startPause, stop)

		if !(crafter.CheckKeys(*xiv) && crafter.CheckFlags(*xiv)) {
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
