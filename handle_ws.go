package mysunxapi

func handleWsData[T any](data []byte) (*T, error) {
	var res T
	err := json.Unmarshal(data, &res)
	if err != nil {
		log.Error("data: ", string(data))
		log.Error("err: ", err.Error())
		return nil, err
	}
	return &res, nil
}

type WsMarketCh struct {
	Ch string `json:"ch"`
}

// Public
type WsBBOTick struct {
	Mrid    int64     `json:"mrid"`
	Id      int64     `json:"id"`
	Bid     []float64 `json:"bid"`
	Ask     []float64 `json:"ask"`
	Ts      int64     `json:"ts"`
	Version int64     `json:"version"`
	Ch      string    `json:"ch"`
}

type WsBBORes struct {
	Ch   string    `json:"ch"`
	Ts   int64     `json:"ts"`
	Tick WsBBOTick `json:"tick"`
}

type WsBBO struct {
	Ch   string `json:"ch"`
	Ts   int64  `json:"ts"`
	Tick struct {
		Mrid    int64      `json:"mrid"`
		Id      int64      `json:"id"`
		Bid     PriceLevel `json:"bid"`
		Ask     PriceLevel `json:"ask"`
		Ts      int64      `json:"ts"`
		Version int64      `json:"version"`
		Ch      string     `json:"ch"`
	} `json:"tick"`
}

func (b *WsBBORes) convertToWsBBO() *WsBBO {
	return &WsBBO{
		Ch: b.Ch,
		Ts: b.Ts,
		Tick: struct {
			Mrid    int64      `json:"mrid"`
			Id      int64      `json:"id"`
			Bid     PriceLevel `json:"bid"`
			Ask     PriceLevel `json:"ask"`
			Ts      int64      `json:"ts"`
			Version int64      `json:"version"`
			Ch      string     `json:"ch"`
		}{
			Mrid: b.Tick.Mrid,
			Id:   b.Tick.Id,
			Bid:  PriceLevel{Price: b.Tick.Bid[0], Volume: b.Tick.Bid[1]},
			Ask:  PriceLevel{Price: b.Tick.Ask[0], Volume: b.Tick.Ask[1]},

			Ts:      b.Tick.Ts,
			Version: b.Tick.Version,
			Ch:      b.Ch,
		},
	}
}

// Private
type WsPrivateResCommon struct {
	Op    string `json:"op"`
	Topic string `json:"topic"`
	Ts    int64  `json:"ts"`
	Uid   string `json:"uid"`
	Event string `json:"event"`
}

type WsAccountRes struct {
	Op           string `json:"op"`
	Topic        string `json:"topic"`
	Ts           int64  `json:"ts"`
	Uid          string `json:"uid"`
	Event        string `json:"event"`
	ContractCode string `json:"contract_code"`
	Data         struct {
		Equity  string `json:"equity"`
		State   string `json:"state"`
		Details []struct {
			Currency              string `json:"currency"`
			Equity                string `json:"equity"`
			Available             string `json:"available"`
			ProfitUnreal          string `json:"profit_unreal"`
			InitialMargin         string `json:"initial_margin"`
			MaintenanceMargin     string `json:"maintenance_margin"`
			MaintenanceMarginRate string `json:"maintenance_margin_rate"`
			InitialMarginRate     string `json:"initial_margin_rate"`
			Voucher               string `json:"voucher"`
			VoucherValue          string `json:"voucher_value"`
			WithdrawAvailable     string `json:"withdraw_available"`
			CreatedTime           string `json:"created_time"`
			UpdatedTime           string `json:"updated_time"`
		} `json:"details,omitempty"`
		InitialMargin         string `json:"initial_margin"`
		MaintenanceMargin     string `json:"maintenance_margin"`
		MaintenanceMarginRate string `json:"maintenance_margin_rate"`
		ProfitUnreal          string `json:"profit_unreal"`
		AvailableMargin       string `json:"available_margin"`
		VoucherValue          string `json:"voucher_value"`
		CreatedTime           string `json:"created_time"`
		UpdatedTime           string `json:"updated_time"`
		Version               int64  `json:"version"`
	} `json:"data"`
}

