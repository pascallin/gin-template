package posts

import "time"

type Post struct {
	UserID 		string 		`bson:"userId" json:"userId"`
	Title 		string 		`bson:"title" json:"title"`
	Content 	string		`bson:"content" json:"content"`
	Tags		[]string	`bson:"tags" json:"tags"`
	CreatedAt 	time.Time 	`bson:"createdAt" json:"createdAt"`
	UpdatedAt 	time.Time	`bson:"updatedAt" json:"updatedAt"`
}

type Tag struct {
	UserID		string 		`bson:"userId" json:"userId"`
	Values 		[]string	`bson:"values" json:"values"`
}