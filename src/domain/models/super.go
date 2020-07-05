package models

// Super entity represents a Villain or Hero
//
// Base            - Field that represent the Base entity
// Name            - A field that represent the name how super
// FullName        - Represent the personal name
// Intelligence    - A field that represent the super intelligence
// Power           - A field that represent the super intelligence
// Occupation      - Represent the personal super occupation
// Image           - A field with url of hero image
// Alignment 	   - A field that represent if super its hero or villain
// RelativesAmount - Represents the number of hero relatives
// Relatives       - Represents the hero relatives names
//
type Super struct {
	Base
	Name             string   `gorm:"size:255;not null" json:"name"`
	FullName         string   `gorm:"size:255;not null" json:"fullname"`
	Intelligence     int      `gorm:"not null" json:"intelligence"`
	Power            int      `gorm:"not null" json:"power"`
	Occupation       string   `gorm:"size:255" json:"occupation"`
	Image            string   `gorm:"size:750" json:"image"`
	Alignment        string   `gorm:"size:255" json:"alignment"`
	RelativesAmount  int      `json:"relativesAmount"`
	Relatives        string   `json:"relatives"`
	GroupsAssociated []*Group `gorm:"many2many:group_super" json:"groupsAssociated"`
}
