package model

type Agent struct {
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Matricule string `gorm:"size:100;not null;unique" json:"matricule"`
	Nom       string `gorm:"size:100;not null" json:"nom"`
	Prenoms   string `gorm:"size:100;not null" json:"prenoms"`
	Poste     string `gorm:"size:100" json:"poste"`
	Agence    string `gorm:"size:100" json:"agence"`
	Region    string `gorm:"size:100" json:"region"`
	Email     string `gorm:"size:150;unique" json:"email"`
}

func (Agent) TableName() string {
	return "agents"
}