//go:build !debug

package errors

import (
	"io"
	"os"
	"os/exec"
	"testing"
)

func TestDebug(t *testing.T) {
	// Test with -tags debug to run the tests in debug_test.go
	cmd := exec.Command("go", "test", "-tags", "debug")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		t.Fatalf("external go test failed: %v", err)
	}
}

func TestSeparator(t *testing.T) {
	defer func(prev string) {
		Separator = prev
	}(Separator)
	Separator = ":: "

	// Same pattern as above.
	user := UserName("joe@blow.com")

	// Single error. No user is set, so we will have a zero-length field inside.
	e1 := E(Op("Get"), IO, "network unreachable")

	// Nested error.
	e2 := E(Op("Read"), user, Other, e1)

	want := "Read, user joe@blow.com: I/O error:: Get: network unreachable"
	if errorAsString(e2) != want {
		t.Errorf("expected %q; got %q", want, e2)
	}
}

func TestDoesNotChangePreviousError(t *testing.T) {
	err := E(Permission)
	err2 := E(Op("I will NOT modify err"), err)

	expected := "I will NOT modify err: permission denied"
	if errorAsString(err2) != expected {
		t.Fatalf("Expected %q, got %q", expected, err2)
	}
	kind := err.(*Error).Kind
	if kind != Permission {
		t.Fatalf("Expected kind %v, got %v", Permission, kind)
	}
}

func TestNoArgs(t *testing.T) {
	defer func() {
		err := recover()
		if err == nil {
			t.Fatal("E() did not panic")
		}
	}()
	_ = E()
}

type matchTest struct {
	err1, err2 error
	matched    bool
}

const (
	john = UserName("john@doe.io")
	jane = UserName("jane@doe.io")
)

const (
	op  = Op("Op")
	op1 = Op("Op1")
	op2 = Op("Op2")
)

var matchTests = []matchTest{
	// Errors not of type *Error fail outright.
	{nil, nil, false},
	{io.EOF, io.EOF, false},
	{E(io.EOF), io.EOF, false},
	{io.EOF, E(io.EOF), false},
	// Success. We can drop fields from the first argument and still match.
	{E(io.EOF), E(io.EOF), true},
	{E(op, Invalid, io.EOF, jane), E(op, Invalid, io.EOF, jane), true},
	{E(op, Invalid, io.EOF, jane), E(op, Invalid, io.EOF, jane), true},
	{E(op, Invalid, io.EOF), E(op, Invalid, io.EOF, jane), true},
	{E(op, Invalid), E(op, Invalid, io.EOF, jane), true},
	{E(op), E(op, Invalid, io.EOF, jane), true},
	// Failure.
	{E(io.EOF), E(io.ErrClosedPipe), false},
	{E(op1), E(op2), false},
	{E(Invalid), E(Permission), false},
	{E(jane), E(john), false},
	{E(op, Invalid, io.EOF, jane), E(op, Invalid, io.EOF, john), false},
	{E(Str("something")), E(jane), false}, // Test nil error on rhs.
	// Nested *Errors.
	{E(op1, E(john)), E(op1, E(op2, john)), true},
	{E(op1, E(john)), E(op1, E(op2, jane)), false},
}

func TestMatch(t *testing.T) {
	for _, test := range matchTests {
		matched := Match(test.err1, test.err2)
		if matched != test.matched {
			t.Errorf("Match(%q, %q)=%t; want %t", test.err1, test.err2, matched, test.matched)
		}
	}
}

type kindTest struct {
	err  error
	kind Kind
	want bool
}

var kindTests = []kindTest{
	// Non-Error errors.
	{nil, NotExist, false},
	{Str("not an *Error"), NotExist, false},

	// Basic comparisons.
	{E(NotExist), NotExist, true},
	{E(Exist), NotExist, false},
	{E("no kind"), NotExist, false},
	{E("no kind"), Other, false},

	// Nested *Error values.
	{E("Nesting", E(NotExist)), NotExist, true},
	{E("Nesting", E(Exist)), NotExist, false},
	{E("Nesting", E("no kind")), NotExist, false},
	{E("Nesting", E("no kind")), Other, false},
}

func TestKind(t *testing.T) {
	for _, test := range kindTests {
		got := Is(test.kind, test.err)
		if got != test.want {
			t.Errorf("Is(%q, %q)=%t; want %t", test.kind, test.err, got, test.want)
		}
	}
}

func errorAsString(err error) string {
	if e, ok := err.(*Error); ok {
		e2 := *e
		return e2.Error()
	}
	return err.Error()
}
