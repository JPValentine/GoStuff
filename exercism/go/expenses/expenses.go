package expenses


import(
	"errors"
	"fmt"
)


// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	var out []Record
	for _,v := range in{
		if predicate(v){out=append(out, v)}
	}//for
	return out
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(r Record) bool{
		flag:=false
		for i:=p.From;p.To+1>i;i++{
			if i==r.Day{
				flag=true
			}
		}
		return flag
	}
}

// bycategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool{
		flag:=false
		if r.Category==c{
			flag=true
		}
                return flag
        }

}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	total:=0.0
	for j:=0;len(in)>j;j++{
		for i:=p.From;p.To+1>i;i++{
			if i==in[j].Day{
				total+=in[j].Amount
			}
		}//i
	}//j
        return total
}


// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	total:=0.0
	flag:=false
	for x:=0;len(in)>x;x++{
		if in[x].Category == c{ flag = true}
	}
        for j:=0;len(in)>j;j++{
                for i:=p.From;p.To+1>i;i++{
                        if i==in[j].Day && in[j].Category == c{
                                total+=in[j].Amount
                        }
                }//i
        }//j
	if flag{
		fmt.Println(total)
        	return total,nil
	}
        return total,errors.New(fmt.Sprintf("unknown category %s",c))
	
}
