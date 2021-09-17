package world

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	// 定义一个表示测试用例的结构体
	type test struct {
		str  string
		want bool
	}
	//用一个map表示一个测试组
	tests := map[string]test{
		"simple":       {"沙河有沙又有河", false},
		"englishFalse": {"abc", false},
		"englishTrue":  {"abcba", true},
		"ChineseTrue":  {"油灯少灯油", true},
		"withXx":       {"Madam,I'mAdam", true},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := IsPalindrome(tc.str)
			if got != tc.want {
				t.Errorf("want:%#v,got:%#v", tc.want, got)
			}
		})
	}
}
