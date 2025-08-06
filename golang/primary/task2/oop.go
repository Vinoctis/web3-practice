package main 
import(
	"fmt"
)

// 题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。
// 然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
// 在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
// 考察点 ：接口的定义与实现、面向对象编程风格。
type Shape interface{
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width 	float64
	Height 	float64
}

func NewRectangle(w, h float64) Rectangle{
	return Rectangle{
		Width: w,
		Height: h,
	}
}

func(r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func NewCircle(r float64) Circle{
	return Circle{
		Radius: r,
	}
}


func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

// 题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，
// 再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
// 考察点 ：组合的使用、方法接收者。
type Person struct {
	Name string 
	Age int 
}

type Employee struct {
	EmployeeID int
	Person
}

func NewEmployee(id int, p Person) Employee {
	return Employee{
		EmployeeID: id,
		Person: p,
	}
}

func (e Employee) PrintInfo(id int) {
	fmt.Println("员工姓名：", e.Name)
	fmt.Println("员工年龄：", e.Age)
	fmt.Println("员工ID：", e.EmployeeID)
}


func main() {
	rectangle := NewRectangle(10, 20)
	fmt.Println("矩形面积：", rectangle.Area())
	fmt.Println("矩形周长：", rectangle.Perimeter())

	circle := NewCircle(5.5)
	fmt.Println("圆面积：", circle.Area())
	fmt.Println("圆周长：", circle.Perimeter())

	p:= Person{
		Name: "Vinoctis",
		Age: 18,
	}
	e:= NewEmployee(1, p)
	e.PrintInfo(1)
}