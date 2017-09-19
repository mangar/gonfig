package log

import (
	"fmt"
	"math/rand"
	"project/constants"

	emoji "gopkg.in/kyokomi/emoji.v1"
)

func LogDebug(message string) {
	newMessage := emoji.Sprint(":beetle: ", message)
	if constants.FlagDebug {
		fmt.Println(newMessage)
	}
}

func LogError(message string) {
	newMessage := emoji.Sprint(":heavy_exclamation_mark: ", message)
	fmt.Println(newMessage)
}

func LogGit(message string) {
	newMessage := emoji.Sprint(":sleeping: ", message)
	fmt.Println(newMessage)
}

func LogSuccess(message string) {
	newMessage := emoji.Sprint(message, " Enjoy! "+constants.DoneOKEmoji[rand.Intn(len(constants.DoneOKEmoji))])
	fmt.Println(newMessage)
}

func Log(message string) {
	newMessage := emoji.Sprint(":point_right: ", message)
	fmt.Println(newMessage)
}
