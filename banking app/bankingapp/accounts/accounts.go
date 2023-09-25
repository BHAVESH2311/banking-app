package accounts
import (
	"fmt"
	"time"
    "bankingapp/passbook"
)

var id = uint(10000)
var accNumber = uint(10000)

type Account struct {
	id            uint   
	bankName      string
	balance       float64  
	holderName    string
	accountType string
	accNumber uint
	isActive bool
	passbook []*passbook.Passbook
	
}

func CreateAccount(bankName,holderName,accountType string,) *Account {

	id++
	accNumber++

	return &Account{
		id: id, 
		bankName: bankName,
		balance: 0,
		holderName : holderName,
		accNumber: accNumber,
		accountType : accountType,
		isActive: true,
		passbook: []*passbook.Passbook{},
	}
}

func (account *Account)Withdraw( amount float64) bool {
	if account.balance < amount {
	fmt.Println("insufficient funds")
	}
	account.balance -= amount
	return true
}

func (account *Account)Deposit(amount float64) {
		account.balance += amount 
}


func (account *Account) GetAccountNum() uint {
	return account.accNumber
}

func (account *Account) GetId() uint{
	return account.id
}


func (a *Account) GetPassbook(fromDate, toDate interface{}) []*passbook.Passbook {
	from, err := time.Parse("2006-01-02", fromDate.(string))

	if err != nil {
		fmt.Println(err)
		return nil
	}

	to, err := time.Parse("2006-01-02", toDate.(string))

	if err != nil {
		fmt.Println(err)
		return nil
	}

	if fromDate == nil && toDate == nil {
		return a.passbook
	}

	if fromDate == nil {
		var passbookLogs []*passbook.Passbook
		for _, passbook := range a.passbook {
			if to.Before(passbook.GetTime()) {
				passbookLogs = append(passbookLogs, passbook)
			}
		}

		return passbookLogs
	}

	if toDate == nil {
		var passbookLogs []*passbook.Passbook
		for _, passbook := range a.passbook {
			if from.After(passbook.GetTime()) {
				passbookLogs = append(passbookLogs, passbook)
			}
		}

		return passbookLogs
	}

	var passbookLogs []*passbook.Passbook
	for _, passbook := range a.passbook {
		if from.After(passbook.GetTime()) && to.Before(passbook.GetTime()) {
			passbookLogs = append(passbookLogs, passbook)
		}
	}

	return passbookLogs
}


func (account *Account) GetBankName() string {
	return account.bankName
}

func (account *Account) GetAccBalance() uint {
	return uint(account.balance)
}