package schema

type QueueUsers struct {
	ID   uint   // adds ID, CreatedAt, UpdatedAt, DeletedAt fields
	Name string `json:"name"`
}
