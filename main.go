//package main
//
//import "fmt"
//
//func main() {
//	var a, b float32
//	var sym string
//	println("Вам нужно ввести пример (2 числа и между ними знак действия) через пробел")
//	fmt.Scanf("%f %s %f", &a, &sym, &b)
//	switch sym {
//	case "+":
//		println(a + b)
//	case "-":
//		println(a - b)
//	case "*":
//		println(a * b)
//	case "/":
//		println(a / b)
//	case ">":
//		println(a < b)
//	case "<":
//		println(a < b)
//	case ">=":
//		println(a >= b)
//	case "<=":
//		println(a <= b)
//	case "==":
//		println(a == b)
//	case "!=":
//		println(a != b)
//	}
//}

//package main
//
//import (
//	"fmt"
//	"strconv"
//)
//
//func main() {
//	var a string
//	fmt.Scan(&a)
//	b, e := strconv.Atoi(a)
//	if e != nil {
//		panic("ПОШЕЛ НАФИГ")
//	}
//	if b%2 == 0 {
//		fmt.Println("Четное число")
//	} else {
//		fmt.Println("Нечетное число")
//	}
//}
//package main
//
//import "fmt"
//
//type error interface {
//	Error() string
//}
//
//func main() {
//	var a rune
//	b, e := fmt.Scan(&a)
//	if b == 0 {
//		panic("Вы ввели не число")
//	}
//}

//package main
//
//import (
//	"fmt"
//	"math"
//)
//
//func main() {
//	var a, b, c float64
//	println("Введите коэффициент a при x^2 в квадратном уравнении: ")
//	kod, err := fmt.Scanf("%f", &a)
//	if err != nil || kod == 0 {
//		println("Вы ввели неверное значение коэффициента, так что будет взято значение по умолчанию: 0")
//		a = 0
//	}
//	println("Введите коэффициент b при x в квадратном уравнении: ")
//	kod, err = fmt.Scanf("%f", &b)
//	if err != nil {
//		println("Вы ввели неверное значение коэффициента, так что будет взято значение по умолчанию: 0")
//		b = 0
//	}
//	println("Введите коэффициент c при x^0 в квадратном уравнении: ")
//	kod, err = fmt.Scanf("%f", &c)
//	if err != nil {
//		println("Вы ввели неверное значение коэффициента, так что будет взято значение по умолчанию: 0")
//		c = 0
//	}
//	D := ((b * b) - 4*a*c)
//	if a == 0 && b == 0 && c != 0 {
//		fmt.Errorf("Const != 0")
//	} else if a == 0 && b == 0 && c == 0 {
//		println("Корней нет, 0 = 0")
//	} else if a == 0 {
//		println("ЭТО НЕ КВАДРАТНОЕ УРАВНЕНИЕ")
//	} else if D < 0 {
//		println("Тут компексные корни будут, а мне лень их прописывать поэтому типо нет решения")
//	} else if D == 0 {
//		x := -b / (2 * a)
//		fmt.Printf("Уравнение будет иметь единственный корень равный: %f\n", x)
//	} else {
//		x1 := (-b + math.Sqrt(D)) / (2 * a)
//		x2 := (-b - math.Sqrt(D)) / (2 * a)
//		fmt.Printf("Уравнение имеет два действительных корня: %f: %f\n", x1, x2)
//	}
//}

//package main
//
//import "fmt"
//
//func main() {
//	students := make(map[string]int)
//	students["John"] = 1
//	students["Jane"] = 2
//	students["Paul"] = 3
//	students["Jack"] = 4
//	fmt.Println(students)
//}

//package main
//
//import "fmt"
//
//func main() {
//	var a float32
//	b, err := fmt.Scanf("%f", &a)
//	if err != nil || b == 0 {
//		fmt.Errorf("Неверный ввод числа! %s", &a)
//	} else {
//		if a < 0 {
//			println("Отрицательное число")
//			return
//		}
//		println("Положительное число или 0")
//	}
//}

