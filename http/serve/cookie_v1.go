/**
 * @time: 2019-08-22 22:31
 * @author: louis
 */
package serve

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", Cookie)
	http.HandleFunc("/2", Cookie2)
	http.HandleFunc("/read", ReadCookie)
	http.ListenAndServe(":8000", nil)
}

func Cookie(w http.ResponseWriter, r *http.Request) {
	ck := &http.Cookie{
		Name:   "myCookie",
		Value:  "hello",
		Path:   "/",
		Domain: "localhost",
		MaxAge: 12000,
	}

	http.SetCookie(w, ck)
	w.Write([]byte("<b>设置cookie成功.<b>\n"))
}

func ReadCookie(w http.ResponseWriter, r *http.Request)  {
	h:= r.Header["Cookie"]
	fmt.Fprintln(w,h)
}

func Cookie2(w http.ResponseWriter, r *http.Request) {
	ck := &http.Cookie{
		Name:   "myCookie",
		Value:  "helloworld",
		Path:   "/",
		Domain: "localhost",
		MaxAge: 120,
	}

	w.Header().Set("Set-Cookie",ck.String())

	ck2, _ := r.Cookie("myCookie")
	//cs := r.Cookies()
	fmt.Fprintln(w,ck2)
	//fmt.Fprintln(w,cs)
}