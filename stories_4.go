package main

func Stories4() string {
	var name string = ""

	for i := 0; i < len(Periods); i++ {
		if name != "" {
			name += ", "
		}
		name += Periods[i].Winner.Name
	}

	return name
}