package main

import (
	v100 "github.com/hentan/task13/100"
	v110 "github.com/hentan/task13/110"
	v200 "github.com/hentan/task13/200"
	v210 "github.com/hentan/task13/210"
)

func main() {
	v100.Do("patients.txt", "res.txt")
	v110.Do("patients.txt", "res2.txt")
	v200.Do("patients.txt", "res3.txt")
	v210.Do("patients.txt", "res4.txt")
}
