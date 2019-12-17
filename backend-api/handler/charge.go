package handler

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"strconv"

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
		log.Fatalln(err)
	}
	identifer, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = c.JSON(http.StatusInternalServerError, err)
		log.Println(err)
	}

	// id から item情報取得
	res, err := db.SelectItem(int64(identifer))
	if err != nil {
		err = c.JSON(http.StatusInternalServerError, err)
		log.Println(err)
	}
	// gRPC サーバーに送る Request を作成
	payReq := &paypb.PayReq{
		Id:          int64(identifer),
		Token:       t.Token,
		Amount:      res.Amount,
		Name:        res.Name,
		Description: res.Description,
	}

	//IPアドレス(ここではlocalhost)とポート番号(ここでは50051)を指定して、サーバーと接続する
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		err = c.JSON(http.StatusForbidden, err)
		log.Println(err)
	}
	defer conn.Close()
	client := paypb.NewPayManagerClient(conn)

	// gRPCマイクロサービスの支払い処理関数を叩く
	gres, err := client.Charge(context.Background(), payReq)
	if err != nil {
		err = c.JSON(http.StatusForbidden, err)
		log.Println(err)
		return
	}
	err = c.JSON(http.StatusOK, gres)
	log.Println(err)
}
