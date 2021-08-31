package models

import (
    "time"
)

type Post struct {
    ID *string `json:"id,omitempty" bson:"_id,omitempty"`
    Title *string `json:"title"`
    Author *string `json:"author"`
    Tags *string `json:"tags"`
    Content *string `json:"content"`
    Read int `json:"read"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
}
