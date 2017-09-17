package main

import "fmt"

const(
    white = iota
    black
    blue
    red
    yellow
)

type Color byte
type Box struct {
    width, height, depth float64
    color Color
}

type BoxList []Box

func (b Box) Volume() float64 {
    return b.width * b.height * b.depth
}

// 方法
func (b *Box) SetColor(c Color) {
    b.color = c
}

// 方法
func (bl BoxList) BiggestColor() Color {
    v := 0.0
    k := Color(white)
    for _, b := range bl {
        if bv := b.Volume(); bv > v {
            v = bv
            k = b.color
        }
    }
    return k
}

// 方法
func (bl BoxList) PainItBlack() {
    for i, _ := range bl {
        bl[i].SetColor(black)
    }
}

func (c Color) String() string {
    strings := []string {"white", "black", "blue", "red", "yellow"}
    return strings[c]
}

func main() {
    boxes := BoxList {
        Box{4, 4, 4, red},
        Box{1, 2, 3, yellow},
        Box{2, 3, 4, black},
        Box{2, 2, 2, blue},
    }

    fmt.Printf("total boxes: %d \n", len(boxes))
    fmt.Println("first volume: ", boxes[0].Volume())
    fmt.Println("color of last box: ", boxes[len(boxes)-1].color.String())
    fmt.Println("the biggest box: ", boxes.BiggestColor().String())

    boxes.PainItBlack()
    fmt.Println("color of secode: ", boxes[1].color.String())

    fmt.Println("bigest box: ", boxes.BiggestColor().String())

}
