package combinations

/* translated from my original version of implementation was perl.
 * the method is called "no crossed finger".
 * I will update the explanation after uploading perl version.
 */
func SomeIndexes( M int, N int ) [][]int {
	// M: number of selection ( 0 ... (M-1) )
	// N: number of choice
	if M < N {
		return [][]int{}
	}

	initRoomSize := M - N
	room := make([]int, N)
	pos  := []int{}
	for i := 0; i < N; i++ {
		pos = append (pos, i);
	}

	// https://stackoverflow.com/questions/39984957/is-it-possible-to-initialize-slice-with-specific-values
	for i := range room {
		room[i] = initRoomSize
	}

	var combis [][]int
	new_case := make([]int, N)
	copy(new_case, pos)
	combis = append( [][]int{}, new_case )

	cursor := N - 1 // initial: index of last elements in selection

	for {
		if room[cursor] > 0 {
			room[cursor]--
			pos[cursor]++
			new_case := make([]int, N)
			copy(new_case, pos)
			combis = append(combis, new_case)
		} else {
			cursor_moved := false
			for i := cursor; i > 0; i-- {
				if room[i-1] > 0 {
					cursor = i-1
					cursor_moved = true
					break
				}
			}
			if cursor_moved {
				new_room := room[cursor] - 1
				base_pos := pos[cursor];
				for p, i := 1, cursor; i < N; i++ {
					room[i] = new_room
					pos[i]  = base_pos + p
					p++ // p++, i++ not working on for()
				}
				new_case := make([]int, N)
				copy(new_case, pos)
				combis = append(combis, new_case)
				cursor = N - 1
			} else {
				break
			}
		}

	}

	return combis
}

func AllIndexes( M int ) [][] int {
	var all_combis [][]int

	for n := 1; n <= M; n++ {
		for _, cb := range SomeIndexes( M, n ) {
			all_combis = append(all_combis, cb)
		}
	}

	return all_combis
}
