package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// GENERATE A RANDOM NUMBER
func generateNumber(num int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(num)
}

// GENERATE LOTTO NUMBERS
func lotteryNumbers(amount, maxNum int) []int {
	xi := make([]int, amount)

	// GENERATE A RANDOM VALUE FOR EACH INDEX IN x
	for i := 0; i < amount; i++ {
		tempNum := generateNumber(maxNum)

		// CHECK UNIQUENESS OF VALUE IN x
		for _, v := range xi {

			// KEEP GENERATING A NEW VALUE IN x[i] UNTIL VALUE IS UNIQUE AND NOT NULL IN x
			for tempNum == v || tempNum == 0 {
				tempNum = generateNumber(maxNum)
			}
		}
		xi[i] = tempNum
	}
	sort.Ints(xi)
	return xi
}

func main() {
	lottery := flag.String("l", "6of49", "Kind of Lottery: 6of49 or euro")
	superNum := flag.Bool("s", false, "Genereate SuperZahl for 6/49 Lottery")
	lotteryFields := flag.Int("n", 1, "Amount of genereated lottery fileds")
	flag.Parse()

	switch *lottery {
	// LOTTERY 6of49 (6 NUMBERS OUT OF 49)
	case "6of49":
		for i := 0; i < *lotteryFields; i++ {
			fmt.Println(lotteryNumbers(6, 49))
		}
		if *superNum {
			fmt.Println(lotteryNumbers(1, 10))
		}
	// EUROLOTTO 5 OUT OF 50 + 2 OUT OF 10
	case "euro":
		for i := 0; i < *lotteryFields; i++ {
			fmt.Println(lotteryNumbers(5, 50), lotteryNumbers(2, 10))
		}
	default:
		flag.Usage()
	}
}
