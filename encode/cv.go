package encode

type Curriculum struct {
	Person   Person    `json:"person" yaml:"person"`
	Contact  Contact   `json:"contact" yaml:"contact"`
	Sections []Section `json:"sections" yaml:"sections"`
}

type Person struct {
	Name string `json:"name" yaml:"name"`
	Role string `json:"role" yaml:"role"`
}

type Contact struct {
	Email    string `json:"email" yaml:"email"`
	Location string `json:"location" yaml:"location"`
	LinkedIn string `json:"linkedin" yaml:"linkedin"`
}

type Section struct {
	Title    string    `json:"title" yaml:"title"`
	Articles []Article `json:"articles" yaml:"articles"`
}

type Article struct {
	What     string   `json:"what" yaml:"what"`
	Where    string   `json:"where" yaml:"where"`
	When     string   `json:"when" yaml:"when"`
	Desc     string   `json:"desc" yaml:"desc"`
	List     []string `json:"list" yaml:"list"`
	FullList []string `json:"full-list" yaml:"full-list"`
}
