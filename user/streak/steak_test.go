package streak

import (
	"testing"
)

func TestNew(t *testing.T) {
	scenario := []struct {
		Desc   string
		Input  int
		Output error
	}{
		{"zero", 0, nil},
		{"one", 1, nil},
		{"negative", -1, NegativeStreakErr},
	}
	for _, v := range scenario {
		_, err := New(v.Input)
		if v.Output != err {
			t.Errorf("Test \"%s\" failed: wanted %s, but got %s\n", v.Desc, v.Output, err)
		} else {
			t.Logf("Test \"%s\" has been passed!\n", v.Desc)
		}
	}
}

func TestStreak_String(t *testing.T) {
	scenario := []struct {
		Desc   string
		Input  int
		Output string
	}{
		{"zero", 0, "0"},
		{"one", 1, "1"},
		{"theAnswer", 42, "42"},
	}

	for _, v := range scenario {
		st, _ := New(v.Input)
		if st.String() != v.Output {
			t.Errorf("Test \"%s\" failed: wanted %s, but got %s\n", v.Desc, v.Output, st.String())
		} else {
			t.Logf("Test \"%s\" has been passed!\n", v.Desc)
		}
	}
}

func TestStreak_Increment(t *testing.T) {
	scenario := []struct {
		Desc   string
		Input  int
		Output string
	}{
		{"zero++", 0, "1"},
		{"one++", 1, "2"},
		{"theAnswer++", 42, "43"},
	}

	for _, v := range scenario {
		st, _ := New(v.Input)
		st.Increment()
		if st.String() != v.Output {
			t.Errorf("Test \"%s\" failed: wanted %s, but got %s\n", v.Desc, v.Output, st.String())
		} else {
			t.Logf("Test \"%s\" has been passed!\n", v.Desc)
		}
	}

}
