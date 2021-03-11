package main

import (
	"fmt"
	"strconv"
)

func decToHex(target int) ([]string, error) {
	amari := []string{}

	t := target
	for {
		result := t % 16 // 余り
		amari = append(amari, strconv.Itoa(result))
		t = t / 16 // 商
		fmt.Println(result)
		if t == 0 {
			break
		}
	}

	// 余りの10 ~ 15をA ~ Fに置換
	for i, value := range amari {
		switch value {
		case "10":
			amari[i] = "A"
			break
		case "11":
			amari[i] = "B"
			break
		case "12":
			amari[i] = "C"
			break
		case "13":
			amari[i] = "D"
			break
		case "14":
			amari[i] = "E"
			break
		case "15":
			amari[i] = "F"
			break
		}
	}

	// 逆に並び替える
	d := &DefaultUtils{list: amari}
	reversedAmari := d.Reverse()
	return reversedAmari, nil
}
