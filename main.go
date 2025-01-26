// 3d donut but in golang
package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	//declarating variables
	var A float64 = 0.0
	var B float64 = 0.0

	z := [1760]float64{}
	b := [1760]byte{}

	fmt.Printf("\x1b[2J");
	
	for {
		for i := range z {
			z[i] = 0
			b[i] = ' '
		}

		for j := 0.0; j < 6.28; j += 0.07 {
			for i := 0.0; i < 6.28; i += 0.02 {
				c := math.Sin(i)
				d := math.Cos(j)
				e := math.Sin(A)
				f := math.Sin(j)
				g := math.Cos(A)
				h := d + 2
				D := 1 / (c*h*e + f*g + 5)
				l := math.Cos(i)
				m := math.Cos(B)
				n := math.Sin(B);
				t := c * h * g - f * e
				x := int(40 + 30*D * (l*h*m - t*n))
				y := int(12 + 15*D * (l*h*n + t*m))
				o := x + 80*y
				N := int(8 * ((f*e - c*d*g)*m - c*d*e - f*g - l*d*n))
				if 22 > y && y > 0 && x > 0 && 80 > x && D > z[o] {
					z[o] = D
					if N > 0 {
						b[o] = ".,-~:;=!*#$@"[N]
					} else {
						b[o] = '.'
					}
				}
			}
		}

		fmt.Printf("\x1b[H")
		for k := 0; k < 1760; k += 1 {
			if k%80==0 {
				fmt.Print("\n")
			} else {
				fmt.Printf("%c", b[k])
			}
		}

		A += 0.04
		B += 0.02
		time.Sleep(0 * time.Millisecond)
	}
}
