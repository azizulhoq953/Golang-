
//Method
package main

import "fmt"

/*type Doctor struct{

Name string

Education string

Age int
Salary float32

}*/


type Doctor struct{

Name string

Education string

Age string
Salary string


}
//method
/*func (d. Doctor)getSalaryinfo()string{
return d.Salary

}
*/
//method
func (c Doctor)getSalary()string{
return c.Salary

}
//method

func (e Doctor)getEducation()string{

return e.Education

}

func main(){

//var d=Doctor{ "Tarek","MBBS",30,55000,}

//var d=Doctor{Name:"Tarek",Education:"MBBS",Age:30,Salary:55000,}

/*var d Doctor

d.Name="Tarek"
d.Education="MBBS"
d.Age= 30
d.Salary=55000.00

fmt.Println(d.Name,d.Education,d.Age,d.Salary )
d.Speak()*/
//fmt.Println(d.Name,d.Education,d.Age,d.Salary)
var c Doctor

c.Name=" Name: Tareq\n"
c.Education=" Education: MBBS\n"
c.Age="Age: 30\n"
c.Salary="Salary: 50000.04\n"

var e Doctor

e.Name=" Name: Tareq\n"
e.Education=" Education: MBBS\n"
e.Age="Age: 30\n"
e.Salary="Salary: 50000.04\n"

//fmt.Println(c.Name,c.Age,c.Education,c.Salary)

fmt.Println(c.getSalary())
fmt.Println(e.getEducation())


}