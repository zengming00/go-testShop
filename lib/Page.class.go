package lib

import "strings"

func str_replace(search []string, replace []string, str string) string {
	if len(search) != len(replace) {
		panic("len(search) != len(replace)")
	}
	var oldnew = make([]string, len(search)+len(replace))
	for i, v := range search {
		oldnew[i*2] = v
		oldnew[i*2+1] = replace[i]
	}
	r := strings.NewReplacer(oldnew...)
	return r.Replace(str)
}
