package models

import "go.mongodb.org/mongo-driver/v2/bson"

type Todo struct {
	Id      bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Label   string        `json:"label" bson:"label"`
	Checked bool          `json:"checked" bson:"checked"`
	Date    string        `json:"date" bson:"date"`
}

type CreateTodoDTO struct {
	Label   string `json:"label"`
	Checked bool   `json:"checked"`
}

// var todos = []Todo{
// 	{Id: "571e5eb7-f253-4d50-b008-937ada01586e", Label: "Blue Train", Checked: false, Date: time.Now().Format(time.RFC3339)},
// 	{Id: "7bbd5369-fd5a-4ccc-ae8a-a974e420a7a4", Label: "Jeru", Checked: false, Date: time.Now().Format(time.RFC3339)},
// 	{Id: "3aff95e2-3dbe-4443-860b-40641b38557b", Label: "Sarah Vaughan and Clifford Brown", Checked: false, Date: time.Now().Format(time.RFC3339)},
// }
