package structures

type JsonResponseSpec struct {
	EndpointURI string `gorm:"primaryKey"`
	Value       string
}
