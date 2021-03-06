/*
Package geojson is a library for encoding and decoding GeoJSON into Go structs using
the geometries in the orb package. Supports both the json.Marshaler and json.Unmarshaler
interfaces as well as helper functions such as `UnmarshalFeatureCollection` and `UnmarshalFeature`.
*/
package geojson

import (
	"encoding/json"
)

// A FeatureCollection correlates to a GeoJSON feature collection.
type FeatureCollection struct {
	Type     string     `json:"type"`
	Features []*Feature `json:"features"`
}

// NewFeatureCollection creates and initializes a new feature collection.
func NewFeatureCollection() *FeatureCollection {
	return &FeatureCollection{
		Type:     "FeatureCollection",
		Features: []*Feature{},
	}
}

// Append appends a feature to the collection.
func (fc *FeatureCollection) Append(feature *Feature) *FeatureCollection {
	fc.Features = append(fc.Features, feature)
	return fc
}

// MarshalJSON converts the feature collection object into the proper JSON.
// It will handle the encoding of all the child features and geometries.
// Alternately one can call json.Marshal(fc) directly for the same result.
func (fc FeatureCollection) MarshalJSON() ([]byte, error) {
	type featureCollection FeatureCollection

	c := featureCollection{
		Type:     "FeatureCollection",
		Features: fc.Features,
	}

	if c.Features == nil {
		c.Features = []*Feature{}
	}
	return json.Marshal(c)
}

// UnmarshalFeatureCollection decodes the data into a GeoJSON feature collection.
// Alternately one can call json.Unmarshal(fc) directly for the same result.
func UnmarshalFeatureCollection(data []byte) (*FeatureCollection, error) {
	fc := &FeatureCollection{}
	err := json.Unmarshal(data, fc)
	if err != nil {
		return nil, err
	}

	return fc, nil
}
