package main

import (
	"fmt"
	"strings"
)

type User struct {
	Name string
}

func main() {}

// func main() {
// 	// fmt.Println(os.Args[1])
// 	// defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()

// 	// fmt.Println(rand.Int())
// 	// stackUserAlloc()
// 	// heapUserAlloc()
// 	// user := new(User)
// 	// mw := new(MyWorker)
// 	// mw := MyWorker{}
// 	// DoWorkInt(mw, User{})

// 	// s1 := make([]int, 0, 5)
// 	// s2 := make([]int, 0, 5)

// 	// for i := 0; i < 5; i++ {
// 	// 	s1 = append(s1, i)
// 	// }

// 	// for _, v := range s2 {
// 	// 	fmt.Println(v)
// 	// }

// 	// sort.Slice(bms, func(i, j int) bool {
// 	// 	return bms[i] > bms[j]
// 	// })

// 	// showCodeIdx := 2

// 	// stuntItems := []string{
// 	// 	"jfk-the-conspiracy-continues",
// 	// 	"brad-meltzers-greatest-conspiracies-of-all-time",
// 	// 	"the-secrets-of-area-51",
// 	// 	"monsters-across-america",
// 	// 	"fire-over-the-atlantic-the-mystery-of-twa-flight-800",
// 	// 	"riddle-the-search-for-james-r-hoffa",
// 	// 	"scandalous-chappaquiddick"}

// 	// left := stuntItems[:showCodeIdx]
// 	// right := stuntItems[showCodeIdx:]

// 	// stuntItems = append(right, left...)

// 	// items := append([]string{}, stuntItems[showCodeIdx:]...)
// 	// items = append(items, stuntItems[:showCodeIdx]...)

// 	// retrieve from list the following items
// 	// itemsLen := len(stuntItems[showCodeIdx+1:])

// 	// items := make([]string, itemsLen)
// 	// copy(items, stuntItems[showCodeIdx+1:])

// 	// bookmarks := []string{"1", "2", "3", "4", "5", "6", "7"}
// 	// bmIds := make([]string, 0, len(bookmarks))
// 	// for _, v := range bookmarks {
// 	// 	bmIds = append(bmIds, "")
// 	// 	fmt.Printf(v)
// 	// }

// 	// id := "testSuitetest1458839464305" + "sleepy-hollow_03_15"
// 	// dst := make([]byte, hex.EncodedLen(len(id)))
// 	// hex.Encode(dst, []byte(id))

// 	// fmt.Println(string(dst))

// 	// fmt.Println(rand.Intn(1000))

// 	var book Book

// 	if book.id == 0 {
// 		fmt.Println("empty")
// 	} else {
// 		fmt.Println("not empty")
// 	}

// }

// type Book struct {
// 	id int
// }

// // escapes to heap
// func heapAlloc() {
// 	_, err := getDataHeap()

// 	if err != nil {

// 		fmt.Println(err)
// 	}
// }

// func getDataHeap() ([]byte, error) {
// 	data := make([]byte, 1000)

// 	for i := range data {
// 		data[i] = 'a'
// 	}

// 	return data, nil
// }

// func heapUserAlloc() {
// 	_, err := getUserHeap()

// 	if err != nil {

// 		fmt.Println(err)
// 	}
// }

// func getUserHeap() ([]User, error) {
// 	data := make([]User, 1000)

// 	for i := range data {
// 		data[i] = User{Name: "scartezini"}
// 	}

// 	return data, nil
// }

// // does not escape
// func stackAlloc() {
// 	data := make([]byte, 1000)
// 	getDataStack(data)
// }

// func getDataStack(data []byte) {
// 	for i := range data {
// 		data[i] = 'a'
// 	}
// }

// func stackUserAlloc() {
// 	data := make([]User, 1000)
// 	getUserStack(data)
// }

// func getUserStack(data []User) {
// 	for i := range data {
// 		data[i] = User{Name: "scartezini"}
// 	}
// }

// type Worker interface {
// 	Work(user User)
// }

// type MyWorker struct {
// }

// func (mw MyWorker) Work(user User) {

// }

// func DoWorkInt(worker Worker, user User) {
// 	worker.Work(user)
// }

// type Register struct {
// 	Password     string
// 	PasswordHash string
// }

// func (req *Register) HashPasswd() error {
// 	// Using cost "0" to cast as default cost "10".

// 	bytes := make([]byte, 0, len(req.Password))

// 	bytes = append(bytes, req.Password...)
// 	// bytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), 0)
// 	// if err != nil {
// 	// 	return err
// 	// }

// 	req.PasswordHash = *(*string)(unsafe.Pointer(&bytes))
// 	// req.Password = string(bytes)
// 	return nil
// }

func FormatSprintf() string {
	return fmt.Sprintf("%s:%s", "/oauth2/default/v1/logout", "id_token_hint=8427347283472348325828357")
}

func FormatPlusSign() string {
	return "/oauth2/default/v1/logout" + ":" + "id_token_hint=8427347283472348325828357"
}

func FormatBuilder() string {
	var b strings.Builder
	b.WriteString("/oauth2/default/v1/logout")
	b.WriteString(":")
	b.WriteString("id_token_hint=8427347283472348325828357")
	return b.String()
}
