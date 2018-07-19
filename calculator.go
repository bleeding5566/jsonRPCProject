
package main

import(
	"unicode"
	"strconv"
) 

type Num struct{
	IntNum        string
	LenAfterPoint int
	Sign          bool
}

func smaller(s1, s2 string) bool{
    if len(s1) < len(s2){
        return true
    } else if len(s1)==len(s2){
        for i:=0;i<len(s1);i++{
            if s1[i] > s2[i]{
                break
            } else if s1[i]<s2[i]{
                return true
            }
        }
    }
    return false
}

func divString(s1, s2 string) string{
    var ret = ""
    var i int
    if len(s1) > len(s2){
        i = len(s2)
    } else {
        i = len(s1)
    }
    rem := s1[:i]
    lenAfterPoint := 0
    for lenAfterPoint <16{
        tmp := 0
        for !smaller(rem,s2){
            tmp++
            rem, _ = subString(rem,s2)
            zeroIdx :=0
			// Remove Front Zeros 
            for ; zeroIdx<len(rem)-1;zeroIdx++{
                if rem[zeroIdx]!='0'{
                    break
                }
            }
            rem = rem[zeroIdx:]
        }
        ret+= strconv.Itoa(tmp)
        if rem=="0"{
            break
        } else{
            if i==len(s1){
                ret+="."
                rem+="0"
            } else if i<len(s1){
                rem+=string(s1[i])
            } else {
                rem+="0"
                lenAfterPoint++
            }
            i++
        }
    }
    ret = removeZeros(ret)
    return ret
}

func divNum(s1Num,s2Num Num) string{
    s1, s2, _ := matchPoint(s1Num, s2Num)
	s1Sign := s1Num.Sign
	s2Sign := s2Num.Sign
    ret := ""
    sign := (!s1Sign && s2Sign) || (s1Sign && !s2Sign)
    ret =  divString(s1,s2)
    if sign{
        ret = "-" + ret
    }
    return ret
}

func mutString(s1,s2 string) string{
    s1Len  := len(s1)
    s2Len  := len(s2)
    ret := ""
    for i := 0; i < s1Len + s2Len; i++{
        ret+="0"
    }
    for s1Idx := s1Len-1; s1Idx >= 0; s1Idx--{
        s1Digit := int(s1[s1Idx] - '0')
        carry := 0
        for s2Idx := s2Len-1; s2Idx>=0; s2Idx--{
            s2Digit := int(s2[s2Idx] - '0')
            retIdx := s1Idx + s2Idx + 1
            sum := carry + s1Digit*s2Digit + int(ret[retIdx] - '0')
            carry = sum/10
            sum = sum%10
            ret = ret[:retIdx] + strconv.Itoa(sum) + ret[retIdx+1:]
        }
        ret = ret[:s1Idx] + strconv.Itoa(carry) + ret[s1Idx+1:]
    }
    return ret
}


func mutNum(s1Num,s2Num Num) string{
    s1, s2, lenAfterPoint := matchPoint(s1Num, s2Num)
	s1Sign := s1Num.Sign
	s2Sign := s2Num.Sign
    ret := ""
    sign := (!s1Sign && s2Sign) || (s1Sign && !s2Sign)
    ret =  mutString(s1,s2)
    PointLoc := len(ret) - 2*lenAfterPoint
    ret = ret[:PointLoc] + "." + ret[PointLoc:]
    ret = removeZeros(ret)
    if sign{
        ret = "-" + ret
    }
    return ret
}

func subString(s1,s2 string) (string, bool){
	if smaller(s1,s2) {
		ret, _ := subString(s2,s1)
		return ret, true
	}
    s1Idx := len(s1)-1
    s2Idx := len(s2)-1
    carry := 0
    ret := ""
    for s1Idx >=0 || s2Idx >=0 || carry<0{
        s1Digit := 0
        s2Digit := 0
        if s1Idx >= 0{
            s1Digit = int(s1[s1Idx] - '0')
            s1Idx--
        }
        
        if s2Idx >= 0{
            s2Digit = int(s2[s2Idx] - '0')
            s2Idx--
        }
        sum := s1Digit - s2Digit + carry
        if sum<0{
            carry = -1
            sum += 10
        } else {
            carry = 0
        }
        ret= strconv.Itoa(sum) + ret
    }
    return ret, false
}

