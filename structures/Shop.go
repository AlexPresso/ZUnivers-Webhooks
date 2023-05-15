package structures

import "gorm.io/gorm"

type ShopEntry struct {
	gorm.Model

	ShopEntryID   string         `json:"id"`
	BeginDate     *DateTime      `json:"beginDate"`
	InitialStock  int            `json:"initialStock"`
	ShopInventory *ShopInventory `json:"shopInventory" gorm:"-"`
}

type ShopInventory struct {
	Balance      int   `json:"balance"`
	Golden       bool  `json:"isGolden"`
	UpgradeLevel int   `json:"upgradeLevel"`
	Item         *Item `json:"item"`
}
