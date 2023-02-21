package solution_offer_2_63

import (
	"strings"
)

type node map[rune]node

func replaceWords(dictionary []string, sentence string) string {
	root := make(map[rune]node)
	for _, s := range dictionary {
		tmp := root
		for _, c := range s {
			if tmp[c] == nil {
				tmp[c] = node{}
			}
			tmp = tmp[c]
		}
		tmp['#'] = node{}
	}

	var rst []rune
	for _, s := range strings.Split(sentence, " ") {
		tmp := root
		try := true
		for _, c := range s {
			rst = append(rst, c)
			if try {
				if v, ok := tmp[c]; ok {
					if v['#'] != nil {
						break
					} else {
						tmp = v
					}
				} else {
					try = false
				}
			}
		}
		rst = append(rst, ' ')
	}
	rst = rst[:len(rst)-1]
	return string(rst)
}
