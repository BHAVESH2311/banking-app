package bank
// Bank struct
import (
	"bankingapp/accounts"
	"bankingapp/passbook"
	"fmt"
)

var id = uint(0)

type Bank struct {
	id uint 
	Name string
	Abbreviation string 
	netWorth int
	accounts []*accounts.Account
}


func CreateBank(name, abbreviation string) *Bank {
	id++
 
	return &Bank {
		id: id,
		Name: name,
		Abbreviation: abbreviation,
		netWorth: 0,
		accounts: []*accounts.Account{},
	}
	
}


func (bank *Bank) CreateAccount(holderName,accountType string) *accounts.Account {
      newAccount := accounts.CreateAccount(bank.Name,holderName,accountType)
	  bank.accounts = append(bank.accounts, newAccount)
	  return newAccount
}


func (bank *Bank) DepositMoneyIntoAccount(acId uint,amount int){
	for _,i := range bank.accounts{
		if(i.GetId()==acId){
			i.Deposit(float64(amount))
			return
		}
	}
	return
}


func (bank *Bank) WithdrawMoneyFromAccount(acId uint,amount int){
	for _,i := range bank.accounts{
		if(i.GetId()==acId){
			i.Withdraw(float64(amount))
			return
		}
	}
	return
}

func (b *Bank) getAccount(accNum uint) *accounts.Account {
	for _, account := range b.accounts {
		if account.GetAccountNum() == accNum {
			return account
		}
	}

	return nil
}


func (b *Bank) TransferMoney(srcAcc, desAcc uint, amount int) {
	var srcAccount *accounts.Account
	var desAccount *accounts.Account

	for _, j := range b.accounts {
		if j.GetId() == uint(srcAcc) {
			srcAccount = j
		}
	}

	for _, j := range b.accounts {
		if j.GetId() == uint(desAcc) {
			desAccount = j
		}
	}

	srcAccount.Withdraw(float64(amount))
	desAccount.Deposit(float64(amount))
}


func (b *Bank) GetAccBalance(accNum uint) uint {
	account := b.getAccount(accNum)

	if account == nil {
		fmt.Println("Account not found")
	}

	return account.GetAccBalance()
}

// func (b *Bank) UpdateAccount(accNum uint, name string) {
// 	account := b.getAccount(accNum)

// 	if account == nil {
// 		fmt.Println("Account not found")
// 	}

// 	account.UpdateAccount(name)
// }



func (b *Bank) GetBalance(accNum uint) uint {
	account := b.getAccount(accNum)

	if account == nil {
	}

	return account.GetAccBalance()
}


func (b *Bank) GetNetWorth() float64 {
	worth := 0.0
	for _, bank := range b.accounts {
		worth += float64(bank.GetAccBalance())
	}

	return worth
}


func (b *Bank) GetBank(id uint) uint {
	return b.id
}



func (b *Bank) GetPassbook(accNum uint, fromDate, toDate string) []*passbook.Passbook {
	account := b.getAccount(accNum)

	passbooks := account.GetPassbook(fromDate, toDate)
	return passbooks
}












