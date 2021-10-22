package context

type Context struct {
	Protocol ContextProtocol
	RSX      ContextRSX
	Treasury ContextTreasury
}

type ContextProtocol struct {
	RSX ContextProtocolRSX
}

type ContextProtocolRSX struct {
	Debt ContextProtocolRSXDebt
}

type ContextProtocolRSXDebt struct {
	Amount float64
	Value  float64
}

type ContextRSX struct {
	Price ContextRSXprice
}

type ContextRSXprice struct {
	Ceiling float64
	Floor   float64
}

type ContextTreasury struct {
	RSX ContextTreasuryRSX
	DAI ContextTreasuryDAI
}

// ContextTreasuryRSX is the RSX amount managed by the treasury for various
// reasons. These are e.g. minting RSX at price ceiling or redeeming RSX at
// price floor.
type ContextTreasuryRSX struct {
	Minted float64

	Supply ContextTreasuryRSXSupply
}

type ContextTreasuryRSXSupply struct {
	Circulating float64
	MarketCap   float64
	Total       float64
}

// ContextTreasuryDAI is the DAI added to the treasury by various funding
// mechanisms. These are e.g. seed investment or RSX purchases at price ceiling.
type ContextTreasuryDAI struct {
	Backing float64
	DAO     float64
	Excess  float64
	Inflow  float64
}
