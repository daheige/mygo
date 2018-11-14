package main

import (
	"bytes"
	"fmt"
)

const target int = 32 << (^uint(0) >> 63) //判断当前系统是32位还是64位

type IntSet struct {
	words []uint
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/target, uint(x%target)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(numbers ...int) {
	for _, numebr := range numbers {
		x := numebr
		word, bit := x/target, uint(x%target)
		for word >= len(s.words) {
			s.words = append(s.words, 0)
		}
		s.words[word] |= 1 << bit
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

//并集，合并s,t
func UnionWith(s, t *IntSet) *IntSet {
	uw := &IntSet{}
	slen, tlen := len(s.words), len(t.words)
	minlen := slen
	if minlen > tlen {
		minlen = tlen
	}
	for i := 0; i < minlen; i++ {
		uw.words = append(uw.words, s.words[i]|t.words[i])
	}
	if minlen == tlen {
		union(uw, s, minlen, slen)
	} else {
		union(uw, t, minlen, tlen)
	}
	return uw
}

//交集,元素在s,t中均出现
func InersectWith(s, t *IntSet) *IntSet {
	iw := &IntSet{}
	slen, tlen := len(s.words), len(t.words)
	minlen := slen
	if minlen > tlen {
		minlen = tlen
	}
	for i := 0; i < minlen; i++ {
		iw.words = append(iw.words, s.words[i]&t.words[i])
	}
	if minlen == tlen { //s元素更多, 多余的元素全部置为0
		for minlen < slen {
			iw.words = append(iw.words, s.words[minlen]&0)
			minlen++
		}
	}
	return iw
}

//差集，元素出现在s，未出现在t
func DifferenceWith(s, t *IntSet) *IntSet {
	dw := &IntSet{}
	slen, tlen := len(s.words), len(t.words)
	minlen := slen
	if minlen > tlen {
		minlen = tlen
	}
	for i := 0; i < minlen; i++ {
		mask := s.words[i] & t.words[i]
		dw.words = append(dw.words, s.words[i]^mask)
	}
	if minlen == tlen && minlen < slen {
		union(dw, s, minlen, slen)
	}
	return dw
}

//并差集，元素出现在s但未出现在t，或者元素出现在t但未出现在s
func SymmetricDifference(s, t *IntSet) *IntSet {
	sd := &IntSet{}
	slen, tlen := len(s.words), len(t.words)
	minlen := slen
	if minlen > tlen {
		minlen = tlen
	}
	for i := 0; i < minlen; i++ {
		sd.words = append(sd.words, s.words[i]^t.words[i])
	}
	if minlen == tlen {
		union(sd, s, minlen, slen)
	} else {
		union(sd, t, minlen, tlen)
	}
	return sd
}

//从tmplen开始，将is中的元素复制到res
func union(res, is *IntSet, tmplen, xlen int) {
	for tmplen < xlen {
		res.words = append(res.words, is.words[tmplen])
		tmplen++
	}
}

func main() {
	var x, y IntSet
	x.Add(0)
	x.Add(1, 3, 5, 7)
	x.Add(9)
	x.Add(63)
	x.Add(100)
	fmt.Println("x : ", x.String())

	y.Add(9)
	y.Add(42)
	y.Add(7)
	y.Add(5)
	fmt.Println("y : ", y.String())

	result1 := UnionWith(&x, &y)
	fmt.Println("x UnionWith y : ", result1.String())
	result2 := InersectWith(&x, &y)
	fmt.Println("x InersectWith y : ", result2.String())
	result3 := DifferenceWith(&x, &y)
	fmt.Println("x DifferenceWith y : ", result3.String())
	result4 := SymmetricDifference(&x, &y)
	fmt.Println("x SymmetricDifference y : ", result4.String())

}
