package mcuser

import(
	"encoding/json"
    "fmt"
    "io/ioutil"
	"net/http"
	"github.com/satori/go.uuid"
	"os"
	"reflect"
)

type MCUser struct {
	UUID     string `json:"id"`
	Username string `json:"name"`
}

type MCWhiteListUser struct {
	UUID	 string `json:"uuid"`
	Name	 string `json:"name"`
}

func NewUser(username string) *MCUser {
	fmt.Println("Creating new MCUser")
	return &MCUser{"",username}
}

func (user *MCUser) SetUUID(uuid string) {
	user.UUID = uuid
}

func (user *MCUser) GetUUID() string {
	return user.UUID
}

func (user *MCUser) GetUsername() string {
	return user.Username
}

func (user *MCUser) GetUser() *MCUser {
	return user
}

func (user *MCUser) GetMinecraftUser() *MCUser {
	mojangEndpoint := "https://api.mojang.com/users/profiles/minecraft/" + user.GetUsername()
	fmt.Println("Making API call to:", mojangEndpoint)
	request, _ := http.NewRequest("GET", mojangEndpoint, nil)
	request.Header.Set("User-Agent", "Xxplodibot")
	client := &http.Client{}
	response, err := client.Do(request)
    if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
		data, _ := ioutil.ReadAll(response.Body)
		er := json.Unmarshal(data, &user)
		if er != nil {
			fmt.Printf("Failed to convert response to type User %s\n", er)
		} else {
			uuid, _ := uuid.FromString(user.UUID)
			user.SetUUID(uuid.String())
		}
	}
	response.Body.Close()

	return user
}

func UpdateWhitelistFile(user *MCUser) bool {
	// jsonFile, err := os.Open(os.Getenv("HOME") + "/go/src/minecraft/whitelist.json")
	jsonFile, err := os.Open(os.Getenv("HOME") + "/FTBRevelationServer_2.2.0/whitelist.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened whitelist.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var whitelisted []MCWhiteListUser

	json.Unmarshal(byteValue, &whitelisted)

	whitelistUser := MCWhiteListUser{user.GetUUID(), user.GetUsername()}
	isWhitelisted := false

	for _, u := range whitelisted {
		if(u.Name == user.Username) {
			isWhitelisted = true
			break
		}
	}

	if(isWhitelisted) {
		fmt.Println("MCUser is whitelisted:",isWhitelisted)
		return true
	}

	fmt.Println("MCUser is not whitelisted adding to whitelist.")

	_, err = json.Marshal(whitelistUser)
	if err != nil {
		fmt.Println(err)
		return false
	}

	whitelisted = append(whitelisted, whitelistUser)

	result, _ := json.Marshal(whitelisted)

	// err = ioutil.WriteFile(os.Getenv("HOME") + "/go/src/minecraft/whitelist.json", result, 0644)
	err = ioutil.WriteFile(os.Getenv("HOME") + "/FTBRevelationServer_2.2.0/whitelist.json", result, 0644)

	return true
}

func in_array(val interface{}, array interface{}) (exists bool, index int) {
    exists = false
    index = -1

    switch reflect.TypeOf(array).Kind() {
    case reflect.Slice:
        s := reflect.ValueOf(array)

        for i := 0; i < s.Len(); i++ {
            if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
                index = i
                exists = true
                return
            }
        }
    }

    return
}