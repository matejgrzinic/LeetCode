package main

type RecentCounter struct {
	Data []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	this.Data = append(this.Data, t)
	for i, v := range this.Data {
		if t-v <= 3000 {
			this.Data = this.Data[i:]
			break
		}
	}
	return len(this.Data)
}
