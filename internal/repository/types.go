package repository

type User struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Score   int    `json:"score"`
	Active  bool   `json:"active"`
	Country string `json:"country"`
	Team    Team   `json:"team"`
	Logs    Logs   `json:"logs"`
}

type Team struct {
	Name     string   `json:"name"`
	Leader   string   `json:"leader"`
	Projects Projects `json:"projects"`
}

type Projects struct {
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

type Logs struct {
	Date   string `json:"date"`
	Action string `json:"action"`
}
