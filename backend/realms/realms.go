package realms

import (
	"github.com/go-resty/resty"
	"encoding/json"
	"fmt"
	"log"
)
var RealmList []FullRealmInfo
func GetRealmNames(realmList []FullRealmInfo) []string {
	var realmNames = make([]string, len(realmList))
	for i,realm := range realmList {
		realmNames[i] = realm.Name
	}
	return realmNames
}
func GetRealmList() []FullRealmInfo {
	if RealmList != nil {
		// TODO: CACHING?
		return RealmList
	}
	apiUrl := "https://us.api.battle.net/wow/realm/status?locale=en_US&apikey=rcnpk45pfgjxge4p96udb9cstck2wrra"
	resp, err := resty.R().Get(apiUrl)
	if err != nil {
		log.Fatalf("Error getting Realm List Response: %v\n", err)
	}
	var respStruct RealmListResponse
	err  = json.Unmarshal(resp.Body(), &respStruct)
	if err != nil {
		log.Fatalf("Error unmarshalling Realm List Response: %v\n", err)
	}
	fmt.Printf("Found %v realms!\n", len(respStruct.Realms))
	var realmNames = make([]FullRealmInfo, 0)
	for _,realm := range respStruct.Realms {
		if realm.Name != "" && realm.Status == true {
			realmNames = append(realmNames, realm)
		}
	}

	if RealmList == nil {
		RealmList = respStruct.Realms
	}
	return realmNames
}
type RealmListResponse struct {
	Realms []FullRealmInfo
}

type FullRealmInfo struct {
	Type string
	Population string
	Queue bool
	Status bool
	Name string
	Slug string
	Battlegroup string
	Locale string
	Timezone string
	Connected_realms []string
}