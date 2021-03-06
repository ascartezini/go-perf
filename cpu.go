package main

import (
	"io"
)

func readbyte(r io.Reader) (rune, error) {
	var buf [1]byte
	_, err := r.Read(buf[:])
	return rune(buf[0]), err
}

// func main() {
// 	defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
// 	// defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()

// 	f, err := os.Open(os.Args[1])
// 	if err != nil {
// 		log.Fatalf("could not open file %q: %v", os.Args[1], err)
// 	}

// 	words := 0
// 	inword := false
// 	b := bufio.NewReader(f)
// 	for {
// 		r, err := readbyte(b)
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			log.Fatalf("could not read file %q: %v", os.Args[1], err)
// 		}
// 		if unicode.IsSpace(r) && inword {
// 			words++
// 			inword = false
// 		}
// 		inword = unicode.IsLetter(r)
// 	}
// 	fmt.Printf("%q: %d words\n", os.Args[1], words)
// }
