package main
import ("bankingapp/user")
func main() {
	user.CreateAdmin("bhavesh", "Mishra")

	admin := user.GetAdmin()
	admin.CreateUser("Rajesh", "sharma")

	user, _ := user.GetUserById(2)
	user.UpdateFirstName("Rakesh")

	user.DepositMoney(1, 1001, 1000)

	user.WithdrawMoney(1, 1001, 500)

	user.TransferMoney(1, 2, 1001, 2001, 200)

	user.GetBalance(1, 1001)

	user.GetPassbook(1, 1001, "2020-01-01", "2020-12-31")

	user.GetNetWorth()

	admin.ReadAllUsers()

 }

