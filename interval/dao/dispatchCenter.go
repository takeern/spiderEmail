package dao

// import (
// 	"regexp"
// 	"strings"
// 	"os"
// 	"fmt"

// 	// "spider/interval/modal"
// )

// func CreateDispatch(url string)  {
// 	re := regexp.MustCompile(`(http|https):\/\/?([^/]*)`)
// 	dbName := strings.Replace(string(re.Find([]byte(url))), ".", "", -1) + "/"
// 	fmt.Println(dbName)
// 	mb := NewDb(dbName)
// 	emailModals, err :=  mb.SelectData(10)
// 	if err != nil {
// 		fmt.Println(err)
// 		Log.Error("select emails error", err)
// 		os.Exit(3)
// 	}
// 	fmt.Println(mb)
// 	fmt.Println(len(emailModals))
// 	mb.UpdateStatus("128d9406@st.kumamoto", true)
// }

