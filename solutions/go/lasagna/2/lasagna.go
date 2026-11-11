package lasagna


// 1. 定义预期的烤箱时间是 40 分钟
const OvenTime = 40


func RemainingOvenTime(actualMinutesInOven int) int {
	return OvenTime - actualMinutesInOven
}


func PreparationTime(numberOfLayers int) int {
	return numberOfLayers * 2
}

func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	// 总耗时 = 准备时间 + 烤箱中已经过的时间
	return PreparationTime(numberOfLayers) + actualMinutesInOven
}