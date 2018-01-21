package auctions

import (
	"github.com/go-resty/resty"
	"log"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"sort"
	"github.com/thilltbc/ahrank/backend/auction_structs"
)
type PlayerScore struct {
	Owner string `json:"owner"`
	Count int `json:"count"`
}
//var conn *pgx.Conn
var GlobalAuctions []auction_structs.Auction
func INIT() {
	//auctionRecordsURL := getAHRecordsFileURL("azuremyst")
	//importAuctionRecords(auctionRecordsURL)
/*	var config pgx.ConnConfig
	config.Database = "ahrank"
	config.User = "ahrank"
	config.Password = "lolpass"
	var err error
	conn, err = pgx.Connect(config)
	if err != nil {
		log.Fatalf("Error connecting to db: %v\n", err)
	}
*/
	GlobalAuctions = getDownloadedAuctions().Auctions

}
func GetAuctionCountRanking() []PlayerScore {
	// attempts imitates the query below
	// SELECT COUNT(auc), owner FROM auctions GROUP BY owner;
	var score = make(map[string]int)
	for _, auc := range GlobalAuctions {

		score[auc.Owner]++

	}
	var sortedScores []PlayerScore
	for k, v := range score {
		sortedScores = append(sortedScores, PlayerScore{k, v})
	}
	sort.Slice(sortedScores, func(i, j int) bool {
		return sortedScores[i].Count > sortedScores[j].Count
	})
	for i, player := range sortedScores {
		if i > 50 {
			break
		}
		fmt.Printf("%v: %v\n", player.Owner, player.Count)

	}
	return sortedScores
}
/*
avoiding db for now
func insertAuctions(auctions auction_structs.AuctionsJSON) {
	trans,_ := conn.Begin()
	defer trans.Rollback()
	start := time.Now()
	for index, auc:= range auctions.Auctions{
		_, err := trans.Exec(`INSERT INTO auctions(auc,item,owner,ownerRealm,bid,buyout,quantity,timeLeft,seed,rand) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`,
			auc.Auc, auc.Item, auc.Owner, auc.OwnerRealm, auc.Bid, auc.Buyout, auc.Quantity, auc.TimeLeft, auc.Seed, auc.Rand)
		if err != nil {
			log.Fatalf("Error exec: %v", err)
		}
		if index % 5000 == 0  && index > 1{
			transErr := trans.Commit()
			trans,_ = conn.Begin()
			if transErr != nil {
				log.Fatalf("Error commiting 100 auction transaction to PG: %v\n", transErr)
			}
		}
	}
	transErr := trans.Commit()
	t := time.Now();
	elapsed := t.Sub(start)
	fmt.Printf("Took %v secs to insert all the rows.\n", elapsed)
	if transErr != nil {
		log.Fatalf("Error commiting last < 100 auction transaction to PG: %v\n", transErr)
	}

}*/
func getDownloadedAuctions() auction_structs.AuctionsJSON{
	// for testing to avoid exhausting api/bandwidth
	auctionsJSON, err := ioutil.ReadFile("./auctions.json")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	var respStruct auction_structs.AuctionsJSON
	err  = json.Unmarshal(auctionsJSON, &respStruct)
	if err != nil {
		log.Fatalf("Error unmarshalling Auction JSON File: %v\n", err)
	}
	fmt.Printf("Found %v auctions!\n", len(respStruct.Auctions))
	return respStruct
}
func getLiveAuctions(auctionRecordsURL string) auction_structs.AuctionsJSON {
	// gets auctionfile and returns them
	resp, err := resty.R().Get(auctionRecordsURL)
	if err != nil {
		log.Fatalf("Error getting Auction JSON Data File: %v\n", err)
	}
	var respStruct auction_structs.AuctionsJSON
	err  = json.Unmarshal(resp.Body(), &respStruct)
	if err != nil {
		log.Fatalf("Error unmarshalling Auction JSON File: %v\n", err)
	}
	fmt.Printf("Found %v auctions!\n", len(respStruct.Auctions))
	return respStruct
}

func getAHRecordsFileURL(realmSlug string) string {
	resp, err := resty.R().Get("https://us.api.battle.net/wow/auction/data/" + realmSlug + "?locale=en_US&apikey=rcnpk45pfgjxge4p96udb9cstck2wrra")
	if err != nil {
		log.Fatalf("Error getting Auction JSON Files: %v\n", err)
	}
	var respStruct auction_structs.BaseAuctionResponse
	err  = json.Unmarshal(resp.Body(), &respStruct)
	if err != nil {
		log.Fatalf("Error unmarshalling Initial Auction JSON response: %v\n", err)
	}
	fmt.Printf("%v", respStruct)
	return respStruct.Files[0].Url
}