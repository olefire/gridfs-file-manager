package utils

import "go.mongodb.org/mongo-driver/bson/primitive"

func ConvertToPrimitiveObjectId(stringArray []string) ([]primitive.ObjectID, error) {
	var primitiveSlice []primitive.ObjectID
	for _, s := range stringArray {
		id, err := primitive.ObjectIDFromHex(s)
		if err != nil {
			return nil, err
		}
		primitiveSlice = append(primitiveSlice, id)
	}
	return primitiveSlice, nil
}
