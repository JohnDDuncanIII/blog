package comment

import (
	"fmt"
	"strconv"
	"html/template"
	"time"
	//"net/mail"
	//"faces/parse"
)

func main() {
	entry_num := 1
	comment_num := 1


	subject := "Test"
	ip := "127.0.0.1"
	name := "John"
	email := "duncjo01@gettysburg.edu"
	x_face := ""
	face := ""
	homepage := "http://cs.gettysburg.edu/~duncjo01/"
	content := ":)"
	epoch := strconv.Itoa(int32(time.Now().Unix()))
	deeper := ""

	toInsert := subject+"¦"+ip+"¦"+email+"¦"+homepage+"¦"+epoch+"¦"+content+"¦"+face+"¦"+x_face+"¦"+deeper
	//sub2¦::1¦bwk@research.att.com¦https://www.bell-labs.com/usr/bwk/www/¦1489799000¦:)¦¦¦..
}

/*func insert(parentIndex int, toInsert string) []string {
	var s []string
	i := 1
	s = append(s, "0")
	s = append(s, "5")
	s = append(s, "10")
	s = s[0 : len(s)+1]
	copy(s[i+1:], s[i:])
	s[i] = toInsert
	fmt.Println(s)
	return s
}*/

/*func is_valid_email() string {
	e, err := mail.ParseAddress("alice@example.com")
	if err == nil {
		return e.Name+e.Address
	}
	return ""
}*/
