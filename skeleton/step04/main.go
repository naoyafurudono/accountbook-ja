// STEP04: 複数データの記録

package main

import "fmt"

type Item struct {
	Category string
	Price    int
}

func main() {
	var n int
	fmt.Print("件数>")
	fmt.Scan(&n)

	items := make([]Item, 0, n)

	for i := 0; i < cap(items); i++ {
		items = inputItem(items)
	}

	showItems(items)
}

func inputItem(items []Item) []Item {
	var item Item

	fmt.Print("品目>")
	fmt.Scan(&item.Category)

	fmt.Print("値段>")
	fmt.Scan(&item.Price)

	items = append(items, item)

	return items
}

// 一覧の表示を行う関数
func showItems(items []Item) {
	fmt.Println("===========")

	for i := 0; i < len(items); i++ {
		fmt.Printf("%s:%d円\n", items[i].Category, items[i].Price)
	}

	fmt.Println("===========")
}