//package main
//
//import "fmt"
//
//func main() {
//	var a, b int
//	println("Введите 2 числа через пробел")
//	c, e := fmt.Scanf("%d %d", &a, &b)
//	if e != nil || c != 2 {
//		fmt.Println("Не удалось определить числа")
//		return
//	} else {
//		if a > b {
//			fmt.Printf("%d > %d\n", a, b)
//		} else {
//			fmt.Printf("%d <= %d\n", a, b)
//		}
//	}
//}

//package main
//
//import "fmt"
//
//func main() {
//	var check bool
//	var a int
//	b, e := fmt.Scanf("%d", &a)
//	if e != nil || b == 0 {
//		println("Ошибка ввода")
//		return
//	}
//	check = a%3 == 0
//	println(check)
//}

//package main
//
//import "fmt"
//
//func main() {
//	var a int
//	println("Введите целое число: ")
//	b, e := fmt.Scanf("%d", &a)
//	if e != nil || b == 0 {
//		println("Ошибка ввода")
//		return
//	}
//	if a%3 == 0 && a%5 == 0 {
//		println("FizzBuzz")
//	} else if a%3 == 0 {
//		println(fmt.Println("Fizz"))
//	} else if a%5 == 0 {
//		println(fmt.Println("Buzz"))
//	} else {
//		println(a)
//	}
//}

//package main
//
//func main() {
//	technologies := []string{"go", "database", "sql", "bots", "stickers", "json"}
//	for _, technologyi := range technologies {
//		println(technologyi)
//	}
//}

// package main
//
// import "fmt"
//
//	func main() {
//		user_subscribe := []int{2, 3, 0, 5, 6, 7, 1, 20, 34, 256}
//		for _, user := range user_subscribe {
//			if user%2 == 0 {
//				fmt.Printf("%d ", user)
//			}
//		}
//
// }
//
// package main
//
//	func main() {
//		user_subscribe := []int{2, 3, 0, 5, 6, 7, 1, 20, 34, 256, 37}
//		for _, user := range user_subscribe {
//			if user%7 == 0 && user%2 != 0 {
//				println(user)
//			}
//		}
//	}
//package main
//
//func main() {
//	user_purchase := []int{65400, 23200, 871, 4800, 322}
//	sum := 0
//	for _, user := range user_purchase {
//		sum += user
//	}
//	println(sum / len(user_purchase))
//}

// package main
//
// import (
//
//	"fmt"
//	"strconv"
//
// )
//
//	func main() {
//		var a string
//		fmt.Scan(&a)
//		b, err := strconv.Atoi(a)
//		if err != nil {
//			fmt.Printf("Неверный ввод числа!")
//			return
//		} else {
//			if b < 0 {
//				println("Отрицательное число")
//				return
//			}
//			println("Положительное число или 0")
//		}
//	}
//package main
//
//func main() {
//	user_purchase := []int{65400, 23200, 871, 4800, 322}
//	var sum int = 0
//	for _, user := range user_purchase {
//		sum += user
//	}
//	println(float64(sum / len(user_purchase)))
//}

//package main
//
//import "fmt"
//
//func greet(name string) {
//	fmt.Printf("Привет, %s!\n", name)
//}
//
//func main() {
//	var a string
//	println("Введите ваше имя (ну или че нибудь): ")
//	fmt.Scanf("%s", &a)
//	greet(a)
//}

//package main
//
//import (
//	"fmt"
//	"strconv"
//)
//
//func add(first, second int) int {
//	return first + second
//}
//
//func main() {
//	var a, b string
//	println("Введите 2 числа через пробел: ")
//	fmt.Scanf("%s %s", &a, &b)
//	c, err := strconv.Atoi(a)
//	d, err1 := strconv.Atoi(b)
//	if err1 != nil || err != nil {
//		println("Неверный ввод!")
//		return
//	}
//	println(add(c, d))
//}