type PriceLevel struct {
	Price  float64 `json:"price"`
	Volume float64 `json:"volume"`
}

type WsDepthRes struct {
	Ch   string `json:"ch"`
	Ts   int64  `json:"ts"`
	Tick struct {
		Mrid    int64       `json:"mrid"`
		Id      int64       `json:"id"`
		Bids    [][]float64 `json:"bids"`
		Asks    [][]float64 `json:"asks"`
		Ts      int64       `json:"ts"`
		Version int64       `json:"version"`
		Ch      string      `json:"ch"`
	} `json:"tick"`
}

type WsDepth struct {
	Ch   string `json:"ch"`
	Ts   int64  `json:"ts"`
	Tick struct {
		Mrid    int64        `json:"mrid"`
		Id      int64        `json:"id"`
		Bids    []PriceLevel `json:"bids"`
		Asks    []PriceLevel `json:"asks"`
		Ts      int64        `json:"ts"`
		Version int64        `json:"version"`
		Ch      string       `json:"ch"`
	} `json:"tick"`
}

func (d *WsDepthRes) convertToWsDepthRes() *WsDepth {
	var asks, bids []PriceLevel
	for _, bid := range d.Tick.Bids {
		bids = append(bids, PriceLevel{
			Price:  bid[0],
			Volume: bid[1],
		})
	}
	for _, ask := range d.Tick.Asks {
		asks = append(asks, PriceLevel{
			Price:  ask[0],
			Volume: ask[1],
		})
	}
	return &WsDepth{
		Ch: d.Ch,
		Ts: d.Ts,
		Tick: struct {
			Mrid    int64        `json:"mrid"`
			Id      int64        `json:"id"`
			Bids    []PriceLevel `json:"bids"`
			Asks    []PriceLevel `json:"asks"`
			Ts      int64        `json:"ts"`
			Version int64        `json:"version"`
			Ch      string       `json:"ch"`
		}{
			Mrid:    d.Tick.Mrid,
			Id:      d.Tick.Id,
			Bids:    bids,
			Asks:    asks,
			Ts:      d.Tick.Ts,
			Version: d.Tick.Version,
			Ch:      d.Ch,
		},
	}
}

type WsDepthHighFreqRes struct {
	Ch   string `json:"ch"`
	Ts   int64  `json:"ts"`
	Tick struct {
		Asks    [][]float64 `json:"asks"`
		Bids    [][]float64 `json:"bids"`
		Ch      string      `json:"ch"`
		Event   string      `json:"event"`
		Id      int64       `json:"id"`
		Mrid    int64       `json:"mrid"`
		Ts      int64       `json:"ts"`
		Version int64       `json:"version"`
	} `json:"tick"`
}

type WsDepthHighFreq struct {
	Ch   string `json:"ch"`
	Ts   int64  `json:"ts"`
	Tick struct {
		Asks    []PriceLevel `json:"asks"`
		Bids    []PriceLevel `json:"bids"`
		Ch      string       `json:"ch"`
		Event   string       `json:"event"`
		Id      int64        `json:"id"`
		Mrid    int64        `json:"mrid"`
		Ts      int64        `json:"ts"`
		Version int64        `json:"version"`
	} `json:"tick"`
}

