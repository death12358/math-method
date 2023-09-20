package dot

import (
	"fmt"

	float64conversion "github.com/death12358/text-conversion/float64conversion"
)

func test() {
	float64conversion.Decimal([]float64{123})
}
func dot() {
	vector1 := []float64{1, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110, 120, 130, 140, 150, 160, 170, 180, 190, 200, 220, 240, 260, 280, 300, 350, 400, 450, 500, 1000}
	vector2 := []float64{700, 570, 570, 600, 550, 500, 400, 300, 100, 60, 60, 50, 50, 50, 20, 20, 20, 10, 10, 10}

	dotProd, err := dotProduct(vector1, vector2)
	if err != nil {
		fmt.Println("错误：", err)
		return
	}

	fmt.Printf("内积结果：%.2f\n", dotProd)
}

// []float64{700, 570, 570, 600, 550, 500, 400, 300, 100, 60, 60, 50, 50, 50, 20, 20, 20, 10, 10, 10},
// 	[]float64{0, 0, 0, 0, 0, 0, 500, 350, 400, 60, 60, 50, 50, 20, 20, 10, 0, 0, 0, 0}, //	21	FISHC06 財神

func dotProduct(a, b []float64) (float64, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("切片长度不相等")
	}

	var result float64
	for i := 0; i < len(a); i++ {
		result += a[i] * b[i]
	}
	return result, nil
}