package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
Design an algorithm to encode a list of strings to a string. The encoded string
is then sent over the network and is decoded back to the original list of
strings.

Machine 1 (sender) has the function:

string encode(vector<string> strs) {
  // ... your code
  return encoded_string;
}

Machine 2 (receiver) has the function:

vector<string> decode(string s) {
  //... your code
  return strs;
}

So Machine 1 does:

string encoded_string = encode(strs);

and Machine 2 does:

vector<string> strs2 = decode(encoded_string);

strs2 in Machine 2 should be the same as strs in Machine 1.

Implement the encode and decode methods.

You are not allowed to solve the problem using any serialize methods (such as
eval).



Example 1:

Input: dummy_input = ["Hello","World"]
Output: ["Hello","World"]
Explanation:
Machine 1:
Codec encoder = new Codec();
String msg = encoder.encode(strs);
Machine 1 ---msg---> Machine 2

Machine 2:
Codec decoder = new Codec();
String[] strs = decoder.decode(msg);

Example 2:

Input: dummy_input = [""]
Output: [""]



Constraints:

    1 <= strs.length <= 200
    0 <= strs[i].length <= 200
    strs[i] contains any possible characters out of 256 valid ASCII characters.



Follow up: Could you write a generalized algorithm to work on any possible set of characters?
*/

type Codec struct{}

// Encodes a list of strings to a single string.
func (codec *Codec) Encode(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var res strings.Builder
	for _, str := range strs {
		res.WriteString(fmt.Sprint(len(str), ","))
		res.WriteString(str)
	}
	return res.String()
}

// Decodes a single string to a list of strings.
func (codec *Codec) Decode(strs string) []string {
	if len(strs) == 0 {
		return []string{""}
	}
	result := make([]string, 0)
	for i := 0; i < len(strs); {
		index := strings.Index(strs[i:], ",")
		length, _ := strconv.Atoi(strs[i : i+index])
		word := string(strs[i+index+1 : i+index+length+1])
		fmt.Println(i, strs[i:], index, strs[i:i+index], word)
		result = append(result, word)
		i += index + 1 + length
	}
	return result
}

// Your Codec object will be instantiated and called as such:
// var codec Codec
// codec.Decode(codec.Encode(strs));

func main() {
	var codec Codec
	fmt.Println(codec.Decode(codec.Encode([]string{"Hello", "World"})))
}
