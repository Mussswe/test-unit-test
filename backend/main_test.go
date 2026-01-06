package backend

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestValidateUser_ValidUser(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run("accept case ", func(t *testing.T) {
		user := User{
			Name:  "Alice",
			Email: "alice@example.com",
		}
		valid, err := ValidateUser(user)
		g.Expect(err).To(BeNil())
		g.Expect(valid).To(BeTrue())
	})
	t.Run("reject case - invalid email", func(t *testing.T) {
		user := User{
			Name:  "Bob",
			Email: "invalid-email",
		}
		valid, err := ValidateUser(user)
		g.Expect(err).ToNot(BeNil())
		g.Expect(valid).To(BeFalse())
	})
	t.Run("reject case - empty name", func(t *testing.T) {
		user := User{
			Name:  "",
			Email: "bob@example.com",
		}
		valid, err := ValidateUser(user)
		g.Expect(err).ToNot(BeNil())
		g.Expect(valid).To(BeFalse())
	})
}
