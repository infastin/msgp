package _generated

import (
	"testing"
	"time"
)

func compareStructD(t *testing.T, a, b *CompatibleStructD) {
	t.Helper()

	if !time.Time(a.Time).Equal(time.Time(b.Time)) {
		t.Fatal("not same time")
	}

	if a.Duration != b.Duration {
		t.Fatal("not same duration")
	}

	if len(a.MapString) != len(b.MapString) {
		t.Fatal("not same map")
	}

	for k, v1 := range a.MapString {
		if v2, ok := b.MapString[k]; !ok || v1 != v2 {
			t.Fatal("not same map")
		}
	}
}

func compareStructC(t *testing.T, a, b *CompatibleStructC) {
	t.Helper()

	if a.Float32 != b.Float32 {
		t.Fatal("not same float32")
	}

	if a.Float64 != b.Float64 {
		t.Fatal("not same float64")
	}

	compareStructD(t, (*CompatibleStructD)(&a.StructD), (*CompatibleStructD)(&b.StructD))
}

func compareStructB(t *testing.T, a, b *CompatibleStructB) {
	t.Helper()

	if a.Array8 != b.Array8 {
		t.Fatal("not same array")
	}

	if a.Any != b.Any {
		t.Fatal("not same any")
	}

	compareStructC(t, (*CompatibleStructC)(&a.StructC), (*CompatibleStructC)(&b.StructC))
}

func compareStructA(t *testing.T, a, b *CompatibleStructA) {
	t.Helper()

	if a.Int != b.Int {
		t.Fatal("not same int")
	}

	compareStructB(t, (*CompatibleStructB)(&a.StructB), (*CompatibleStructB)(&b.StructB))
}

func compareStructI(t *testing.T, a, b *CompatibleStructI) {
	t.Helper()

	if *a.Int != *b.Int {
		t.Fatal("not same int")
	}

	if *a.Uint != *b.Uint {
		t.Fatal("not same uint")
	}
}

func compareStructS(t *testing.T, a, b *CompatibleStructS) {
	t.Helper()

	if len(a.Slice) != len(b.Slice) {
		t.Fatal("not same slice")
	}

	for i := 0; i < len(a.Slice); i++ {
		if a.Slice[i] != b.Slice[i] {
			t.Fatal("not same slice")
		}
	}
}

func TestReplace_ABCD(t *testing.T) {
	d := CompatibleStructD{
		Time:     Time(time.Now()),
		Duration: Duration(time.Duration(1234)),
		MapString: map[string]string{
			"foo":   "bar",
			"hello": "word",
			"baz":   "quux",
		},
	}

	c := CompatibleStructC{
		StructD: StructD(d),
		Float32: 1.0,
		Float64: 2.0,
	}

	b := CompatibleStructB{
		StructC: StructC(c),
		Any:     "sup",
		Array8:  [8]byte{'f', 'o', 'o'},
	}

	a := CompatibleStructA{
		StructB: StructB(b),
		Int:     10,
	}

	t.Run("D", func(t *testing.T) {
		bytes, err := d.MarshalMsg(nil)
		if err != nil {
			t.Fatal(err)
		}

		ud := CompatibleStructD{}

		_, err = ud.UnmarshalMsg(bytes)
		if err != nil {
			t.Fatal(err)
		}

		compareStructD(t, &d, &ud)
	})

	t.Run("C", func(t *testing.T) {
		bytes, err := c.MarshalMsg(nil)
		if err != nil {
			t.Fatal(err)
		}

		uc := CompatibleStructC{}

		_, err = uc.UnmarshalMsg(bytes)
		if err != nil {
			t.Fatal(err)
		}

		compareStructC(t, &c, &uc)
	})

	t.Run("B", func(t *testing.T) {
		bytes, err := b.MarshalMsg(nil)
		if err != nil {
			t.Fatal(err)
		}

		ub := CompatibleStructB{}

		_, err = ub.UnmarshalMsg(bytes)
		if err != nil {
			t.Fatal(err)
		}

		compareStructB(t, &b, &ub)
	})

	t.Run("A", func(t *testing.T) {
		bytes, err := a.MarshalMsg(nil)
		if err != nil {
			t.Fatal(err)
		}

		ua := CompatibleStructA{}

		_, err = ua.UnmarshalMsg(bytes)
		if err != nil {
			t.Fatal(err)
		}

		compareStructA(t, &a, &ua)
	})
}

func TestReplace_I(t *testing.T) {
	var int0 int = -10
	var uint0 uint = 12

	i := CompatibleStructI{
		Int:  (*Int)(&int0),
		Uint: (*Uint)(&uint0),
	}

	bytes, err := i.MarshalMsg(nil)
	if err != nil {
		t.Fatal(err)
	}

	ui := CompatibleStructI{}

	_, err = ui.UnmarshalMsg(bytes)
	if err != nil {
		t.Fatal(err)
	}

	compareStructI(t, &i, &ui)
}

func TestReplace_S(t *testing.T) {
	s := CompatibleStructS{
		Slice: []Int{10, 12, 14, 16},
	}

	bytes, err := s.MarshalMsg(nil)
	if err != nil {
		t.Fatal(err)
	}

	us := CompatibleStructS{}

	_, err = us.UnmarshalMsg(bytes)
	if err != nil {
		t.Fatal(err)
	}

	compareStructS(t, &s, &us)
}

func TestReplace_Dummy(t *testing.T) {
	dummy := Dummy{
		StructA: StructA{
			StructB: StructB{
				StructC: StructC{
					StructD: StructD{
						Time:     Time(time.Now()),
						Duration: Duration(time.Duration(1234)),
						MapString: map[string]string{
							"foo":   "bar",
							"hello": "word",
							"baz":   "quux",
						},
					},
					Float32: 1.0,
					Float64: 2.0,
				},
				Any:    "sup",
				Array8: [8]byte{'f', 'o', 'o'},
			},
			Int: 10,
		},
		StructI: StructI{
			Int:  new(Int),
			Uint: new(Uint),
		},
		StructS: StructS{
			Slice: []Int{10, 12, 14, 16},
		},
		Uint:   10,
		String: "cheese",
	}

	*dummy.StructI.Int = 1234
	*dummy.StructI.Uint = 555

	bytes, err := dummy.MarshalMsg(nil)
	if err != nil {
		t.Fatal(err)
	}

	udummy := Dummy{}

	_, err = udummy.UnmarshalMsg(bytes)
	if err != nil {
		t.Fatal(err)
	}

	compareStructA(t, (*CompatibleStructA)(&dummy.StructA), (*CompatibleStructA)(&udummy.StructA))
	compareStructI(t, (*CompatibleStructI)(&dummy.StructI), (*CompatibleStructI)(&udummy.StructI))
	compareStructS(t, (*CompatibleStructS)(&dummy.StructS), (*CompatibleStructS)(&udummy.StructS))

	if dummy.Uint != udummy.Uint {
		t.Fatal("not same uint")
	}

	if dummy.String != udummy.String {
		t.Fatal("not same string")
	}
}
