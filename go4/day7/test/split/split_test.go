package split

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	// 存放测试数据的结构体
	type test struct {
		str  string
		sep  string
		want []string
	}
	//创建一个存放所有测试用例的map
	var tests = map[string]test{
		"normal": test{"a:b:c", ":", []string{"a", "b", "c"}},
		"none":   test{"a:b:c", "*", []string{"a:b:c"}},
		"multi":  test{"abcfbcabcb", "bc", []string{"a", "f", "a", "b"}},
		"num":    test{"1231", "1", []string{"", "23", ""}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			t.Log("开始测试", name)
			ret := Split(tc.str, tc.sep)
			if !reflect.DeepEqual(ret, tc.want) {
				t.Errorf("期望得到:%#v,实际得到:%#v", tc.want, ret)
			}
		})
	}
}

//基准测试
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a:b:c", ":")
	}
}

//并行测试
func BenchmarkSplitParallel(b *testing.B) {
	//b.Setparallelism(1) //设置使用的CPU数
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Split("a:b:c", ":")
		}
	})
}
