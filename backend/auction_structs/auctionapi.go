package auction_structs

type Auction struct {
	Auc uint64
	Item uint64
	Owner string
	OwnerRealm string
	Bid uint64
	Buyout uint64
	Quantity uint64
	TimeLeft string //TODO: Create ENUM
	Rand int64
	Seed int64
}
type Realm struct {
	Name string
	Slug string
}
type AuctionsJSON struct {
	Realms []Realm
	Auctions []Auction
}
type BaseAuctionResponse struct {
	Files []AuctionResponseFiles
}
type AuctionResponseFiles struct {
	Url string
	LastModified uint64
}