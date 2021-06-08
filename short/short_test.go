package short_t_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestShort(t *testing.T) {

	url := "https://v2.alapi.cn/api/url/query"

	payload := strings.NewReader("url=http://t.cn/A6tYYKyI")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("token", "v2WDehKoih2NfVcU")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

func TestHitokoto(t *testing.T) {
	url := "https://v2.alapi.cn/api/hitokoto"

	payload := strings.NewReader("type=a&format=json")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("token", "v2WDehKoih2NfVcU")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

func TestDeepEqual(t *testing.T) {
	f := reflect.DeepEqual(struct {
		A int
	}{A: 2}, struct{ B int }{B: 2})
	t.Logf("%v", f)

	type PP struct {
		P int
	}

	type D struct {
		V int
		P PP
	}

	var a, b D
	a = D{V: 1, P: PP{P: 1}}
	b = D{V: 1, P: PP{P: 1}}
	t.Logf("%v", a == b)
}

func TestDefer(t *testing.T) {
	outA()
}

func outA() {
	defer func() {
		fmt.Println("out_a")
		defer outB()
	}()
}

func outB() {
	fmt.Println("out_b")
}

type D struct{}

type F interface {
}

var (
	a map[string]interface{}
	d map[D]interface{}

	f map[F]interface{}
)

func TestMap(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	a1 := s[1:2]
	a1 = append(a1, 1)
	fmt.Println(len(a1), cap(a1))

	a2 := s[1:2:2]
	a2 = append(a2, 1)
	fmt.Println(len(a2), cap(a2))
}
