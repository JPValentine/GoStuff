package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	out:=map[string]int{}
	out["quarter_of_a_dozen"]=3
	out["half_of_a_dozen"]=6
	out["dozen"]=12
	out["small_gross"]=120
	out["gross"]=144
	out["great_gross"]=1728
	return out
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill:=map[string]int{}
	return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	_,uExists := units[unit]
	if uExists == false{
		return false
	}

	_,bExists:= bill[item]
	if bExists{
		bill[item]+=units[unit]
	}else{
		bill[item]=units[unit]
	}
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_,uExists:=units[unit]
	_,bExists:=bill[item]
	if bExists==false || uExists==false{
		return false
	}

	x:=bill[item]-units[unit]
	if x<0{
		return false
	}
	if x==0{
		delete(bill,item)
		return true
	}else{
		bill[item]-=units[unit]
		return true
	}
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	_,exist:=bill[item]
	if exist==false{
		return 0,false
	}
	qty:=bill[item]
	return qty,true
}
