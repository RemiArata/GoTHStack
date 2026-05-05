package models

type Country struct {
	Name        string
	Capital     string
	Population  int
	Description string
}

var Countries = []Country{
	{
		Name:        "USA",
		Capital:     "Washington DC",
		Population:  349_000_000,
		Description: "Land of the free, home of the brave",
	},
	{
		Name:        "Japan",
		Capital:     "Tokyo",
		Population:  123_000_000,
		Description: "The cherry blossoms bloom every spring",
	},
	{
		Name:        "South Africa",
		Capital:     "Pretoria",
		Population:  65_000_000,
		Description: "The home of Nelson Mandela",
	},
}
