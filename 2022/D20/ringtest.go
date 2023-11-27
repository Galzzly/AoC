package main

import (
	"container/ring"
	"fmt"
	"time"

	"github.com/Galzzly/AoC/utils"
)

type Pos struct {
	i, v int
}

type Rings map[Pos]*ring.Ring

func main() {
	start := time.Now()
	nums := utils.ReadIntsByLine("input.txt")
	t1 := time.Now()
	fmt.Println("Part 1:", sol(nums, 1, 1), "Took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 1:", sol(nums, 10, 811589153), "Took:", time.Since(t2))
	fmt.Println("Total time:", time.Since(start))
}

func sol(nums []int, round, m int) (res int) {
	for i := range nums {
		nums[i] *= m
	}
	rings, zero := makeRing(nums)
	length := len(nums) - 1
	half := length / 2
	for i := 0; i < round; i++ {
		for i, v := range nums {
			r := rings[Pos{i, v}].Prev()
			T := r.Unlink(1)
			if v > half || v < -half {
				v %= length
				switch {
				case v > half:
					v -= length
				case v < -half:
					v += length
				}
			}
			r.Move(v).Link(T)
		}
	}
	r := rings[zero]
	t := 1000 % len(nums)
	for i := 1; i <= 3; i++ {
		r = r.Move(t)
		// fmt.Println(i, r.Value.(int))
		res += r.Value.(int)
	}
	return
}

func makeRing(nums []int) (map[Pos]*ring.Ring, Pos) {
	rings := make(map[Pos]*ring.Ring)
	R := ring.New(len(nums))
	zero := Pos{v: 0}
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zero.i = i
		}
		n := nums[i]
		rings[Pos{i, n}] = R
		R.Value = n
		R = R.Next()
	}
	return rings, zero
}
