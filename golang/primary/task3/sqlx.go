package main 
import(
	"github.com/jmoiron/sqlx"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"time"
)	

type Employee struct {
	ID int `db:"id"`
	Name string `db:"name"`
	Department string `db:"department"`
	Salary float64 `db:"salary"`
}

func FindByDempartment(db *sqlx.DB, department string) ([]*Employee, error) {
	employees := make([]*Employee, 0)
	err := db.Select(&employees, "SELECT id, name, department, salary FROM employees WHERE department = ?", department)
	if err!= nil {
		return nil, err
	}
	return employees, nil
}

func FindTopSalary(db *sqlx.DB) (*Employee, error) {
	var employee Employee 
	query := "SELECT id, name, department, salary FROM employees ORDER BY salary DESC LIMIT 1"
	err := db.Get(&employee, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no employees found")
		}
		return nil, err
	}
	return &employee, nil
}

type Book struct {
	ID int `db:"id"`
	Title string `db:"title"`
	Author string `db:"author"`
	Price float64 `db:"price"`
}

func getBooksOverPrice(db *sqlx.DB, price float64) ([]*Book, error) {
	books := make([]*Book, 0)
	err := db.Select(&books, "SELECT id, title, author, price FROM `books` WHERE price > ?", price)
	if err!= nil {
		return nil, err
	}
	return books, nil
}

func main() {
	//dsn
	host := "localhost"
	user := "vinoctis"
	pwd  := "123456"
	database := "test"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pwd, host, database)
	fmt.Println(dsn)
	db, err := sqlx.Connect("mysql", dsn)
	if err!= nil {
		panic(err)
	}
	defer db.Close()

	db.SetMaxOpenConns(20) //最大连接数
	db.SetMaxIdleConns(10) //最大空闲连接数
	db.SetConnMaxLifetime(time.Hour) //连接最大存活时间

	//
	if err := db.Ping();err!= nil {
		panic(fmt.Sprintf("数据库连接失败:%v",err))
	}
	fmt.Println("数据库连接成功")

	employees, err := FindByDempartment(db, "技术部")
	if err!= nil {
		panic(err)
	}
	for _, employee := range employees {
		fmt.Printf("技术部员工：%v\n",employee)
	}

	topEmployee, err := FindTopSalary(db)
	if err!= nil {
		panic(err)
	}
	fmt.Printf("最高工资员工:%v\n", topEmployee)

	books, err := getBooksOverPrice(db, float64(50))
	if err!= nil {
		panic(err)
	}
	for _, book := range books {
		fmt.Printf("价格大于50的书籍：%v\n", book)
	}
	
}