package main

import (
	"fmt"
)

/*
Design a time-based key-value data structure that can store multiple values for
the same key at different time stamps and retrieve the key's value at a certain
timestamp.

Implement the TimeMap class:

    TimeMap() Initializes the object of the data structure.
	void set(String key, String value, int timestamp) Stores the key key with
	the value value at the given time timestamp.
	String get(String key, int timestamp) Returns a value such that set was
	called previously, with timestamp_prev <= timestamp. If there are multiple
	such values, it returns the value associated with the largest
	timestamp_prev. If there are no values, it returns "".



Example 1:

Input
["TimeMap", "set", "get", "get", "set", "get", "get"]
[[], ["foo", "bar", 1], ["foo", 1], ["foo", 3], ["foo", "bar2", 4], ["foo", 4], ["foo", 5]]
Output
[null, null, "bar", "bar", null, "bar2", "bar2"]

Explanation
TimeMap timeMap = new TimeMap();
timeMap.set("foo", "bar", 1);  // store the key "foo" and value "bar" along with timestamp = 1.
timeMap.get("foo", 1);         // return "bar"
timeMap.get("foo", 3);         // return "bar", since there is no value corresponding to foo at timestamp 3 and timestamp 2, then the only value is at timestamp 1 is "bar".
timeMap.set("foo", "bar2", 4); // store the key "foo" and value "bar2" along with timestamp = 4.
timeMap.get("foo", 4);         // return "bar2"
timeMap.get("foo", 5);         // return "bar2"



Constraints:

    1 <= key.length, value.length <= 100
    key and value consist of lowercase English letters and digits.
    1 <= timestamp <= 107
    All the timestamps timestamp of set are strictly increasing.
    At most 2 * 105 calls will be made to set and get.

*/

type TimeMap struct {
	m map[string][]Node
}

type Node struct {
	timestamp int
	value     string
}

func Constructor() TimeMap {
	t := new(TimeMap)
	t.m = make(map[string][]Node)
	return *t
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.m[key] = append(this.m[key], Node{timestamp: timestamp, value: value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	k := this.m[key]
	n := len(k)
	l, r := 0, n
	mx := 0
	for l < r {
		mid := l + (r-l)/2
		if k[mid].timestamp <= timestamp {
			l = mid + 1
			mx = max(mid, mx)
		} else if k[mid].timestamp > timestamp {
			r = mid
		}
	}
	if l == 0 {
		return ""
	}
	return k[mx].value
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
func main() {
	obj := Constructor()
	obj.Set("foo", "bar", 1)
	fmt.Println(obj.Get("foo", 1))
	fmt.Println(obj.Get("foo", 3))
	obj.Set("foo", "bar2", 4)
	fmt.Println(obj.Get("foo", 4))
	fmt.Println(obj.Get("foo", 5))
}
