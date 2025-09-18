package main

func in_str(item string, list []string) bool {
	for _, i := range list {
		if item == i {
			return true
		}
	}
	return false
}

func int_list_to_strings(l []interface{}) []string {
	lr := make([]string, len(l))
	for i, _ := range l {
		lr[i] = l[i].(string)
	}
	return lr
}
