/*task 4.1*/
package main

import "fmt"

func main() {
	res := helloFunc()
	fmt.Println(res)

}

func helloFunc() string {
	return fmt.Sprint("Hello Go!")
}

/*task 4.2
package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hello Go!")
	}()

}
*/

/*task 4.3
package main

import "fmt"

func main() {
	var f func() string

	f = func() string {
		return fmt.Sprint("Hello Go!")
	}

	hello(f)

}

func hello(v func() string) {
	res := v()
	fmt.Println(res)
}*/

/*task 4.4
package main

import "fmt"

func hello() func() {
	return func() {
		fmt.Println("Hello Go!")
	}
}

func main() {
	hello()()
}
 */

/*task 4.5

package main

import "fmt"

func main() {
	defer fmt.Println("завершение работы!")
	fmt.Println(helloFunc())
}

func helloFunc() string {
	return fmt.Sprint("Hello Go!")
}*/

/*task 4.6
package main

import "fmt"

func fib(a, b, c int) {
	if c == 0 {
		return
	}
	fmt.Println(a)
	fib(b, a+b, c-1)
}

func main() {
	fib(0, 1, 23)
}
*/