//package main
//
//import (
//	"fmt"
//	"strconv"
//)
//
//func isEven(n int) bool {
//	return n%2 == 0
//}
//
//func main() {
//	var a string
//	println("Введите число: ")
//	fmt.Scanln(&a)
//	b, err := strconv.Atoi(a)
//	if err != nil || b == 0 {
//		println("Ошибка ввода")
//		return
//	}
//	println(isEven(b))
//}

//package main
//
//import (
//	"fmt"
//	"strconv"
//)
//
//func MAX(x int, y int) int {
//	if x > y {
//		return x
//	}
//	return y
//}
//
//func main() {
//	var x, y string
//	println("Введите 2 числа через пробел")
//	fmt.Scanf("%s %s", &x, &y)
//	a, err := strconv.Atoi(x)
//	b, err1 := strconv.Atoi(y)
//	if err != nil || err1 != nil {
//		println("Ошибка ввода!")
//		return
//	}
//	println(MAX(a, b))
//}

//	func check_for_digit(str string) bool {
//		m, _ := regexp.MatchString("^[0-9]+$", str)
//		return m
//	}
//package main
//
//import (
//	"fmt"
//	"strconv"
//)
//
//func main() {
//	var a, b string
//	println("Введите 2 числа через пробел!")
//	fmt.Scanf("%s %s", &a, &b)
//	c, err := strconv.ParseFloat(a, 64)

//	d, err1 := strconv.ParseFloat(b, 64)
//	if err != nil || err1 != nil {
//		println("Ошибка ввода!")
//		return
//	}
//	fmt.Printf("%.2f", c+d)
//}

//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func main() {
//	for i := 10; i > 1; i-- {
//		fmt.Println(i)
//		time.Sleep(1 * time.Second)
//	}
//}

//package main
//
//import (
//	"fmt"
//	"strconv"
//)
//
//func Factorial(n int) int {
//	if n < 2 {
//		return 1
//	}
//	sum := 1
//	for i := 1; i < n+1; i++ {
//		sum *= i
//	}
//	return sum
//}
//
//func main() {
//	var a string
//	println("Введите число: ")
//	fmt.Scanf("%s", &a)
//	b, err := strconv.Atoi(a)
//	if err != nil {
//		println("Ошибка ввода!")
//		return
//	}
//	println(Factorial(b))
//}

//package main
//
//import (
//	"fmt"
//	"strconv"
//)
//
//func main() {
//	var a string
//	b := 0
//	c := 1
//	println("Введите номер N чисел Фибоначчи:")
//	fmt.Scan(&a)
//	n, e := strconv.Atoi(a)
//	if e != nil {
//		println("Ошибка ввода")
//		return
//	}
//	for i := 0; i < n+1; i++ {
//		print(b, " ")
//		b, c = c, b+c
//	}
//}

//package main
//
//import (
//	"fmt"
//	"math/rand"
//	"strconv"
//)
//
//func main() {
//	a := rand.Intn(100000)
//	println("Компьютер загадал рандомное число от 0 до 100000. Введите ваше число: ")
//	var b string
//	niznaya := 0
//	verkhaya := 100000
//	for {
//		fmt.Scanf("%s", &b)
//		c, err := strconv.Atoi(b)
//		if err != nil || c > 100000 || c < 0 {
//			println("Ошибка ввода. Либо слишком большое число либо оно меньше нуля")
//			return
//		}
//		if c < niznaya || c > verkhaya {
//			println("Вы уже вводили число по данную гриницу, попробуйте снова!")
//			continue
//		}
//
//		if c > a {
//			println("Загаданное число меньше вашего")
//			verkhaya = c
//		} else if c < a {
//			println("Загаданное число больше вашего")
//			niznaya = c
//		} else {
//			println("ПОЗДРАВЛЯЮ, ВЫ ОТГАДАЛИ ЧИСЛО!")
//			break
//		}
//	}
//}

