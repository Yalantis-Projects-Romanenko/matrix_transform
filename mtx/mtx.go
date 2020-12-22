package mtx

func SpiralTransform(matrix [][]int) []int {
	columns := len(matrix[0])
	rows := len(matrix)
	totalLen := columns * rows

	result := make([]int, 0, totalLen)

	x, y := 0, 0
	state := 3

	xMin, xMax, yMin, yMax := 0, columns-1, 0, rows-1

	for i := 0; i < totalLen; i++ {
		result = append(result, matrix[y][x])

		/*
		   shift
		   state is a value that dictates direction
		   0 = \/;
		   1 = <;
		   2 = /\;
		   3 = >;
		*/
		shift := func(x,y int) (int, int) {
			switch state {
			case 0:
				y++
				break
			case 1:
				x--
				break
			case 2:
				y--
				break
			case 3:
				x++
				break
			}
			return x, y
		}

		rotatePhase := func() {
			state++
			if state > 3 {
				state = 0
			}
			x, y = shift(x,y)
		}

		x1, y1 := shift(x,y)
		switch {
		case x1 < xMin:
			yMax--
			rotatePhase()
			break
		case y1 < yMin:
			xMin++
			rotatePhase()
			break
		case x1 > xMax:
			yMin++
			rotatePhase()
			break
		case y1 > yMax:
			xMax--
			rotatePhase()
			break
		default:
			x, y = x1, y1
			break
		}
		//fmt.Println(result, state)
	}

	return result
}
