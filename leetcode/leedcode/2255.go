package leedcode

func countPrefixes(words []string, s string) int {
    counter := 0

    for _, y := range words {

        checker := true

        if len(y) > len(s) {
            checker = false
        } else {
            for a, b := range y {
                if rune(s[a]) != b {
                    checker = false

                    break
                }
            }
        }

        if checker {
            counter++
        }
    }

    return counter
}