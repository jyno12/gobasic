package specification_test

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

type Food int

const (
	Noodle Food = iota
	Rice
)

type FakeFood int

var (
	SteamBread Food = 2
)

type Snack = int

const (
	IceCream Snack = iota
	Chip
)

func GetSundayFood() Food {
	rand.Seed(time.Now().UnixNano())
	return Food(rand.Intn(2))
}

func GetSundaySnack() Snack {
	rand.Seed(time.Now().Unix())
	return rand.Intn(2)
}

// SwitchStmt = ExprSwitchStmt | TypeSwitchStmt .

// Switch表达式
func TestExprSwitch(t *testing.T) {
	// switch表达式仅执行一次
	// 简单语句
	// weekDay:=time.Now().Weekday()
	rand.Seed(time.Now().UnixNano())
	weekDay := time.Weekday(rand.Intn(7))
	switch weekDay {
	// 执行顺序为从上到下，从左到右
	// 	case为常量
	case time.Saturday:
		t.Log("周六")
		t.Log("没有感觉就进入第二天")
		// fallthrough 连续到下一个case
		fallthrough
	case time.Sunday:
		// 没有fallthrough直接退出了
		t.Log("周日")
		// 表达式形式
		switch meal := GetSundayFood(); meal {
		// case 语句中如果没有定义类型，会隐式转换为switch中的类型
		case 2: // Rice， 默认为执行 Food(2) == meal，类型要允许执行比较操作
			t.Log("周日吃饭")
		case Noodle:
			t.Log("周日吃面")
		}

		// 分号后面没有内容，默认为true switch true{}
		// snack为临时变量且没有指定决定类型
		switch snack := GetSundaySnack(); {
		// 注意：这里没有分号，这种情况是不允许的。但是类型表达式 switch snackType:=someSnack.(type) {}允许
		// switch snack := GetSundaySnack() {
		case snack == IceCream:
			t.Log("零食冰激凌")
		case snack == Chip:
			t.Log("零食薯片")
		}

	// 	执行顺序为从左到右
	case time.Monday, time.Tuesday, time.Thursday, time.Wednesday, time.Friday:
		t.Log("周一到周五的工作日")
		// 省略条件，默认表达式为true
		switch {
		case time.Monday == weekDay:
			t.Log("周一")
		case time.Tuesday == weekDay:
			t.Log("周二")
		case time.Thursday == weekDay:
			t.Log("周三")
		case time.Wednesday == weekDay:
			t.Log("周四")
		case time.Friday == weekDay:
			t.Log("周五")
		}
	}
}

// switch 类型比较
func TestTypeSwitch(t *testing.T) {
	var bread interface{} = SteamBread
	// switch中的变量类型必须是interface type，即类型必须为interface{}
	switch bread.(type) {
	// 下面的类型必须都是不同的
	case Food:
		t.Log("是主食")
	case Snack:
		t.Log("是零食")
	case FakeFood:
		// 和Food相同都是int，但是不算重复类型
		t.Log("是个假的主食")
	}

	for _, thing := range []interface{}{
		interface{}(Chip),
		interface{}(Rice),
		interface{}(SteamBread),
	} {
		thingType := reflect.TypeOf(thing)
		// 这里的赋值类型
		switch thingValue := thing.(type) {
		// case条件还是类型
		case Food:
			t.Logf("bread的类型是%v，具体值：%v", thingType.Name(), thingValue)
		case Snack:
			t.Logf("bread的类型是%v，具体值：%v", thingType.Name(), thingValue)
		case FakeFood:
			t.Logf("bread的类型是%v，具体值：%v", thingType.Name(), thingValue)
		}
	}
}
