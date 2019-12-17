package handler

import (
	"google.golang.org/grpc"
	"net/http"

	"github.com/tetsuzawa/vue-go-pay-tutorial/backend-api/db"
	"github.com/tetsuzawa/vue-go-pay-tutorial/backend-api/domain"
	paypb "github.com/tetsuzawa/vue-go-pay-tutorial/payment-service/protocols/pay"
)

var addr = "localhost:50051"

// Charge exec payment-service charge
func Charge(c Context) {
	//パラメータや body をうけとる
	t := domain.Payment{}
	err := c.Bind(&t)
	if err != nil {
		err = c.JSON(http.StatusInternalServerError, err)
		panic(err)
	}
	identifer := c.Param("id")

	// id から item情報取得
	res, err := db.SelectItem(identifer)
	if err != nil {
		err = c.JSON(http.StatusInternalServerError, err)
		panic(err)
	}
	// gRPC サーバーに送る Request を作成
	pay := &paypb.PayReq{
		Id:          identifer,
		Token:       t.Token,
		Amount:      res.Amount,
		Name:        res.Name,
		Description: res.Description,
	}

	//IPアドレス(ここではlocalhost)とポート番号(ここでは50051)を指定して、サーバーと接続する
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		c.JSON(http.StatusForbidden, err)
	}
	defer conn.Close()
	client := pay.NewPayManagerClient(conn)

	// gRPCマイクロサービスの支払い処理関数を叩く
	gres, err := client.Charge(context.Background(), pay)
	if err != nil {
		c.JSON(http.StatusForbidden, err)
		return
	}
	c.JSON(http.StatusOK, gres)
}