func subNum(s1Num,s2Num Num) string{
	s1, s2, lenAfterPoint := matchPoint(s1Num, s2Num)
	s1Sign := s1Num.Sign
	s2Sign := s2Num.Sign
    ret := ""
    sign := false
    if !s1Sign && !s2Sign{
        ret, sign  = subString(s1,s2)
        
    } else if s1Sign && s2Sign{
        ret,sign = subString(s2,s1)
    } else if !s1Sign && s2Sign{
        ret = addString(s1,s2)
        sign = false
    } else {
        ret = addString(s1,s2)
        sign = true
    }
    var PointLoc = len(ret) - lenAfterPoint
    ret = ret[:PointLoc] + "." + ret[PointLoc:]
    ret = removeZeros(ret)
    if sign{
        ret = "-" + ret
    }
    return ret
}

func addString(s1,s2 string) string{
    s1Idx := len(s1)-1
    s2Idx := len(s2)-1
    carry := 0
    ret := ""
    for s1Idx >= 0 || s2Idx >= 0 || carry>0{
        s1Digit := 0
        s2Digit := 0
        if s1Idx >= 0{
            s1Digit = int(s1[s1Idx] - '0')
            s1Idx--
        }
        if s2Idx >=0 {
            s2Digit = int(s2[s2Idx] - '0')
            s2Idx--
        }
        sum := s1Digit + s2Digit + carry
        if sum>=10{
            carry = 1
            sum -= 10
        } else {
            carry = 0
        }
        ret = strconv.Itoa(sum) + ret
    }
    return ret
}

func addNum(s1Num,s2Num Num) string{
	s1, s2, lenAfterPoint := matchPoint(s1Num, s2Num)
	s1Sign := s1Num.Sign
	s2Sign := s2Num.Sign
    ret := ""
    sign := false
    if (s1Sign && s2Sign) || (!s1Sign && !s2Sign){
        ret  = addString(s1,s2)
        sign = s1Sign
    } else if !s1Sign && s2Sign{
        ret,sign = subString(s1,s2)
    } else {
        ret,sign = subString(s2,s1)
    }
    var PointLoc = len(ret) - lenAfterPoint
    ret = ret[:PointLoc] + "." + ret[PointLoc:]
    ret = removeZeros(ret)
    if sign{
        ret = "-" + ret
    }
    return ret
}

func removeZeros(s string) string{
    PIdx:=0
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
    s = s[start:end+1]
    return s
}

func check(s string) (Num , bool){
    PointExist := false
    SNum := Num{}
    IntNum := ""
    LenAfterPoint := 0
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
            }
        } else if !unicode.IsDigit(ch){
            return SNum, false
        } else{
            IntNum += string(ch)
            if PointExist{
                LenAfterPoint++
            }
        }
    }
	// No number exist after point
	if LenAfterPoint==0 && PointExist {
		return SNum, false
	}
	// No number exist 
    if len(IntNum)==0{
        return SNum, false
    }
    SNum.IntNum = IntNum
    SNum.LenAfterPoint = LenAfterPoint
    SNum.Sign = Sign
    return SNum, true
}

func matchPoint(s1Num,s2Num Num) (string, string, int){
	s1 := s1Num.IntNum
    s2 := s2Num.IntNum
    s1LenAfterPoint := s1Num.LenAfterPoint
    s2LenAfterPoint := s2Num.LenAfterPoint
    for s1LenAfterPoint < s2LenAfterPoint{
        s1 += "0"
        s1LenAfterPoint++
    }
    for s1LenAfterPoint > s2LenAfterPoint{
        s2 += "0"
        s2LenAfterPoint++
    }
	return s1, s2, s1LenAfterPoint
}

