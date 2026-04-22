package models

type UserRole string

const (
	RoleParent     UserRole = "parent"
	RoleSpecialist UserRole = "specialist"
	RoleAdmin      UserRole = "admin"
)

type AppointmentStatus string

const (
	StatusPending  AppointmentStatus = "pending"
	StatusApproved AppointmentStatus = "approved"
	StatusRejected AppointmentStatus = "rejected"
)

type RecommendationType string

const (
	RecommendationChild  RecommendationType = "child"
	RecommendationParent RecommendationType = "parent"
)
