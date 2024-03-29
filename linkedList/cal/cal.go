package cal

import (
	"bytes"
	"dsa/linkedList/genericStack"
	_ "fmt"
	"strconv"
	_ "strings"
)

func Cal(str string) int {
	numStack := genericStack.InitGenericStack()
	operStack := genericStack.InitGenericStack()
	cursor := 0
	array := []rune(str)
	length := len(array)
	for cursor < length {
		if !IsOper(array[cursor]) && !IsBrackets(array[cursor]) { //如果是数字的话直接入数字栈 要判断是不是多位数
			if cursor == length-1 { //到切片的最后一位必然是个一位数
				temp, _ := strconv.Atoi(string(array[cursor]))
				genericStack.Push(numStack, temp)
			} else { //可能是个多位数

				var b bytes.Buffer
				b.Grow(length)
				b.WriteRune(array[cursor])
				for !IsOper(array[cursor+1]) && !IsBrackets(array[cursor+1]) { //下一位还是数字
					cursor++
					b.WriteRune(array[cursor])
				}
				tempNumber, _ := strconv.Atoi(b.String())
				genericStack.Push(numStack, tempNumber)

			}
			//游标往后移动
			cursor++
		} else if IsOper(array[cursor]) { //如果是运算符
			if genericStack.IsEmpty(operStack) { //如果运算栈为空直接加入
				genericStack.Push(operStack, array[cursor])
			} else {
				_, operInter := genericStack.Peek(operStack)
				operRune := operInter.(rune)
				if Priority(array[cursor]) > Priority(operRune) { //如果当前oper的优先级高于栈的oper优先级入栈
					genericStack.Push(operStack, array[cursor])
				} else { //如果当前运算符的优先级比较低则从numStack pop2个元素，从operStack pop一个运算符
					//再把结果push到numStack 当前运算符入栈
					genericStack.Iter(numStack)
					_, num1Inter := genericStack.Pop(numStack)
					_, num2Inter := genericStack.Pop(numStack)
					num1 := num1Inter.(int)
					num2 := num2Inter.(int)
					_, operInter := genericStack.Pop(operStack)
					oper := operInter.(rune)
					genericStack.Push(numStack, Mathematical(num1, num2, oper))
					genericStack.Push(operStack, array[cursor])
				}
			}

			cursor++
		} else if IsBrackets(array[cursor]) { //如果是括号找到匹配的括号最后运算必然得到一个数字值push到numStack中
			matchIndex := MatchBrackets(array[cursor+1:length]) + cursor + 1 //计算这里的matchIndex
			tempArray := array[cursor+1 : matchIndex]                        //得到新的rune
			resNum := Cal(string(tempArray))                                 //用括号的得到运算结果值
			genericStack.Push(numStack, resNum)                              //递归
			cursor = matchIndex + 1                                          //游标移动到matchIndex的后一位
		}

	}
	//计算最后的结果
	return complete(numStack, operStack)

}

func MatchBrackets(str []rune) int { //匹配括号返回坐标
	//左括号为1
	left, right := 1, 0 //左括号==1右括号==0
	resIndex := 0
	for index, temp := range str {
		if temp == '(' {
			left++
		}
		if temp == ')' {
			right++
		}
		if left == right {
			resIndex = index
			break
		}
	}
	return resIndex
}

func complete(numStack *genericStack.GenericStack, operStack *genericStack.GenericStack) int {
	for !genericStack.IsEmpty(operStack) {
		_, num1Inter := genericStack.Pop(numStack)
		_, num2Inter := genericStack.Pop(numStack)
		num1 := num1Inter.(int)
		num2 := num2Inter.(int)
		_, operInter := genericStack.Pop(operStack)
		oper := operInter.(rune)
		genericStack.Push(numStack, Mathematical(num1, num2, oper))
	}
	_, resIter := genericStack.Pop(numStack)
	res := resIter.(int)
	return res

}

func IsRightBrackets(ch rune) bool {
	return ch == ')'
}
func IsLeftBrackets(ch rune) bool {
	return ch == '('
}
func Priority(ch rune) int {
	if ch == '*' || ch == '/' {
		return 1
	} else {
		return 0
	}
}
func IsOper(ch rune) bool {
	return ch == '+' || ch == '-' || ch == '*' || ch == '/'
}
func IsBrackets(ch rune) bool {
	return ch == '(' || ch == ')'
}

func Mathematical(num1 int, num2 int, oper rune) int {
	total := 0
	switch oper {
	case '+':
		total = num1 + num2
	case '-':
		total = num2 - num1
	case '*':
		total = num1 * num2
	case '/': //todo 验证num1!=0
		total = num2 / num1
	default:
		total = 0
	}
	return total
}
