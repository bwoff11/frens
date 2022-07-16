package test

type APITest struct {
	Description  string      // description of the test case
	Method       string      // HTTP method
	Route        string      // route path to test
	ReqBody      interface{} // request body
	ExpectedCode int         // expected HTTP status code
	// expectedBody string // expected response body
}
