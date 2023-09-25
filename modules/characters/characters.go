package characters

type Character struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Health  int    `json:"health"`
	Attack  int    `json:"attack"`
	Defense int    `json:"defense"`
}
