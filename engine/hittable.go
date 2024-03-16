package engine

type HitRecord struct {
	t        float64  // which t generated the hit
	p        Point3   // which point when hit
	normal   Vec3     // normal at that point
	material Material // the material associated to this record
}

// Hittable defines the interface of objects that can be hit by a ray
type Hittable interface {
	hit(r *Ray, interval *Interval) (bool, *HitRecord)
}

// HittableList defines a simple list of hittable
type HittableList []Hittable

// hit defines the method for a list of hittables: will return the one closest
func (hl HittableList) hit(r *Ray, interval *Interval) (bool, *HitRecord) {
	var res *HitRecord
	hitAnything := false

	closestSoFar := interval.max

	for _, h := range hl {
		if hit, hr := h.hit(r, &Interval{interval.min, closestSoFar}); hit {
			hitAnything = true
			res = hr
			closestSoFar = hr.t
		}
	}

	return hitAnything, res
}