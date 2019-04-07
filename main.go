package main
import 	("time"
	"fmt"
	"os"
	"strings"
	"strconv"
)
type Cashier struct {
	CashierNo int
	NoOfCustomerServed int
}

type Cashiers struct{
	cashiers chan *Cashier
}

type  Customer struct {
	customerNo int
}

type CustomerQueue struct {
	customers chan *Customer
}

func NewCustomer(CustomerNo int) *Customer {
	return &Customer{CustomerNo}
}

func NewCustomerQueue(NoOfCustomers int) *CustomerQueue{
	queue:=&CustomerQueue{}
	queue.customers=make(chan *Customer,NoOfCustomers)
	for i:=1;i<=NoOfCustomers;i++{
		queue.customers<-NewCustomer(i)
	}
	return queue
}

func (cq CustomerQueue) Next() *Customer{
	var cust *Customer
	select {
	case cust=<-cq.customers:{
		return cust
		}
	default: {
		return nil
		}
	}
}
func GetCashiers(NoOfCashiers int) *Cashiers {
	c:=&Cashiers{}
	c.cashiers=make(chan *Cashier,NoOfCashiers)
	for i:=1;i<=NoOfCashiers;i++{
		c.cashiers<-&Cashier{i,0}
	}
	return c
}



func ServeCustomer(cash *Cashier,cust *Customer,duration int,cashs *Cashiers){

	fmt.Println(time.Now().Format("2006-01-02 15:04:05 "),"--> Cashier ",cash.CashierNo,": Customer ",cust.customerNo," Started")
	time.Sleep(time.Duration(duration)*time.Second)
	cash.NoOfCustomerServed=cash.NoOfCustomerServed+1
	fmt.Println(time.Now().Format("2006-01-02 15:04:05 "),"--> Cashier ",cash.CashierNo,": Customer ",cust.customerNo," Completed")
	cashs.cashiers<-cash //Add Cashier into free list
}

func main(){

	var Customers,Cashiers,ServiceTime int
	args:=os.Args
	if(len(args)<3){
		panic("command line arguments Not found..:)")
	}
	Customers,_=strconv.Atoi(strings.Split(args[2],"=")[1])
	Cashiers,_=strconv.Atoi(strings.Split(args[1],"=")[1])
	ServiceTime,_=strconv.Atoi(strings.Split(args[3],"=")[1])
	customerQueue:=NewCustomerQueue(Customers)
	cashiers:=GetCashiers(Cashiers)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05 "),"--> Bank Simulation Started")
	for true{
		cust:=customerQueue.Next()
		if(cust!=nil){
			cash:=<-cashiers.cashiers   //getFreeCashier
			go ServeCustomer(cash,cust,ServiceTime,cashiers)
		}else{
		break
		}

	}
	fmt.Println(time.Now().Format("2006-01-02 15:04:05")," --> Bank Simulated Ended")
}

