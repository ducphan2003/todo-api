package common

var RewardPointType = struct {
	ORDER_REWARD     string
	USE_FOR_ORDER    string
	ORDER_CANCEL     string
	PROMOTION_REWARD string
}{
	ORDER_REWARD:     "order_reward",
	USE_FOR_ORDER:    "use_for_order",
	ORDER_CANCEL:     "order_cancel",
	PROMOTION_REWARD: "promotion_reward",
}

var PublishedScope = struct {
	POS    string
	GLOBAL string
}{
	POS:    "pos",
	GLOBAL: "global",
}
