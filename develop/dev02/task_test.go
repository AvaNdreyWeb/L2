package main

import "testing"

func TestUnpackString1(t *testing.T) {

	var got, want, ps string

	ps = "a4bc2d5e"
	want = "aaaabccddddde"
	got = UnpackString(ps)

	if got != want {
		t.Errorf("UnpackString(%q) == %q, want %q", ps, got, want)
	}
}

func TestUnpackString2(t *testing.T) {

	var got, want, ps string

	ps = "abcd"
	want = "abcd"
	got = UnpackString(ps)

	if got != want {
		t.Errorf("UnpackString(%q) == %q, want %q", ps, got, want)
	}
}

func TestUnpackString3(t *testing.T) {

	var got, want, ps string

	ps = "45"
	want = ""
	got = UnpackString(ps)

	if got != want {
		t.Errorf("UnpackString(%q) == %q, want %q", ps, got, want)
	}
}

func TestUnpackString004(t *testing.T) {

	var got, want, ps string

	ps = ""
	want = ""
	got = UnpackString(ps)

	if got != want {
		t.Errorf("UnpackString(%q) == %q, want %q", ps, got, want)
	}
}
