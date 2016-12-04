package main

import (
	"fmt"
	"math"
)

var (
	Pr, T, fsc                        float64
	sw                                float64
	Pwf                               float64
	p1, p2, p3                        float64
	sw1, sw2, sw3                     float64
	freewater, sportwater, waterratio float64
)

func WH2O(Pr, T, fsc float64) float64 {
	p10nd2 := math.Pow(10, -2)
	p10nd4 := math.Pow(10, -4)
	A := 3.14 + 418.0278/Pr
	B := 3.2147 + 3.8537*p10nd2*Pr - 4.7752*p10nd4*math.Pow(Pr, 2)
	W := 1.6019 * p10nd4 * A * math.Pow((0.32*(5.625*p10nd2*T+1)), B) * fsc
	return W
}

func main() {
	for { //循环计算
		fmt.Println("气藏地层水水型比例计算程序 v1.0")
		fmt.Printf("算法制作：张慧先；软件制作：侯莹杰\n\n")
		fmt.Println("步骤1：计算凝析水量")
		fmt.Println("输入数据：Pr T fsc（数值以空格间隔，其中 Pr：地层压力(MPa)；T：气藏温度(T)；fsc:含盐量校正系数)")
		fmt.Scanln(&Pr, &T, &fsc)

		lw := WH2O(Pr, T, fsc)

		fmt.Printf("凝析水量WH2O= %v m3/10^4m3\n\n", lw) //%v 全格式.

		fmt.Println("步骤2：比较水气比与凝析水量")
		fmt.Println("输入水气比Sw(%)：")
		fmt.Scanln(&sw)
		if sw <= lw {
			fmt.Printf("该气藏只有凝析水\n凝析水量WH2O= %v m3/10^4m3\n\n", lw)
		} else {
			fmt.Println("输入井底压力 Pwf(MPa)")
			fmt.Scanln(&Pwf)
			fmt.Println("输入 Pwg(Sw1) Sw1 (Pwg(Sw1)：原始含水饱和度对应压力(MPa)；Sw1：原始含水饱和度(%))")
			fmt.Scanln(&p1, &sw1)
			fmt.Println("输入 Pwg(Sw2) Sw2 (Pwg(Sw2)：临界含水饱和度对应压力(MPa)；Sw2 :临界含水饱和度(%))")
			fmt.Scanln(&p2, &sw2)
			fmt.Println("输入 Pwg(Sw3) Sw3 (Pwg(Sw3)：束缚水饱和度对应压力(MPa)；Sw3 :束缚水饱和度(%))")
			fmt.Scanln(&p3, &sw3)
			dp := Pr - Pwf
			dp1 := p2 - p1
			dp2 := p3 - p2
			freewater = sw1 - sw2
			sportwater = sw2 - sw3
			waterratio = freewater / sportwater

			if dp <= dp1 {
				fmt.Println("该地层水中全部是自由水")
				fmt.Printf("凝析水量 = %v m3/10^4m3\n", lw)
				fmt.Printf("dP1 = %v Mpa\n", dp1)
				fmt.Printf("自由水量 = %v %\n", freewater)
				fmt.Println()
			} else {
				fmt.Println("该地层水中包含有自由水和可动水")
				fmt.Printf("凝析水量 = %v m3/10^4m3\n", lw)
				fmt.Printf("dP1 = %v Mpa\n", dp1)
				fmt.Printf("自由水量 = %v %\n", freewater)
				fmt.Printf("dP2 = %v Mpa\n", dp2)
				fmt.Printf("可动水量 = %v %\n", sportwater)
				fmt.Printf("自由水量/可动水量 = %v %\n", waterratio)
				fmt.Println()
			}
		}
	}
}
