package main

import "fmt"

func main() {
	colors := map[string]string{
		"red": "I am color red",
		"blue": "Le Mans Blue",
	}
	fmt.Println(colors);

	// wrong way to initialise map
	var bmws map[int]string
	fmt.Println(bmws);
	// bmws[1] = "7 Series" // throws error
	// bmws[7] = "5 Series" // panic: assignment to entry in nil map

	// correct way to initialise empty map
	var mercs = make(map[string]int)
	fmt.Println(mercs)
	mercs["E Class"] = 211;
	mercs["E Class Old"] = 210;
	fmt.Println(mercs)
	delete(mercs, "E Class Old");
	fmt.Println(mercs)

}