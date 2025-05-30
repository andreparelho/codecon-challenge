package repository

type User struct {
	Id      string `json:"id"`
	Name    string `json:"nome"`
	Age     int    `json:"idade"`
	Score   int    `json:"score"`
	Active  bool   `json:"ativo"`
	Country string `json:"pais"`
	Team    Team   `json:"equipe"`
	Logs    []Logs `json:"logs"`
}

type Team struct {
	Name     string     `json:"nome"`
	Leader   bool       `json:"lider"`
	Projects []Projects `json:"projetos"`
}

type Projects struct {
	Name      string `json:"nome"`
	Completed bool   `json:"concluido"`
}

type Logs struct {
	Date   string `json:"data"`
	Action string `json:"acao"`
}

type Countries struct {
	Country string `json:"country"`
	Total   int    `json:"total"`
}

type ActiveUsers struct {
	Date  string `json:"date"`
	Total int    `json:"total"`
}

type TeamInsights struct {
	Team              string `json:"team"`
	TotalMembers      int    `json:"total_members"`
	Leaders           int    `json:"leaders"`
	CompletedProjects int    `json:"completed_projects"`
	ActivePercentage  int    `json:"active_percentage"`
}
