//dataStructures
Cashier{
	CashierNO int
	Occupied bool
	NoOfCustomerServed int
}

Cashiers{
	chan Cashier  //as queue
}
Customer{
	customerNo int
}
//it don't need so not implemented as records
/*ServedCustomer{
	customerNo int
	cashierNO int
	ServiceStartTime time
}

AllServedCustomers{
	array{ServedCustomer}

}

(AllServedCustomers) Add(ServedCustomer) bool
*/
UnservedCustomersQueue{
	chan Customer //as Queue
	//NextCustomerNo int doesn't need it
	TotalCustomers int
}

(UnservedCustomersQueue) Add(customer) bool
(UnservedCustomersQueue) GetNextCustomer(customer) Customer

interface BankHandler{

NewCashier(cashierNO) Cashier
NewCustomer(customerNO) Customer
NewServedCustomer(customerNO,cashierNO,ServiceTime) ServedCustomer
NewUnservedCustomerQueue() 
StartCashierWork(Cashier,UnservedCustomerQueue)
}

func main(){
	customers:=NewUnservedCustomerQueue(NoofCustomers)
	cashiers:=getCashiers(NOofCashiers)
	for i:=1;i<NoOfCashiers;i++{
		go StartCashierWork(NewCashier(i),&customers)
}
}
