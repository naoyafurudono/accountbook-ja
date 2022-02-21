// STEP06: ブラッシュアップ

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// AccountBookをNewAccountBookを使って作成
	ab, err := NewAccountBook("accountbook.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "fail to open/create file")
	}

LOOP: // 以下のループにラベル「LOOP」をつける
	for {

		// モードを選択して実行する
		var mode int
		fmt.Println("[1]入力 [2]最新10件 [3]終了")
		fmt.Printf(">")
		fmt.Scan(&mode)

		// モードによって処理を変える
		switch mode {
		case 1: // 入力
			var n int
			fmt.Print("何件入力しますか>")
			fmt.Scan(&n)

			for i := 0; i < n; i++ {
				if err := ab.AddItem(inputItem()); err != nil {
					fmt.Fprintf(os.Stderr, "入力に失敗 %s", err.Error())
					break LOOP
				}
			}
		case 2: // 最新10件
			items, err := ab.GetItems(10)
			if err != nil {
				fmt.Fprintln(os.Stderr, "エラー:", err)
				break LOOP
			}
			showItems(items)
		case 3: // 終了
			fmt.Println("終了します")
			os.Exit(0)
		default:
			continue
		}
	}
}

type Item struct {
	Category string
	Price    int
}

type accountbook struct {
	filename string
}

func NewAccountBook(filename string) (ab accountbook, err error) {
	// file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	ab = accountbook{filename: filename}
	return
}

func (ab accountbook) AddItem(item *Item) error {
	file, err := os.OpenFile(ab.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(fmt.Sprintf("%s %d\n", item.Category, item.Price))
	return err
}

// Itemを入力し返す
func inputItem() *Item {
	var item Item

	fmt.Print("品目>")
	fmt.Scan(&item.Category)

	fmt.Print("値段>")
	fmt.Scan(&item.Price)

	return &item
}

func parseItem(s string) (*Item, error) {
	itemarr := strings.Split(s, " ")
	if len(itemarr) != 2 {
		return nil, errors.New("fail to parse item")
	}

	price, err := strconv.Atoi(itemarr[1])
	if err != nil {
		return nil, err
	}
	item := &Item{Category: itemarr[0], Price: price}
	return item, err
}

func (ab accountbook) GetItems(n int) ([]*Item, error) {
	items := make([]*Item, 0, n)
	file, err := os.Open(ab.filename)
	if err != nil {
		return items, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	oldest := 0
	for i := 0; scanner.Scan(); i++ {
		ip, err := parseItem(scanner.Text())
		if err != nil {
			return nil, err
		}
		if i < n {
			items = append(items, ip)
		} else {
			items[oldest] = ip
			oldest = (oldest + 1) % n
		}
	}

	// 並び替える
	res := make([]*Item, 0, n)
	for i := 0; i < len(items); i++ {
		res = append(res, items[(i+oldest)%n])
	}
	return res, nil
}

// Itemの一覧を出力する
func showItems(items []*Item) {
	fmt.Println("===========")
	// itemsの要素を1つずつ取り出してitemに入れて繰り返す
	for _, item := range items {
		fmt.Printf("%s:%d円\n", item.Category, item.Price)
	}
	fmt.Println("===========")
}
