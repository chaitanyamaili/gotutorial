package main

import (
	"fmt"
	"sort"
)

func main() {
	// declare/init and allocate the memory to maps
	states := make(map[string]string)
	fmt.Println(states)

	// assign values in to the map
	states["MH"] = "Maharastra"
	states["KA"] = "Karnataka"
	states["DL"] = "Delhi"
	states["GJ"] = "Gujrat"
	fmt.Println(states)

	// fetch the value from the map
	fmt.Println("KA value is", states["KA"])

	// delete the value from the map
	delete(states, "GJ")
	fmt.Println(states)

	states["KL"] = "Kerala"
	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)
	}

	// as the iteration over the map is not consistent
	// we will take the help of slice
	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	// sort in ASC
	//sort.Strings(keys)
	// sort in DESC
	sort.Sort(sort.Reverse(sort.StringSlice(keys)))
	fmt.Println("States keys is", keys)

	fmt.Println("\nSorted")
	for i := range keys {
		fmt.Println(states[keys[i]], keys[i])
	}
}
