package pkg

import "fmt"

type RoadStrategy struct {
}

func (r *RoadStrategy) Route(startPoint int, endPoint int) {
	avgSpeed := 30
	trafficJam := 2
	total := endPoint - startPoint
	totalTime := total * avgSpeed * trafficJam

	fmt.Printf("Road A:[%d] to B:[%d] Avg speed:[%d] TrafficJam:[%d] Total:[%d] Total time[%d] min\n",
		startPoint, endPoint, avgSpeed, trafficJam, total, totalTime)
}
