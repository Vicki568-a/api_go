package taxi



type  Cars struct {
	ID				string 	`json:"id" bson:"_id,omitempty"`
	StateNumber		string 	`json:"state_number" bson:"state_number"`
	Photo			string 	`json:"photo" bson:"photo"`
	Brand			string 	`json:"brand" bson:"brand"`
	Сolor			string 	`json:"сolor" bson:"сolor"`
	Passenger		int	   	`json:"passenger" bson:"_passenger,omitempty"`
	Driver			int		`json:"driver" bson:"_driver,omitempty"`
}
