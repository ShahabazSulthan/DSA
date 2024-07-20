package main

func LongCommonPrefix(strs []string) string {
	if len(strs) == 0 {
        return ""
    }
    prefix := strs[0]
    for _, str := range strs[1:] {
        for len(str) < len(prefix) || prefix != str[:len(prefix)] {
            if len(prefix) == 0 {
                return ""
            }
            prefix = prefix[:len(prefix)-1]
        }
    }
    return prefix
}
