import os
import arrays
import time

fn main() {
	start := time.now()
	file := os.read_file(os.args[1]) or {
		return
	}

	lines := file.split('\n\n')
	elves := get_elves(lines)!
	t1 := time.now()
	println(elves[0])
	println(time.since(t1))
	t2 := time.now()
	println(arrays.sum(elves[0..3])!)
	println(time.since(t2))
	println(time.since(start))
}

fn get_elves(lines []string) ![]int {
	mut res := []int{}
	for l in lines {
		mut cal := 0
		for num in l.split('\n') {
			cal += num.int()
		}
		res << cal
	}
	res.sort(a > b)
	return res
}