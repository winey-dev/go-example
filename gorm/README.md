# gorm
## ORM 이란?   
객체와 관계형 데이터베이스의 데이터를 자동으로 연결 해주는 것 

## GORM을 사용 해야 하는 이유 ?
- 개발자는 개발에만   
- DB 관련 Query를 직접적으로 개발자가 다루지 않아도 된다 ?  
- 객체를 통해서 특정 Method를 호출 시 DB관련 CRUD가 된다   

## 주의할 점 
- 객체를 통해서 DB에 접근하는 만큼 원하는 결과에 사용해야할 적절한 Method를 사용해야한다.
### ORM을 사용하지 않을 경우 
```Go
package main
import (
    ... 
)

const (
    INSERTUSER = "INSERT INTO User Values (?,?,?,?,?)"
    SELECTUSER = "SELECT * FROM User WHERE id=?"
)
func main() {

    // init DB
    // User Insert
    db.Exec(INSERTUSER,"id","password","email","phone","email")

    // User Select
    row, _ := db.QueryRow(SELECTUSER,"id")
    row.Result(&id, &password, &email, &phone, &email)
}
```
### ORM을 사용 할 경우
```Go
package main
import (
   	"gorm.io/gorm" 
    "gorm.io/driver/mysql"
)

func main() {

    // some dns example : "id:pwd@tcp(ip:port)/dbname?charset=utf8&parseTime=true"
    db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

    // User Insert
    db.Create(&User)

    // User Select
    db.Where("id = ?", "smlee").First(&u)
}
```