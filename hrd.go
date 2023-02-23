package main

import (
	"flag"
	"log"
	"pkg_test/HRD"
	"time"
)

var ch = make(chan *HRD.StateNode, 100000)

func main() {
	s := flag.String("m", "", "init hrd map")

	flag.Parse()
	if *s == "" {
		log.Fatal("请输入初始值")
	}

	if len(*s) != 20 {
		log.Fatal("初始值无效")
	}

	initMap := HRD.StoM(*s)
	log.Println("********init map********")
	HRD.PrintMv1(initMap)
	HRD.AddToMap(initMap)
	log.Println("********init map********")
	startTime := time.Now().UnixMilli()
	ch <- &HRD.StateNode{
		Pre:   nil,
		State: HRD.MtoI(initMap),
	}
	log.Println("Starting calculating...ヾ(◍°∇°◍)ﾉﾞ")
	total := 1
	for len(ch) > 0 {
		n := <-ch
		m := HRD.ItoM(n.State)
		if m[3][1] == HRD.SQUARE {
			HRD.PrintStep(n)
			log.Printf("total time is %dms", time.Now().UnixMilli()-startTime)
			return
		}
		for i := 0; i < HRD.ROWSIZE; i++ {
			for j := 0; j < HRD.COLUMNSIZE; j++ {
				nodes := make([]*HRD.StateNode, 0)
				switch m[i][j] {
				case HRD.BLOCK:
					nodes = HRD.BlockMove(m, i, j, n)
				case HRD.VERTICAL:
					nodes = HRD.VerticalMove(m, i, j, n)
				case HRD.HORIZONTAL:
					nodes = HRD.HorizontalMove(m, i, j, n)
				case HRD.SQUARE:
					nodes = HRD.SquareMove(m, i, j, n)
				default:

				}
				for _, node := range nodes {
					ch <- node
					total++
					log.Printf("total node %d", total)
					log.Printf("ch len is %d", len(ch))
				}
			}
		}
	}
	log.Println("Failed to find solution! o(╥﹏╥)o")
}
