package generator

import (
	"math/rand"
	"time"

	"fmt"
)

func Generator() string {
	a_z := "abcdefghijklmnopqrstuvwxyz"
	A_Z := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	nb := "0123456789"
	var slice []string
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 3; i++ {
		randInt := rand.Intn(len(a_z))
		slice = append(slice, string(a_z[randInt]))
	}
	for i := 0; i < 3; i++ {
		randInt := rand.Intn(len(A_Z))
		slice = append(slice, string(A_Z[randInt]))
	}
	for i := 0; i < 3; i++ {
		randInt := rand.Intn(len(nb))
		slice = append(slice, string(nb[randInt]))
	}
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})

	fmt.Println(slice)

	var result string
	for _, v := range slice {
		result = result + v
	}
	fmt.Println(result)
	return result
}
