package context

type Context struct {
	Pool     ContextPool
	Protocol ContextProtocol
	RSX      ContextRSX
	Treasury ContextTreasury
}

type ContextPool struct {
	RSXDAI ContextPoolRSXDAI
	RSXOHM ContextPoolRSXOHM
}

type ContextPoolRSXDAI struct {
	ConstantK float64
	Liquidity float64

	RSX ContextPoolRSXDAIRSX
	DAI ContextPoolRSXDAIDAI
}

type ContextPoolRSXDAIRSX struct {
	Amount float64
	Price  float64
	Value  float64
}

type ContextPoolRSXDAIDAI struct {
	Amount float64
	Price  float64
	Value  float64
}

type ContextPoolRSXOHM struct {
	ConstantK float64
	Liquidity float64

	RSX ContextPoolRSXOHMRSX
	OHM ContextPoolRSXOHMOHM
}

type ContextPoolRSXOHMRSX struct {
	Amount float64
	Price  float64
	Value  float64
}

type ContextPoolRSXOHMOHM struct {
	Amount float64
	Price  float64
	Value  float64
}

type ContextProtocol struct {
	Debt ContextProtocolDebt
}

type ContextProtocolDebt struct {
	RSX ContextProtocolDebtRSX
}

type ContextProtocolDebtRSX struct {
	Amount float64
	Price  float64
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
	ExcessReserves float64
	VolumeInflow   float64

	RSX ContextTreasuryRSX
	DAI ContextTreasuryDAI
}

// ContextTreasuryRSX is the RSX amount managed by the treasury for various
// reasons. These are e.g. minting RSX at price ceiling or redeeming RSX at
// price floor.
type ContextTreasuryRSX struct {
	Amount float64
}

// ContextTreasuryDAI is the DAI added to the treasury by various funding
// mechanisms. These are e.g. seed investment or RSX purchases at price ceiling.
type ContextTreasuryDAI struct {
	Amount float64
}
