package example

import (
	"context"
	"fmt"
)


//Goods -
type Goods struct {
	ID   int64
	Name string
}

//Order -
type Order struct {
	OrderID int64
	Name    string
}

type Req struct {

}

type Res struct{
	Order Order
	Goods Goods
}

//GetGoodsByID -
func (g *Goods) GetGoodsByID(ctx context.Context, req *Req, res *Res) error {
	res.Goods = Goods{ID: 10001, Name: "苹果"}
	return nil
}

//GetGoodsByID -
func (o *Order) GetOrderByID(ctx context.Context, req *Req, res *Res) error {
	res.Order = Order{OrderID: 10002, Name: "购买剃须刀"}
	return nil
}

func NewPreKey(servername string) string{
	return fmt.Sprintf("/pre%v",servername)
}
