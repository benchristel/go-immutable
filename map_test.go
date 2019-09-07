package immutable_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/benchristel/immutable-map"
)

var _ = Describe("immutable.Map", func() {
	It("starts empty", func() {
		m := immutable.NewMap()
		Expect(m.Size()).To(BeEquivalentTo(0))
		Expect(m.HasKey("nonexistent")).To(BeFalse())
	})

	Describe("Size()", func() {
		It("increments after adding an item", func() {
			m := (immutable.NewMap().
				Put("a", 1))
			Expect(m.Size()).To(Equal(1))
		})

		It("increments after adding a second item", func() {
			m := immutable.NewMap().
				Put("a", 1).
				Put("b", 2)
			Expect(m.Size()).To(Equal(2))
		})

		It("does not change when overwriting an entry", func() {
			m := immutable.NewMap().
				Put("same", 1).
				Put("same", 2)
			Expect(m.Size()).To(Equal(1))
		})

		It("decrements after removing an item", func() {
			m := immutable.NewMap().
				Put("a", 1).
				Remove("a")
			Expect(m.Size()).To(Equal(0))
		})
	})

	Describe("HasKey()", func() {
		It("returns false for an empty Map", func() {
			m := immutable.NewMap()
			Expect(m.HasKey("nonexistent")).To(BeFalse())
		})

		It("returns true for a key that has been Put()", func() {
			m := immutable.NewMap().
				Put("key", 1)
			Expect(m.HasKey("key")).To(BeTrue())
		})

		It("returns false for a key that has been Remove()d", func() {
			m := immutable.NewMap().
				Put("key", 1).
				Remove("key")
			Expect(m.HasKey("key")).To(BeFalse())
		})
	})

	Describe("Get()", func() {
		It("returns nil for a key that's not in the Map", func() {
			m := immutable.NewMap()
			Expect(m.Get("nonexistent")).To(BeNil())
		})

		It("returns a value that was previously Put()", func() {
			m := immutable.NewMap().
				Put("key", 1)
			Expect(m.Get("key")).To(Equal(1))
		})

		It("returns the most recent value that was Put()", func() {
			m := immutable.NewMap().
				Put("key", 1).
				Put("key", 2)
			Expect(m.Get("key")).To(Equal(2))
		})
	})

	Describe("Put", func() {
		It("doesn't change the Size() of the original Map", func() {
			original := immutable.NewMap()
			original.Put("a", 1)
			Expect(original.Size()).To(Equal(0))
		})

		It("doesn't change the set of values stored in the original Map", func() {
			original := immutable.NewMap()
			original.Put("a", 1)
			Expect(original.Get("a")).To(BeNil())
		})

		It("doesn't change the set of keys stored in the original Map", func() {
			original := immutable.NewMap()
			original.Put("a", 1)
			Expect(original.HasKey("a")).To(BeFalse())
		})
	})

	Describe("Remove", func() {
		It("doesn't change the Size() of the original Map", func() {
			original := immutable.NewMap().Put("key", 1)
			original.Remove("key")
			Expect(original.Size()).To(Equal(1))
		})

		It("doesn't change the set of values stored in the original Map", func() {
			original := immutable.NewMap().Put("key", 1)
			original.Remove("key")
			Expect(original.Get("key")).To(Equal(1))
		})

		It("doesn't change the set of keys stored in the original Map", func() {
			original := immutable.NewMap().Put("key", 1)
			original.Remove("key")
			Expect(original.HasKey("key")).To(BeTrue())
		})
	})
})