//package main
//
//import (
//	"fmt"
//	"strconv"
//)
//
//func main() {
//	xods := []string{"X", "O"}
//	var a, b string
//	count := 0
//	Pole := [3][3]string{{":", ":", ":"}, {":", ":", ":"}, {":", ":", ":"}}
//	println("Добро пожаловать в крестики нолики для 2 игроков!")
//	println("Введите номер строки и столбца через пробел от 1 до 3, в которые хотите разместить ход (Первые ходят крестики)")
//	println("P.S. Ход 1 1 - означает левую верхнюю клетку, а 1 2 - верхнюю среднюю клетку")
//	for ; count < 9; count++ {
//		fmt.Scanf("%s %s", &a, &b)
//		c, err := strconv.Atoi(a)
//		if err != nil || c > 3 || c < 1 {
//			println("Ошибка, неверный ввод!")
//			return
//		}
//		d, err := strconv.Atoi(b)
//		if err != nil || d > 3 || d < 1 {
//			println("Ошибка, неверный ввод!")
//			return
//		}
//		if Pole[c-1][d-1] == ":" {
//			Pole[c-1][d-1] = xods[count%2]
//		} else {
//			println("Поле уже занято попробуйте снова")
//			count--
//			continue
//		}
//		for i := 0; i < 3; i++ {
//			for j := 0; j < 3; j++ {
//				fmt.Print(Pole[i][j])
//			}
//			fmt.Println()
//		}
//		if count > 3 {
//			if (Pole[c-1][d-1] == Pole[c-1][d%3]) && (Pole[c-1][d-1] == Pole[c-1][(d+1)%3]) {
//				println("Победа по горизонтали - ", xods[count%2])
//				break
//			} else if (Pole[c-1][d-1] == Pole[c%3][d-1]) && (Pole[c-1][d-1] == Pole[(c+1)%3][d-1]) {
//				println("Победа по вертикали - ", xods[count%2])
//				break
//			} else if Pole[1][1] != ":" {
//				if (Pole[0][0] == Pole[1][1]) && (Pole[1][1] == Pole[2][2]) {
//					println("Победа по главной диагонали - ", xods[count%2])
//					break
//				} else if (Pole[2][0] == Pole[1][1]) && (Pole[1][1] == Pole[0][2]) {
//					println("Победа по побочной диагонали - ", xods[count%2])
//					break
//				}
//			}
//		}
//	}
//	if count == 9 {
//		println("Ничья!")
//		return
//	}
//}
//
//package main
//
//func quick_sort(a []int) []int {
//	if len(a) < 2 {
//		return a
//	}
//	start := a[0] // Берем 0 элемент как опорный для быстрой сортировки
//	var left []int
//	var right []int
//	for i := 1; i < len(a); i++ {
//		if start > a[i] {
//			left = append(left, a[i])
//		} else {
//			right = append(right, a[i])
//		}
//	}
//	left = append(left, start)
//	left = quick_sort(left)
//	right = quick_sort(right)
//	for i := 0; i < len(right); i++ {
//		left = append(left, right[i])
//	}
//	return left
//}
//
//func main() {
//	a := []int{1, 2, 3, 78, 22, 12, 5, 7, -3, 0, 19, 77}
//	a = quick_sort(a)
//	for i := 0; i < len(a); i++ {
//		println(a[i])
//	}
//}

//package main
//
//import "fmt"
//
//func main() {
//	println("Лев, введи строку, а прога выведет тебе её в обратном порядке: ")
//	var line string
//	fmt.Scanln(&line)
//	rn := []rune(line)
//	for i := 0; i < (len(rn)/2 + (len(rn) % 2)); i++ {
//		rn[i], rn[len(rn)-1-i] = rn[len(rn)-1-i], rn[i]
//	}
//	str := string(rn)
//	println(str)
//}

//
//package main
//
//import "fmt"
//
//func main() {
//	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
//	for i := 0; i < len(a)/2+(len(a)%2); i++ {
//		a[i], a[len(a)-i-1] = a[len(a)-i-1], a[i]
//	}
//	for j := 0; j < len(a); j++ {
//		fmt.Print(a[j], " ")
//
//package main
