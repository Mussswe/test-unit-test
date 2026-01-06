package backend

import (
	"testing"

	"github.com/asaskevich/govalidator"

	. "github.com/onsi/gomega"
)

func TestValidateUser_ValidUser(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run("กรอกข้อมูลที่ถูกต้อง", func(t *testing.T) {
		user := User{
			Sut_id: "C6600013",
			Name:   "Alice",
			Email:  "alice@example.com",
			Roles:  []Role{{RoleName: "Admin"}},
			Address: []Address{
				{City: "Phuket", PostCode: "83000"},
			},
		}
		valid, err := ValidateUser(user)
		g.Expect(err).To(BeNil())
		g.Expect(valid).To(BeTrue())
	})
	t.Run("กรอกชื่อไม่ถูก", func(t *testing.T) {
		user := User{
			Name:  "alice",
			Email: "alice@example.com",
			Roles: []Role{{RoleName: "Admin"}},
			Address: []Address{
				{City: "Phuket", PostCode: "83000"},
			},
		}
		valid, err := ValidateUser(user)
		g.Expect(err).NotTo(BeNil())
		g.Expect(valid).To(BeFalse())
	})
	t.Run("email ผิด", func(t *testing.T) {
		user := User{
			Name:  "alice",
			Email: "alice-example.com",
			Roles: []Role{{RoleName: "Admin"}},
			Address: []Address{
				{City: "Phuket", PostCode: "83000"},
			},
		}
		valid, err := ValidateUser(user)
		g.Expect(err).NotTo(BeNil())
		g.Expect(valid).To(BeFalse())
	})
	t.Run("reject case - empty name", func(t *testing.T) {
		user := User{
			Name:  "",
			Email: "bob@example.com",
			Roles: []Role{{RoleName: "Admin"}},
			Address: []Address{
				{City: "Bangkok", PostCode: "10110"},
				{City: "Phuket", PostCode: "83000"},
			},
		}
		valid, err := ValidateUser(user)
		g.Expect(err).ToNot(BeNil())
		g.Expect(valid).To(BeFalse())
	})
	t.Run("ใส่ sut id ผิด", func(t *testing.T) {
		user := User{
			Sut_id: "G1234567",
			Name:   "Bob",
			Email:  "bob@example.com",
			Roles:  []Role{{RoleName: "Admin"}},
			Address: []Address{
				{City: "Bangkok", PostCode: "10110"},
				{City: "Phuket", PostCode: "83000"},
			},
		}
		valid, err := ValidateUser(user)
		g.Expect(err).ToNot(BeNil())
		g.Expect(valid).To(BeFalse())
	})
	t.Run("ตรวจสอบ Error แยกตามฟิลด์", func(t *testing.T) {
		user := User{
			Name:   "Alice", // ผิด! เพราะเราตั้งไว้ว่าต้องตัวเล็กล้วน (^[a-z]+$)
			Sut_id: "A123",  // ผิด! เพราะต้องขึ้นต้นด้วย B, C, M และเลข 6 หลัก
		}

		_, err := ValidateUser(user)

		// แปลง error เป็น Map เพื่อให้เช็ครายฟิลด์ได้
		errMap := govalidator.ErrorsByField(err)

		// ตรวจสอบ Error ของฟิลด์ Name
		g.Expect(errMap["Name"]).NotTo(BeNil())

		// ตรวจสอบ Error ของฟิลด์ Sut_id
		g.Expect(errMap["Sut_id"]).To(ContainSubstring("does not validate as matches"))
	})
}
func TestUserAddressValidation(t *testing.T) {
	g := NewWithT(t)

	t.Run("ควรจะ Error เมื่อข้อมูลใน Address ผิดเงื่อนไข", func(t *testing.T) {
		user := User{
			Name:   "Alice",
			Sut_id: "B6312347",
			Email:  "alice@test.com",
			Address: []Address{
				{City: "Bangkok123", PostCode: "10110"}, // ผิดเพราะมีตัวเลข
			},
			Roles: []Role{{RoleName: "admin"}},
		}

		valid, err := ValidateUser(user)

		g.Expect(valid).To(BeFalse())
		g.Expect(err).NotTo(BeNil())

		// วิธีที่ชัวร์ที่สุดสำหรับการตรวจ Nested Error คือตรวจจาก String ของ err โดยตรง
		g.Expect(err.Error()).To(ContainSubstring("does not validate as alpha"))

		// หรือถ้าจะใช้ ErrorsByField ให้ลองเช็คที่ชื่อฟิลด์ของตัวลูก
		errs := govalidator.ErrorsByField(err)
		g.Expect(errs).NotTo(BeEmpty())
	})
}
