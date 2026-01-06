package backend

import (
	"testing"

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
}
