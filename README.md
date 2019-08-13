xbinder
=======

Bind your struct from http query params

## Install
```
go get -u github.com/slavabobik/xbinder
``` 

## Example 
```go
type Pet struct {
	Name    string
	Age     int
	Hobbies []string
}

func MyHandler(w http.ResponseWriter, r *http.Request) {
	//for e.g. you have a query http://example.com?name=snowball&age=2&hobbies=eat,sleep,repeat
	values := r.URL.Query()

	var pet Pet
	err := xbinder.FromQuery(&pet, values)
	if err != nil {
		// Handle error
	}
}
```

Supported field types:
* bool
* int, int8,  int16, int32, int64
* string
* int, string slices

Unsupported field types:
* float
* pointers
* uint