func (d *WsDepthHighFreqRes) convertToWsDepthHighFreq() *WsDepthHighFreq {
	var asks, bids []PriceLevel
	for _, ask := range d.Tick.Asks {
		asks = append(asks, PriceLevel{
			Price:  ask[0],
			Volume: ask[1],
		})
	}
	for _, bid := range d.Tick.Bids {
		bids = append(bids, PriceLevel{
			Price:  bid[0],
			Volume: bid[1],
		})
	}
	return &WsDepthHighFreq{
		Ch: d.Ch,
		Ts: d.Ts,
		Tick: struct {
			Asks    []PriceLevel `json:"asks"`
			Bids    []PriceLevel `json:"bids"`
			Ch      string       `json:"ch"`
			Event   string       `json:"event"`
			Id      int64        `json:"id"`
			Mrid    int64        `json:"mrid"`
			Ts      int64        `json:"ts"`
			Version int64        `json:"version"`
		}{
			Asks:    asks,
			Bids:    bids,
			Ch:      d.Tick.Ch,
			Event:   d.Tick.Event,
			Id:      d.Tick.Id,
			Mrid:    d.Tick.Mrid,
			Ts:      d.Tick.Ts,
			Version: d.Tick.Version,
		},
	}
}

type WsKline struct {
	Ch   string `json:"ch"`
	Ts   int64  `json:"ts"`
	Tick struct {
		Id            int64   `json:"id"`
		Mrid          int64   `json:"mrid"`
		Open          float64 `json:"open"`
		Close         float64 `json:"close"`
		High          float64 `json:"high"`
		Low           float64 `json:"low"`
		Amount        float64 `json:"amount"`
		Vol           float64 `json:"vol"`
		TradeTurnover float64 `json:"trade_turnover"`
		Count         int64   `json:"count"`
	} `json:"tick"`
}

type WsTradeDetail struct {
	Ch   string `json:"ch"`
	Ts   int64  `json:"ts"`
	Tick struct {
		Id   int64 `json:"id"`
		Ts   int64 `json:"ts"`
		Data []struct {
			Amount        float64 `json:"amount"`
			Ts            int64   `json:"ts"`
			Id            int64   `json:"id"`
			Price         float64 `json:"price"`
			Direction     string  `json:"direction"`
			Quantity      float64 `json:"quantity"`
			TradeTurnover float64 `json:"trade_turnover"`
		} `json:"data"`
	} `json:"tick"`
}

type WsPositions struct {
	Op           string `json:"op"`
	Topic        string `json:"topic"`
	ContractCode string `json:"contract_code"`
	Ts           int64  `json:"ts"`
	Uid          string `json:"uid"`
	Event        string `json:"event"`
	Data         []struct {
		ContractCode      string `json:"contract_code"`
		Symbol            string `json:"symbol"`
		PositionMode      string `json:"position_mode"`
		PositionSide      string `json:"position_side"`
		Direction         string `json:"direction"`
		MarginMode        string `json:"margin_mode"`
		OpenAvgPrice      string `json:"open_avg_price"`
		Volume            string `json:"volume"`
		Available         string `json:"available"`
		Fee               string `json:"fee"`
		LeverRate         int64  `json:"lever_rate"`
		AdlRiskPercent    int    `json:"adl_risk_percent"`
		LiquidationPrice  string `json:"liquidation_price"`
		InitialMargin     string `json:"initial_margin"`
		MaintenanceMargin string `json:"maintenance_margin"`
		ProfitUnreal      string `json:"profit_unreal"`
		Profit            string `json:"profit"`
		ProfitRate        string `json:"profit_rate"`
		MarginRate        string `json:"margin_rate"`
		State             string `json:"state"`
		FundingFee        string `json:"funding_fee"`
		MarkPrice         string `json:"mark_price"`
		ContractType      string `json:"contract_type"`
		Version           int64  `json:"version"`
		CreatedTime       string `json:"created_time"`
		UpdatedTime       string `json:"updated_time"`
	} `json:"data"`
}

