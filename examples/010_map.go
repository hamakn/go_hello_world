package examples

import "fmt"

type Geo struct {
	Lat, Long float64
}

func MapExample() {
	// 1. 普通に宣言する
	var map1 map[string]Geo
	map1 = make(map[string]Geo)
	map1["Shibuya Station"] = Geo{
		35.3931, 139.4204,
	}
	fmt.Println(map1)

	// 2. 型推論で宣言する
	map2 := make(map[string]Geo)
	map2["Shibuya Station"] = Geo{
		35.3931, 139.4204,
	}
	fmt.Println(map2)

	// 3. 宣言と同時に代入する
	var map3 = map[string]Geo{
		"Shibuya Station": Geo{
			35.3931, 139.4204,
		},
		"Nerima Station": Geo{
			35.44161, 139.39128,
		},
	}
	fmt.Println(map3)

	// 4. 3.で型名を省略する
	var map4 = map[string]Geo{
		"Shibuya Station": {
			35.3931, 139.4204,
		},
		"Nerima Station": {
			35.44161, 139.39128,
		},
	}
	fmt.Println(map4)

	// mapのvalueには代入はできない
	// map1["Shibuya Station"].Long = 10
	// // => cannot assign to struct field map1["Shibuya Station"].Long in map

	// mapへのaccessの2番目の返り値はvalueが存在するかのbool値になる
	var key string
	var g Geo
	var ok bool
	key = "Nerima Station"
	g, ok = map4[key]
	fmt.Println("Key:", key, "Value:", g, "Present?", ok)
	key = "Yokosuka Station"
	g, ok = map4[key]
	fmt.Println("Key:", key, "Value:", g, "Present?", ok)

	// deleteで消せる
	delete(map4, "Nerima Station")
	key = "Nerima Station"
	g, ok = map4[key]
	fmt.Println("Key:", key, "Value:", g, "Present?", ok)
}
