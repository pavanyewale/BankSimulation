//dataStructures
Cashier{
	CashierNO int
	Occupied bool
	NoOfCustomerServed int
	WorkStartTime int
}

Customer{
	customerNo int
}

ServedCustomer{
	customerNo int
	cashierNO int
	ServiceStartTime time
}

AllServedCustomers{
	array{ServedCustomer}

}

(AllServedCustomers) Add(ServedCustomer) bool

UnservedCustomersQueue{
	array{Customer}
	NextCustomerNO int
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


