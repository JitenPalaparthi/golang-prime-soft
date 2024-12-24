
var arr1 [10]any --> [nil nil nil nil nil nil nil nil nil nil]

Sum of elements of that array 

arr1[0]=100
arr1[1]=true
arr1[2]="123.123"
arr1[3]=1231.123

// is nil do nothing
// if string --> try to parseInt or Float
// if bool -> true add 1 false do nothing

func SumOf(arr [10]any)float64{
sum:=0.0
// write logic here


return sum
}
