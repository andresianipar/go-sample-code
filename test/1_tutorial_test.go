package gosamplecode

import (
	"net/http"
	"net/http/httptest"
	"testing"

	library "github.com/andresianipar/go-sample-code/internal"
	tutorial "github.com/andresianipar/go-sample-code/pkg/1_tutorial"
)

// 1 tutorial > command-line arguments
func TestF1_F2_F3(t *testing.T) {
	// Test F1 function
	f1 := tutorial.F1()

	if "" == f1 {
		t.Errorf("Should not be an empty string")
	}

	// Test F2 function
	f2 := tutorial.F2()

	if "Apple, Pear, Plum, Tomato, Peach" != f2 {
		t.Errorf("Should be \"Apple, Pear, Plum, Tomato, Peach\"")
	}

	// Test F3 function
	f3 := tutorial.F3()

	if "Apple, Pear, Plum, Tomato, Peach" != f3 {
		t.Errorf("Should be \"Apple, Pear, Plum, Tomato, Peach\"")
	}
}

// 1 tutorial > fetching a url
func TestF7(t *testing.T) {
	// Test F7 function
	f7 := tutorial.F7()

	if nil == f7 {
		t.Errorf("Should not be empty")
	}
}

// 1 tutorial > fetching urls concurrently
func TestF8(t *testing.T) {
	// Test F8 function
	f8 := tutorial.F8()

	if nil == f8 {
		t.Errorf("Should not be empty")
	}
}

// 1 tutorial > a web server
func TestF9_F10_F11(t *testing.T) {
	// Test F9 function
	req, err := http.NewRequest("GET", "/", nil)

	if nil != err {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler1 := http.HandlerFunc(library.Handler1)

	handler1.ServeHTTP(rr, req)

	if http.StatusOK != rr.Code {
		t.Errorf("Handler1 returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
	}

	path := "\"/\"\n"

	if "URL.Path = "+path != rr.Body.String() {
		t.Errorf("Handler1 returned unexpected body: got %v want %v", rr.Body.String(), "URL.Path = "+path)
	}

	// Test F10 function
	req1, err := http.NewRequest("GET", "/", nil)
	req2, err := http.NewRequest("GET", "/count", nil)

	if nil != err {
		t.Fatal(err)
	}

	rr1 := httptest.NewRecorder()
	rr2 := httptest.NewRecorder()
	handler2 := http.HandlerFunc(library.Handler2)
	counter := http.HandlerFunc(library.Counter)

	handler2.ServeHTTP(rr1, req1)
	counter.ServeHTTP(rr2, req2)

	if http.StatusOK != rr1.Code {
		t.Errorf("Handler2 returned wrong status code: got %v want %v", rr1.Code, http.StatusOK)
	}

	if http.StatusOK != rr2.Code {
		t.Errorf("Cointer returned wrong status code: got %v want %v", rr2.Code, http.StatusOK)
	}

	path = "\"/\"\n"

	if "URL.Path = "+path != rr1.Body.String() {
		t.Errorf("Handler1 returned unexpected body: got %v want %v", rr1.Body.String(), "URL.Path = "+path)
	}

	if "Count: 1\n" != rr2.Body.String() {
		t.Errorf("Counter returned unexpected body: got %v want %v", rr2.Body.String(), "Count: 1\n")
	}

	// Test F11 function
	req, err = http.NewRequest("GET", "/", nil)

	if nil != err {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	handler3 := http.HandlerFunc(library.Handler3)

	handler3.ServeHTTP(rr, req)

	if http.StatusOK != rr.Code {
		t.Errorf("Handler3 returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
	}

	if "" == rr.Body.String() {
		t.Errorf("Should not be an empty string")
	}
}

// 1 tutorial > loose ends
func TestF12_F13(t *testing.T) {
	// Test F12 function
	heads, tails := tutorial.F12("heads")

	if 1 != heads || 0 != tails {
		t.Errorf("heads should be 1 and tails should be 0")
	}

	// Test F13 function
	result := tutorial.F13(5)

	if 1 != result {
		t.Errorf("result should be 1")
	}
}

func BenchmarkF2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tutorial.F2()
	}
}

func BenchmarkF3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tutorial.F3()
	}
}
