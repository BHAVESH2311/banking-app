package passbook

import "time"

type Passbook struct {
	timestamp             time.Time
	senderId              uint
	receiverId            uint
	senderAccNumber   uint
	receivedAccNumber uint
	amount                uint
	transactionType       string
}

func CreatePassbook(senderId, receiverId, senderAccNumber, receivedAccNumber, amount uint, transactionType string) *Passbook{
	return &Passbook{
		timestamp:             time.Now(),
		senderId:              senderId,
		receiverId:            receiverId,
		senderAccNumber:   senderAccNumber,
		receivedAccNumber: receivedAccNumber,
		amount:                amount,
		transactionType:       transactionType,
	}
}

func (p *Passbook) GetTime() time.Time {
	return p.timestamp
}