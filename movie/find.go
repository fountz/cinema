package movie

func FindMovieName(imbID string) string {
	switch imbID {
	case "tt4154796":
		return "Avengers: End Game"
	case "tt1825683":
		return "Black Panther"
	}
	return "not found"
}
