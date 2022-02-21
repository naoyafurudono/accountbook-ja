// STEP03: データの記録

package main

import "fmt"

// DONE:
// 品目と値段を一緒に扱うために
// Itemという構造体の型を定義する
// Categoryという品目を入れる文字列型のフィールドを持つ
// Priceという値段を入れる整数型のフィールドを持つ
type Item struct {
	Category string
	Price    int
}

func main() {

	// TODO:
	// inputItemという関数を呼び出し
	// 結果をitemという変数に入れる
	var item = inputItem()

	fmt.Println("===========")

	// TODO:
	// 品目に「コーヒー」、値段に「100」と入力した場合に
	// 「コーヒーに100円使いました」と表示する
	fmt.Printf("%sに%d円使いました\n", item.Category, item.Price)

	fmt.Println("===========")
}

// 入力を行う関数
// 入力したItemを返す
func inputItem() Item {
	// Item型のitemという名前の変数を定義する
	var item Item

	fmt.Print("品目>")
	fmt.Scan(&item.Category)

	fmt.Print("値段>")
	fmt.Scan(&item.Price)

	return item
}
