package validation_test
import (
	"testing"
	. "github.com/onsi/gomega"
	"github.com/peraponglim1125/Pre9/entity"
)
func TestValid (t *testing.T){
	g := NewGomegaWithT(t)
	student := entity.Student{
		Fullname: "John Doe",
		Username: "B12345",
		Age: 20,
		GPA: 3.5,	
	}
	ok, err := entity.ValidationStudent(&student)
	g.Expect(ok).To(BeTrue())
	g.Expect(err).To(BeNil())
}
func TestInValidName (t *testing.T){
	t.Run (`Fullname is required` , func(t *testing.T){
		g := NewGomegaWithT(t)
		student := entity.Student{
			Fullname: "",
			Username: "B12345",
			Age: 20,
			GPA: 3.5,
		}
		ok , err := entity.ValidationStudent(&student)
		g.Expect(ok).To(BeFalse())
		g.Expect(err.Error()).To(Equal("Fullname is required"))

	})
	t.Run (`Fullname is not in range` , func(t *testing.T){
		g := NewGomegaWithT(t)
		student := entity.Student{
			Fullname: "hifoiwefwpefjwefpwojepfwefefew",
			Username: "B12345",
			Age: 20,
			GPA: 3.5,
		}
		ok , err := entity.ValidationStudent(&student)
		g.Expect(ok).To(BeFalse())
		g.Expect(err.Error()).To(Equal("Fullname is not in range"))

	})
	
}
func TestInValidUsername (t *testing.T){
	g := NewGomegaWithT(t)
	student := entity.Student{
		Fullname: "John Doe",
		Username: "BV12345fds",
		Age: 20,
		GPA: 3.5,	
	}
	ok, err := entity.ValidationStudent(&student)
	g.Expect(ok).To(BeFalse())
	g.Expect(err.Error()).To(Equal("Username is not in this formal"))
}

func TestInValidAge (t *testing.T){
	g := NewGomegaWithT(t)
	student := entity.Student{
		Fullname: "John Doe",
		Username: "B12345",
		Age: 2,
		GPA: 3.5,	
	}
	ok, err := entity.ValidationStudent(&student)
	g.Expect(ok).To(BeFalse())
	g.Expect(err.Error()).To(Equal("Age is not in range"))
}
func TestInValidGpa (t *testing.T){
	g := NewGomegaWithT(t)
	student := entity.Student{
		Fullname: "John Doe",
		Username: "B12345",
		Age: 20,
		GPA: 55.0,
	}
	ok, err := entity.ValidationStudent(&student)
	g.Expect(ok).To(BeFalse())
	g.Expect(err.Error()).To(Equal("GPA is not in range"))
}
