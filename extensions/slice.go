//Copyright Copyright 2018 Andrey Z.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0

package extensions


func Min(x []int) int {
	if len(x) == 0 {
		return 0
	}
	ans := x[0]
	for _, v := range x[1:] {
		if v < ans {
			ans = v
		}
	}
	return ans
}
