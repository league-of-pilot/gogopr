package chanogo

import "fmt"

type structExport struct {
	name string
	Age  int
}

func CreatePrivate() structExport {
	return structExport{
		name: "chanogo",
		Age:  23,
	}
}

type StructPub struct {
	pubName string
	PubAge  int
}

func CreatePublic() StructPub {
	return StructPub{
		pubName: "chanogo",
		PubAge:  23,
	}
}

func Unpack() {

	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])
	prices[1] = 9.99

	prices = append(prices, 5.99, 12.99, 29.99, 100.10)
	prices = prices[1:]
	fmt.Println(prices)

	discountPrices := []float64{101.99, 80.99, 20.59}
	prices = append(prices, discountPrices...)
	fmt.Println(prices)
}
