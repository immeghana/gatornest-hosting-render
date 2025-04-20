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
	Name           string             `json:"name" binding:"required" validate:"required"`
	Email          string             `json:"email" gorm:"unique;not null" binding:"required,email" validate:"required,email"`
	RoomID         *uint              `json:"room_id"`
	Phone          string             `json:"phone" binding:"required" validate:"required"`
	DormPreference string             `json:"dorm_preference" binding:"required" validate:"required"`
	Password       string             `json:"password" binding:"required" validate:"required"`
	Gender         string             `json:"gender"`
	Age            int                `json:"age"`
	Major          string             `json:"major"`
	LanguageSpoken string             `json:"language_spoken"`
	Preference     Preference         `json:"preference"`
	Cleanliness    Cleanliness        `json:"cleanliness"`
	FoodPref       FoodPreference     `json:"food_preference"`
	PeopleOver     PeopleOver         `json:"people_over"`
	LangPref       LanguagePreference `json:"language_preference"`
	Address        string             `json:"address"`
	University     string             `json:"university"`
	StudentID      string             `json:"student_id" gorm:"unique"`
	PendingPayment float64            `json:"pending_payment" gorm:"default:1250.00"`
	PendingDues    float64            `json:"pending_dues" gorm:"default:10"`
}
