package entity
import "github.com/asaskevich/govalidator"
type Student struct{
	Fullname string `valid:"required~Fullname is required, stringlength(8|10)~Fullname is not in range"`
	Username string `valid:"matches(^[BC]\\d{5}$)~Username is not in this formal"`
	Age int `valid:"range(10|100)~Age is not in range"`
	GPA float32 `valid:"range(0|4)~GPA is not in range"`
}
func ValidationStudent (student *Student) (bool ,error){
	return govalidator.ValidateStruct(student)
}