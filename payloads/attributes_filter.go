package payloads

import "category-rest-service/utils"

// AttributeFilters structure for query string parameters
type AttributeFilters struct {
	ID string   `bson:"ID"`
	Name string `bson:"name"`
}

// Clean - To clean empty fields 
func (af *AttributeFilters) Clean() (newEntity interface{}){
	i := map[string]interface{}{}

	if !utils.IsEmpty(af.ID) {
		i["ID"] = af.ID
	}

	if !utils.IsEmpty(af.Name) {
		i["name"] = af.Name
	}
	return i
}