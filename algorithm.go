package main

import (
	"slices"
	"sort"
)

func ComputeZoo(numEnclosures int, animals []Animal) (Zoo, error) {
	enclosures := []Enclosure{}

	// Sorting for greedy algorithm. Some trial and error could improve this sort
	// Interesting different versions of this sort gives different results
	sort.SliceStable(animals, func(i, j int) bool {

		return animals[i].Popularity + 3*animals[i].MaximumFriends > animals[j].Popularity + 3*animals[j].MaximumFriends

		// return animals[i].Popularity > animals[j].Popularity

		// if animals[i].Popularity != animals[j].Popularity {
		// 	return animals[i].Popularity > animals[j].Popularity
		// }
		// if animals[i].MaximumFriends != animals[j].MaximumFriends {
		// 	return animals[i].MaximumFriends > animals[j].MaximumFriends
		// }
		// return len(animals[i].Incompatible) < len(animals[j].Incompatible)
	})

	for i := 0; i < numEnclosures; i++ {
		enclosures = append(enclosures, Enclosure{Animals: []Animal{}, MaximumFriends: 100, Incompatible: []string{}})
	}

	totalAnimals := 0
	totalPopularity := 0

	for i := 0; i < len(animals); i++ {
		for j := 0; j < numEnclosures; j++ {
			if enclosures[j].MaximumFriends > len(enclosures[j].Animals) &&
				animals[i].MaximumFriends >= len(enclosures[j].Animals) {

				animalCompatible := true
				for m := 0; m < len(enclosures[j].Animals); m++ {
					if slices.Contains(animals[i].Incompatible, enclosures[j].Animals[m].Type) {
						animalCompatible = false
					}
				}

				if slices.Contains(enclosures[j].Incompatible, animals[i].Type) {
					animalCompatible = false
				}

				if !animalCompatible {
					continue
				}

				// Check if it adds any imcompatibilities ?

				enclosures[j].Animals = append(enclosures[j].Animals, animals[i])
				enclosures[j].MaximumFriends = min(enclosures[j].MaximumFriends, animals[i].MaximumFriends)

				totalAnimals += 1
				totalPopularity += animals[i].Popularity

				newEnclosureIncompatibles := []string{}
				for k := 0; k < len(enclosures[j].Animals); k++ {
					newEnclosureIncompatibles = slices.Concat(newEnclosureIncompatibles, enclosures[j].Animals[k].Incompatible)
					slices.Sort(newEnclosureIncompatibles)
					newEnclosureIncompatibles = slices.Compact(newEnclosureIncompatibles)
				}
				enclosures[j].Incompatible = newEnclosureIncompatibles

				break
			}
		}

	}

	zoo := Zoo{TotalAnimals: totalAnimals, TotalPopularity: totalPopularity, Enclosures: enclosures}
	return zoo, nil
}
