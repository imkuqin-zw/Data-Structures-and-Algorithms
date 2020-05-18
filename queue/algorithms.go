package queue

import (
	"container/list"
)

func openLock(deadends []string, target string) int {
	dead := make(map[string]bool)
	exists := make(map[string]bool)
	for _, item := range deadends {
		dead[item] = true
	}
	if _, ok := dead[target]; ok {
		return -1
	}
	up := func(s string, j int) string {
		ch := []byte(s)
		if '0' == ch[j] {
			ch[j] = '9'
		} else {
			ch[j] -= 1
		}
		return string(ch)
	}
	down := func(s string, j int) string {
		ch := []byte(s)
		if '9' == ch[j] {
			ch[j] = '0'
		} else {
			ch[j] += 1
		}
		return string(ch)
	}
	q := list.New()
	q.PushBack("0000")
	exists["0000"] = true
	step := 0
	for q.Len() > 0 {
		size := q.Len()
		for j := 0; j < size; j++ {
			e := q.Front()
			tmp := e.Value.(string)
			q.Remove(e)
			if _, ok := dead[tmp]; ok {
				continue
			}
			if tmp == target {
				return step
			}
			for i := 0; i < 4; i++ {
				upNext := up(tmp, i)
				if _, ok := exists[upNext]; !ok {
					q.PushBack(upNext)
					exists[upNext] = true
				}
				downNext := down(tmp, i)
				if _, ok := exists[downNext]; !ok {
					q.PushBack(downNext)
					exists[downNext] = true
				}
			}
		}
		step++
	}
	return -1
}