type WsNotificationTopic struct {
	Topic string `json:"topic"`
}
type WsMatchOrders struct {
	Op           string `json:"op"`
	Topic        string `json:"topic"`
	ContractCode string `json:"contract_code"`
	Ts           int64  `json:"ts"`
	Uid          string `json:"uid"`
	Data         []struct {
		Side             string `json:"side"`
		Type             string `json:"type"`
		Price            string `json:"price"`
		Volume           string `json:"volume"`
		State            string `json:"state"`
		Id               string `json:"id"`
		ContractCode     string `json:"contract_code"`
		ContractType     string `json:"contract_type"`
		OrderId          string `json:"order_id"`
		PositionSide     string `json:"position_side"`
		PriceMatch       string `json:"price_match"`
		ClientOrderId    string `json:"client_order_id"`
		MarginMode       string `json:"margin_mode"`
		LeverRate        string `json:"lever_rate"`
		OrderSource      string `json:"order_source"`
		ReduceOnly       bool   `json:"reduce_only"`
		TimeInForce      string `json:"time_in_force"`
		CancelReason     string `json:"cancel_reason"`
		TradeId          string `json:"trade_id"`
		TradeVolume      string `json:"trade_volume"`
		TotalTradeVolume string `json:"total_trade_volume"`
		TradePrice       string `json:"trade_price"`
		TradeTurnover    string `json:"trade_turnover"`
		Role             string `json:"role"`
		CreatedTime      string `json:"created_time"`
		MatchTime        string `json:"match_time"`
		SelfMatchPrevent string `json:"self_match_prevent"`
	} `json:"data"`
}

type WsTrade struct {
	Op           string `json:"op"`
	Topic        string `json:"topic"`
	ContractCode string `json:"contract_code"`
	Ts           int64  `json:"ts"`
	Uid          string `json:"uid"`
	Data         []struct {
		Direction     string `json:"direction"`
		PositionSide  string `json:"position_side"`
		Id            string `json:"id"`
		ContractCode  string `json:"contract_code"`
		ContractType  string `json:"contract_type"`
		OrderId       string `json:"order_id"`
		TradeId       string `json:"trade_id"`
		TradeVolume   string `json:"trade_volume"`
		TradePrice    string `json:"trade_price"`
		TradeTurnover string `json:"trade_turnover"`
		Role          string `json:"role"`
		ClientOrderId string `json:"client_order_id"`
		CreatedTime   string `json:"created_time"`
		UpdatedTime   string `json:"updated_time"`
	} `json:"data"`
}

type WsOrders struct {
	Op           string `json:"op"`
	Topic        string `json:"topic"`
	ContractCode string `json:"contract_code"`
	Ts           int64  `json:"ts"`
	Uid          string `json:"uid"`
	Data         struct {
		Side             string `json:"side"`
		Type             string `json:"type"`
		Price            string `json:"price"`
		Volume           string `json:"volume"`
		State            string `json:"state"`
		Id               string `json:"id"`
		ContractCode     string `json:"contract_code"`
		ContractType     string `json:"contract_type"`
		OrderId          string `json:"order_id"`
		PositionSide     string `json:"position_side"`
		PriceMatch       string `json:"price_match"`
		ClientOrderId    string `json:"client_order_id"`
		MarginMode       string `json:"margin_mode"`
		LeverRate        int64  `json:"lever_rate"`
		OrderSource      string `json:"order_source"`
		ReduceOnly       bool   `json:"reduce_only"`
		TimeInForce      string `json:"time_in_force"`
		CancelReason     string `json:"cancel_reason"`
		TradeId          string `json:"trade_id"`
		TradeVolume      string `json:"trade_volume"`
		TotalTradeVolume string `json:"total_trade_volume"`
		TradeAvgPrice    string `json:"trade_avg_price"`
		TradeTurnover    string `json:"trade_turnover"`
		Role             string `json:"role"`
		CreatedTime      string `json:"created_time"`
		UpdatedTime      string `json:"updated_time"`
		Fee              string `json:"fee"`
		FeeCurrency      string `json:"fee_currency"`
		SelfMatchPrevent string `json:"self_match_prevent"`
	} `json:"data"`
}
