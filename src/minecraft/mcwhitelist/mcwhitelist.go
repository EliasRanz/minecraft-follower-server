package mcwhitelist

import (
    "fmt"
	"os"
	"os/exec"
	"strings"
	"minecraft/mcuser"
)

type BotQuery struct {
	Username string `json:"username"`
}

func NewMCWhitelist(username string) bool {
	fmt.Println("Initiating whitelist for", username)
	user := mcuser.NewUser(username)
	user = user.GetMinecraftUser()
	fmt.Println("MCUser:", user.GetUser())
	fmt.Println("MCUser.UUID from MCUser:", user.GetUUID())
	fmt.Println("MCUser.Username from MCUser:", user.GetUsername())

	screenName := GetScreen()
	fmt.Println("screenName:", screenName)

	ReloadWhitelist(screenName)

	return mcuser.UpdateWhitelistFile(user)
}

func GetScreen() string {
	// cmd := "ls -R /var/folders/lm/wj0vbjds0mq8g4cxjry1_b_wflk43d/T/.screen | grep '\\.Server'"
	cmd := "ls -R /var/run/screen/S-root/ | grep '\\.Server'"
	fmt.Println("Running command:", cmd)
	cmdOut, cmdErr := exec.Command("bash", "-c", cmd).Output()
	if cmdErr != nil {
		fmt.Println("Failled to run command.")
		fmt.Fprintln(os.Stderr, "There was an error running screen command: ", cmdErr)
	}
	fmt.Println(string(cmdOut))

	return string(cmdOut)
}

func ReloadWhitelist(screenName string) {
	// cmd := "screen -S 1391.FTBRevelationServer_2.2.0 -p 0 -X stuff \"whitelist reload^M\""
	cmd := "screen -S {screenName} -p 0 -X stuff \"whitelist reload^M\""
	cmd = strings.Replace(cmd, "{screenName}", screenName, -1)
	cmd = strings.Replace(cmd, "\n", " ", -1)
	fmt.Println(cmd)

	cmdOut, cmdErr := exec.Command("bash", "-c", cmd).Output()
	if cmdErr != nil {
		fmt.Println("Failled to run command.")
		fmt.Fprintln(os.Stderr, "There was an error running screen command: ", cmdErr)
		fmt.Println(string(cmdOut))
	}

	fmt.Println("Successfully ran command. Got output\n", string(cmdOut))
}