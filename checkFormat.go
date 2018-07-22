
package main

import(
	"unicode"
) 

func removeZeros(s string) string{
	PIdx :=0
	ret := ""
	for ; PIdx<len(s); PIdx++{
		if s[PIdx] == '.'{
			break
		}
	}
	start := 0
	end := len(s) - 1
	for ; start < PIdx-1; start++{
		if s[start] != '0'{
			break
		}
	}
	for ; end > PIdx; end--{
		if s[end] != '0'{
			break
		}
	}
	ret += s[start:PIdx]
	if end > PIdx {
		ret += s[PIdx:end+1]
	}
	return ret
}

func check(s string) (Num , bool){
	PointExist := false
	SNum := Num{}
	IntNum := ""
	Sign := false
	if s[0] == '-'{
		s = s[1:]
		Sign = true
	}
	for _,ch := range(s){
		if ch =='.'{
			if PointExist{
				return SNum, false
			} else{
				PointExist = true
				IntNum += "."
			}
		} else if !unicode.IsDigit(ch){
			return SNum, false
		} else{
			IntNum += string(ch)
		}
	}
	if len(IntNum)==0 || (len(IntNum)==1 && IntNum[0]=='.') {
		return SNum , false
	}
	if IntNum[0]=='.' {
		IntNum = "0" + IntNum 
	}
	IntNum = removeZeros(IntNum)
	i := 0
	for ; i < len(IntNum) ; i++ {
		if IntNum[i] == '.'{
			SNum.LenAfterPoint = len(IntNum) - i - 1
			break
		}
	}
	if i < len(IntNum) {
		IntNum = IntNum[:i] + IntNum[i+1:]
	}
	SNum.IntNum = IntNum
	SNum.Sign = Sign
	return SNum, true
}
