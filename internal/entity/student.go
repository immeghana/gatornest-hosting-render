package entity

import "time"

type Preference string
type Cleanliness string
type FoodPreference string
type PeopleOver string
type LanguagePreference string

const (
	EarlyBird        Preference         = "Early Bird"
	NightOwl         Preference         = "Night Owl"
	VeryTidy         Cleanliness        = "Very Tidy"
	ModeratelyTidy   Cleanliness        = "Moderately Tidy"
	Messy            Cleanliness        = "Messy"
	NoGuests         PeopleOver         = "Private Space"
	OccasionalGuests PeopleOver         = "Occasional Visitors"
	AlwaysGuests     PeopleOver         = "Social Space"
	Veg              FoodPreference     = "Veg"
	NonVeg           FoodPreference     = "Non-Veg"
	SameLanguage     LanguagePreference = "Yes"
	NoPreference     LanguagePreference = "No"
	EitherWay        LanguagePreference = "Either one is fine"
)

type Student struct {
	ID             uint               `gorm:"primaryKey" json:"id"`
	CreatedAt      time.Time          `json:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at"`
	DeletedAt      *time.Time         `gorm:"index" json:"deleted_at,omitempty"`
	Name           string             `json:"name" binding:"required"`
	Email          string             `json:"email" gorm:"unique;not null" binding:"required,email"`
	RoomID         *uint              `json:"room_id"`
	Phone          string             `json:"phone,omitempty" binding:"required"`
	DormPreference string             `json:"dorm_preference,omitempty" binding:"required"`
	Password       string             `json:"password" binding:"required,min=6"`
	Gender         string             `json:"gender,omitempty" binding:"required"`
	Age            int                `json:"age,omitempty" binding:"required,min=18"`
	Major          string             `json:"major,omitempty" binding:"required"`
	LanguageSpoken string             `json:"language_spoken,omitempty" binding:"required"`
	Preference     Preference         `json:"preference,omitempty" binding:"required"`
	Cleanliness    Cleanliness        `json:"cleanliness,omitempty" binding:"required"`
	FoodPref       FoodPreference     `json:"food_preference,omitempty" binding:"required"`
	PeopleOver     PeopleOver         `json:"people_over,omitempty" binding:"required"`
	LangPref       LanguagePreference `json:"language_preference,omitempty" binding:"required"`
	Address        string             `json:"address" binding:"required"`
	University     string             `json:"university" binding:"required"`
	StudentID      string             `json:"student_id" gorm:"unique;not null" binding:"required"`
	PendingPayment float64            `json:"pending_payment" gorm:"default:1250.00"`
	PendingDues    float64            `json:"pending_dues" gorm:"default:10"`
}
