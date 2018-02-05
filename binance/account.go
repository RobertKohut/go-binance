/*

   account.go
       Account (Signed) Endpoints for Binance Exchange API

*/
package binance

import (
	"fmt"
)

// GetAccountInfo retrieves basic account information
func (b *Binance) GetAccountInfo() (account Account, err error) {

	reqUURL := fmt.Sprintf("api/v3/account")

	_, err = b.client.do("GET", reqUURL, "", true, &account)
	if err != nil {
		return
	}

	return
}

// GetPositions Filters basic account information to retrieve current holdings
func (b *Binance) GetPositions() (positions []Balance, err error) {

	reqURL := fmt.Sprintf("api/v3/account")
	account := Account{}

	_, err = b.client.do("GET", reqURL, "", true, &account)
	if err != nil {
		return
	}

	positions = make([]Balance, len(account.Balances))
	i := 0

	for _, balance := range account.Balances {
		if balance.Free != 0.0 || balance.Locked != 0.0 {
			positions[i] = balance
			i++
		}
	}

	return positions[:i], nil
}

// PlaceLimitOrder Places a Limit Order
func (b *Binance) PlaceLimitOrder(l LimitOrder) (res PlacedOrder, err error) {

	err = l.ValidateLimitOrder()
	if err != nil {
		return
	}

	reqURL := fmt.Sprintf("api/v3/order?symbol=%s&side=%s&type=%s&timeInForce=%s&quantity=%f&price=%f&recvWindow=%d", l.Symbol, l.Side, l.Type, l.TimeInForce, l.Quantity, l.Price, l.RecvWindow)

	_, err = b.client.do("POST", reqURL, "", true, &res)
	if err != nil {
		return
	}

	return
}

// PlaceMarketOrder places a Market Order
func (b *Binance) PlaceMarketOrder(m MarketOrder) (res PlacedOrder, err error) {

	err = m.ValidateMarketOrder()
	if err != nil {
		return
	}

	reqURL := fmt.Sprintf("api/v3/order?symbol=%s&side=%s&type=%s&quantity=%f&recvWindow=%d", m.Symbol, m.Side, m.Type, m.Quantity, m.RecvWindow)

	_, err = b.client.do("POST", reqURL, "", true, &res)
	if err != nil {
		return
	}

	return
}

// CancelOrder cancels an Order
func (b *Binance) CancelOrder(query OrderQuery) (order CanceledOrder, err error) {

	err = query.ValidateOrderQuery()
	if err != nil {
		return
	}

	reqURL := fmt.Sprintf("api/v3/order?symbol=%s&orderId=%d&recvWindow=%d", query.Symbol, query.OrderId, query.RecvWindow)

	_, err = b.client.do("DELETE", reqURL, "", true, &order)
	if err != nil {
		return
	}

	return
}

// CheckOrder checks the Status of an Order
func (b *Binance) CheckOrder(query OrderQuery) (status OrderStatus, err error) {

	err = query.ValidateOrderQuery()
	if err != nil {
		return
	}

	reqURL := fmt.Sprintf("api/v3/order?symbol=%s&orderId=%d&origClientOrderId=%s&recvWindow=%d", query.Symbol, query.OrderId, query.RecvWindow)

	_, err = b.client.do("GET", reqURL, "", true, &status)
	if err != nil {
		return
	}

	return
}

// GetAllOpenOrders retrieves all open orders
func (b *Binance) GetAllOpenOrders() (orders []OrderStatus, err error) {
	_, err = b.client.do("GET", "api/v3/openOrders", "", true, &orders)

	if err != nil {
		return
	}

	return
}

// GetOpenOrders retrieves all open orders for a given symbol
func (b *Binance) GetOpenOrders(query OpenOrdersQuery) (orders []OrderStatus, err error) {

	err = query.ValidateOpenOrdersQuery()
	if err != nil {
		return
	}
	reqURL := fmt.Sprintf("api/v3/openOrders?symbol=%s&recvWindow=%d", query.Symbol, query.RecvWindow)
	_, err = b.client.do("GET", reqURL, "", true, &orders)
	if err != nil {
		return
	}

	return
}

// GetTrades retrieves all trades
func (b *Binance) GetTrades(symbol string) (trades []Trade, err error) {

	reqURL := fmt.Sprintf("api/v3/myTrades?symbol=%s", symbol)

	_, err = b.client.do("GET", reqURL, "", true, &trades)

	if err != nil {
		return
	}
	return
}

// GetWithdrawHistory retrieves all withdrawals
func (b *Binance) GetWithdrawHistory() (withdraws WithdrawList, err error) {

	reqURL := fmt.Sprintf("wapi/v3/withdrawHistory.html")

	_, err = b.client.do("GET", reqURL, "", true, &withdraws)
	if err != nil {
		return
	}
	return
}

// GetDepositHistory retrieves all deposits
func (b *Binance) GetDepositHistory() (deposits DepositList, err error) {

	reqURL := fmt.Sprintf("wapi/v3/depositHistory.html")

	_, err = b.client.do("GET", reqURL, "", true, &deposits)
	if err != nil {
		return
	}
	return
}
