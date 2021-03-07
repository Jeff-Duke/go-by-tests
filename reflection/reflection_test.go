package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	const name = "Jeff"
	const hometown = "Denver"
	const age = 37

	type Profile struct {
		Age  int
		City string
	}

	type Person struct {
		Name    string
		Profile Profile
	}

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{name},
			[]string{name},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{name, hometown},
			[]string{name, hometown},
		},
		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{name, age},
			[]string{name},
		},
		{
			"Struct with nested fields",
			Person{
				name,
				Profile{age, hometown},
			},
			[]string{name, hometown},
		},
		{
			"Pointers to things",
			&Person{
				name,
				Profile{age, hometown},
			},
			[]string{name, hometown},
		},
		{
			"Slices",
			[]Profile{
				{33, "London"},
				{age, hometown},
			},
			[]string{"London", hometown},
		},
		{
			"Arrays",
			[2]Profile{
				{33, "London"},
				{age, hometown},
			},
			[]string{"London", hometown},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	// Maps need to be tested separately where we can not care about the order the calls are run in
	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Berlin"}
			aChannel <- Profile{34, "Queens"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Berlin", "Queens"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Berlin"}, Profile{34, "Queens"}
		}

		var got []string
		want := []string{"Berlin", "Queens"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false

	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it did not", haystack, needle)
	}
}
