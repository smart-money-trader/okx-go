package api

const (
	InstTypeSPOT    = "SPOT"
	InstTypeMARGIN  = "MARGIN"
	InstTypeSWAP    = "SWAP"
	InstTypeFUTURES = "FUTURES"
	InstTypeOPTION  = "OPTION"

	TdModeIsolated = "isolated"
	TdModeCross    = "cross"
	TdModeCash     = "cash"

	SideBuy  = "buy"
	SideSell = "sell"

	PosSideLong  = "long"
	PosSideShort = "short"
	PosSideNet   = "net"

	OrdTypeMarket          = "market"
	OrdTypeLimit           = "limit"
	OrdTypePostOnly        = "post_only"
	OrdTypeFok             = "fok"
	OrdTypeIoc             = "ioc"
	OrdTypeOptimalLimitIoc = "optimal_limit_ioc"
	OrdTypeConditional     = "conditional"     // 单向止盈止损
	OrdTypeOco             = "oco"             // 双向止盈止损
	OrdTypeMoveOrderStop   = "move_order_stop" // 移动止盈止损

	OrdStateCanceled        = "canceled"
	OrdStateLive            = "live"
	OrdStatePartiallyFilled = "partially_filled"
	OrdStateFilled          = "filled"

	PosModeLongShort = "long_short_mode"
	PosModeNet       = "net_mode"

	MgnModeIsolated = "isolated"
	MgnModeCross    = "cross"

	TypeAdd    = "add"
	TypeReduce = "reduce"

	// 触发价类型
	TriggerPxTypeLast  = "last"  // 最新价格
	TriggerPxTypeIndex = "index" // 指数价格
	TriggerPxTypeMark  = "mark"  // 标记价格

	OrdCategoryTwap               = "twap"
	OrdCategoryAdl                = "adl"
	OrdCategoryFullLiquidation    = "full_liquidation"
	OrdCategoryPartialLiquidation = "partial_liquidation"
	OrdCategoryDelivery           = "delivery"
	OrdCategoryDdh                = "ddh"

	StateEffective = "effective"
)
