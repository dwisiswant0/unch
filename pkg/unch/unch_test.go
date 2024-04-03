package unch_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/dwisiswant0/unch/pkg/unch"
)

func TestEncode(t *testing.T) {
	t.Parallel()

	message := "s3cr3t"
	plaintext := "foo bar"

	encoded, err := unch.Encode(message, plaintext)
	if err != nil {
		t.Fatal(err)
	}

	if len(encoded) == len(plaintext) {
		t.Fatalf("'%s' is identical to '%s'", encoded, plaintext)
	}
}

func TestDecode(t *testing.T) {
	t.Parallel()

	message := "s3cr3t"
	plaintext := "foo bar"

	encoded, err := unch.Encode(message, plaintext)
	if err != nil {
		t.Fatal(err)
	}

	decoded, err := unch.Decode(encoded)
	if err != nil {
		t.Fatal(err)
	}

	if len(decoded) == len(encoded) {
		t.Fatalf("'%s' is identical to '%s'", encoded, plaintext)
	}
}

func BenchmarkEncode(b *testing.B) {
	var message strings.Builder
	plaintext := "foo bar"

	for _, n := range []int{1, 10, 100, 1000, 10000} {
		for i := range n {
			for range i {
				message.WriteString("a")
			}
		}

		b.Run(fmt.Sprintf("%d-char", n), func(b *testing.B) {
			b.ResetTimer()
			b.ReportAllocs()

			for range b.N {
				_, _ = unch.Encode(message.String(), plaintext)
			}
		})

		message.Reset()
	}
}

func BenchmarkDecode(b *testing.B) {
	message := "a"
	plaintext := "foo bar"

	encoded, err := unch.Encode(message, plaintext)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for range b.N {
		_, _ = unch.Decode(encoded)
	}
}

func FuzzEncode(f *testing.F) {
	f.Fuzz(func(t *testing.T, _, a string) {
		_, err := unch.Encode(gofakeit.Phrase(), a)
		if err != nil {
			t.Fatal(err)
		}
	})
}

// func FuzzDecode(f *testing.F) {
// 	f.Fuzz(func(t *testing.T, _, a string) {
// 		_, _ = unch.Decode(a)
// 	})
// }

func ExampleEncode() {
	message := "s3cr3t"
	plaintext := "foo bar"

	encoded, err := unch.Encode(message, plaintext)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(os.Stderr, "plaintext: %s (%d)\n", plaintext, len(plaintext))
	fmt.Fprintf(os.Stderr, "encoded: %s (%d)\n", encoded, len(encoded))

	fmt.Println("plaintext == encoded:", plaintext == encoded)
	// Output: plaintext == encoded: false
}
