package engine

import (
	"encoding/json"
	"math"

	"github.com/ath0m/DistributedRaytracer/agent/engine/geometry"
	"github.com/ath0m/DistributedRaytracer/agent/engine/utils"
)

type Sphere struct {
	Center   geometry.Point3 `json:"center"`
	Radius   float64         `json:"radius"`
	Material Material        `json:"material"`
}

// UnmarshalJSON unmarshals JSON data into a Sphere object
func (s *Sphere) UnmarshalJSON(data []byte) error {
	type Alias Sphere
	aux := &struct {
		Material json.RawMessage `json:"material"`
		*Alias
	}{
		Alias: (*Alias)(s),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if aux.Material != nil {
		material, err := UnmarshalMaterial(aux.Material)
		if err != nil {
			return err
		}
		s.Material = material
	}
	return nil
}

// Hit implements the Hit interface for a Sphere
func (s Sphere) Hit(r *geometry.Ray, interval *utils.Interval) (bool, *HitRecord) {
	oc := r.Origin.Sub(s.Center)
	a := r.Direction.LengthSq()
	b := geometry.Dot(oc, r.Direction)
	c := oc.LengthSq() - s.Radius*s.Radius
	discriminant := b*b - a*c

	if discriminant < 0 {
		return false, nil
	}
	discriminantSquareRoot := math.Sqrt(discriminant)

	root := (-b - discriminantSquareRoot) / a
	if !interval.Surrounds(root) {
		root = (-b + discriminantSquareRoot) / a
		if !interval.Surrounds(root) {
			return false, nil
		}
	}

	hitPoint := r.PointAt(root)
	hr := HitRecord{
		T:        root,
		P:        hitPoint,
		Normal:   hitPoint.Sub(s.Center).Scale(1 / s.Radius),
		Material: s.Material,
	}
	return true, &hr
}
