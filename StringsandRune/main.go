package main

import (
	"fmt"
	"strings"
)

func main(){
	var mystring="rEsume"
	var indexed=mystring[0]
	fmt.Printf("%v , %T\n",indexed,indexed)

	for i,v:=range mystring{
		fmt.Println(i,v)
	}

	var myRune='a'
	fmt.Printf("%v , %T\n",myRune,myRune)

	var strSlice=[]string{"S","t","r","i","n","g","s"}
	var StrBuilder strings.Builder
	for i:=range strSlice{  //better than str+="" because it avoids creating a new string each time
		StrBuilder.WriteString(strSlice[i])
	}
	// fmt.Println(StrBuilder.String())
	var catStr= StrBuilder.String()
	fmt.Printf("%s , %T\n",catStr,catStr)
}