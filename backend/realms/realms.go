package realms

import (
	"github.com/go-resty/resty"
	"github.com/thilltbc/ahrank/backend/auction_structs"
	"encoding/json"
	"fmt"
	"log"
)

func GetRealmNames() []string {
	apiUrl := "https://us.api.battle.net/wow/realm/status?locale=en_US&apikey=rcnpk45pfgjxge4p96udb9cstck2wrra"
	resp, err := resty.R().Get(apiUrl)
	if err != nil {
		log.Fatalf("Error getting Realm List Response: %v\n", err)
	}
	var respStruct auction_structs.RealmsListResponse
	err  = json.Unmarshal(resp.Body(), &respStruct)
	if err != nil {
		log.Fatalf("Error unmarshalling Realm List Response: %v\n", err)
	}
	fmt.Printf("Found %v realms!\n", len(respStruct.Realms))
	var realmNames = make([]string, len(respStruct.Realms))
	for _,realm := range respStruct.Realms {
		if realm.Name != "" && realm.Status == true {
			realmNames = append(realmNames, realm.Name)
		}
	}
	return realmNames
}