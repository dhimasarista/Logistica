package handlers

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

// Merchant ID G430193679
// Client ID SB-Mid-client-pZ40T3iL913MGQy1
// Server Key SB-Mid-server-zTZ2r8AhWDPPeBo7H8bWtssm

// func Init() {
// 	if os.Getenv("MIDTRANS_SERVER") == "" {
// 		log.Println("MIDTRANS_SERVER value not found.")
// 	}
// 	if os.Getenv("MIDTRANS_CLIENT") == "" {
// 		log.Println("MIDTRANS_CLIENT value not found.")
// 	}
// }

func OrderMidtrans(orderId string, price int) *snap.Response {
	// Init()
	midtrans.ServerKey = "SB-Mid-server-zTZ2r8AhWDPPeBo7H8bWtssm"
	midtrans.ClientKey = "SB-Mid-client-pZ40T3iL913MGQy1"
	midtrans.Environment = midtrans.Sandbox

	c := coreapi.Client{}
	c.New("SB-Mid-server-zTZ2r8AhWDPPeBo7H8bWtssm", midtrans.Sandbox)

	request := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderId,
			GrossAmt: int64(price),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
	}

	response, _ := snap.CreateTransaction(request)

	return response
}
