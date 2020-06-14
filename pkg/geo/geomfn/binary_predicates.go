// Copyright 2020 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package geomfn

import (
	"fmt"
	"math"

	"github.com/cockroachdb/cockroach/pkg/geo"
	"github.com/cockroachdb/cockroach/pkg/geo/geos"
	"github.com/twpayne/go-geom"
)

// Azimuth returns the azimuth in radians of the segment defined by the given point geometries.
// The azimuth is angle is referenced from north, and is positive clockwise.
// North = 0; East = π/2; South = π; West = 3π/2.
func Azimuth(a *geom.Point, b *geom.Point) (float64, error) {
	if a.X() == b.X() && a.Y() == b.Y() {
		return 0, fmt.Errorf("points are the same")
	}

	return math.Mod(2*math.Pi+math.Pi/2-math.Atan2(b.Y()-a.Y(), b.X()-a.X()), 2*math.Pi), nil
}

// Covers returns whether geometry A covers geometry B.
func Covers(a *geo.Geometry, b *geo.Geometry) (bool, error) {
	if a.SRID() != b.SRID() {
		return false, geo.NewMismatchingSRIDsError(a, b)
	}
	if !a.BoundingBoxIntersects(b) {
		return false, nil
	}
	return geos.Covers(a.EWKB(), b.EWKB())
}

// CoveredBy returns whether geometry A is covered by geometry B.
func CoveredBy(a *geo.Geometry, b *geo.Geometry) (bool, error) {
	if a.SRID() != b.SRID() {
		return false, geo.NewMismatchingSRIDsError(a, b)
	}
	if !a.BoundingBoxIntersects(b) {
		return false, nil
	}
	return geos.CoveredBy(a.EWKB(), b.EWKB())
}

// Contains returns whether geometry A contains geometry B.
func Contains(a *geo.Geometry, b *geo.Geometry) (bool, error) {
	if a.SRID() != b.SRID() {
		return false, geo.NewMismatchingSRIDsError(a, b)
	}
	if !a.BoundingBoxIntersects(b) {
		return false, nil
	}
	return geos.Contains(a.EWKB(), b.EWKB())
}

// ContainsProperly returns whether geometry A properly contains geometry B.
func ContainsProperly(a *geo.Geometry, b *geo.Geometry) (bool, error) {
	if a.SRID() != b.SRID() {
		return false, geo.NewMismatchingSRIDsError(a, b)
	}
	if !a.BoundingBoxIntersects(b) {
		return false, nil
	}
	return geos.RelatePattern(a.EWKB(), b.EWKB(), "T**FF*FF*")
}

// Crosses returns whether geometry A crosses geometry B.
func Crosses(a *geo.Geometry, b *geo.Geometry) (bool, error) {
	if a.SRID() != b.SRID() {
		return false, geo.NewMismatchingSRIDsError(a, b)
	}
	if !a.BoundingBoxIntersects(b) {
		return false, nil
	}
	return geos.Crosses(a.EWKB(), b.EWKB())
}

// Equals returns whether geometry A equals geometry B.
func Equals(a *geo.Geometry, b *geo.Geometry) (bool, error) {
	if a.SRID() != b.SRID() {
		return false, geo.NewMismatchingSRIDsError(a, b)
	}
	// Empty items are equal to each other.
	// Do this check before the BoundingBoxIntersects check, as we would otherwise
	// return false.
	if a.Empty() && b.Empty() {
		return true, nil
	}
	if !a.BoundingBoxIntersects(b) {
		return false, nil
	}
	return geos.Equals(a.EWKB(), b.EWKB())
}

// Intersects returns whether geometry A intersects geometry B.
func Intersects(a *geo.Geometry, b *geo.Geometry) (bool, error) {
	if a.SRID() != b.SRID() {
		return false, geo.NewMismatchingSRIDsError(a, b)
	}
	if !a.BoundingBoxIntersects(b) {
		return false, nil
	}
	return geos.Intersects(a.EWKB(), b.EWKB())
}

// Overlaps returns whether geometry A overlaps geometry B.
func Overlaps(a *geo.Geometry, b *geo.Geometry) (bool, error) {
	if a.SRID() != b.SRID() {
		return false, geo.NewMismatchingSRIDsError(a, b)
	}
	if !a.BoundingBoxIntersects(b) {
		return false, nil
	}
	return geos.Overlaps(a.EWKB(), b.EWKB())
}

// Touches returns whether geometry A touches geometry B.
func Touches(a *geo.Geometry, b *geo.Geometry) (bool, error) {
	if a.SRID() != b.SRID() {
		return false, geo.NewMismatchingSRIDsError(a, b)
	}
	if !a.BoundingBoxIntersects(b) {
		return false, nil
	}
	return geos.Touches(a.EWKB(), b.EWKB())
}

// Within returns whether geometry A is within geometry B.
func Within(a *geo.Geometry, b *geo.Geometry) (bool, error) {
	if a.SRID() != b.SRID() {
		return false, geo.NewMismatchingSRIDsError(a, b)
	}
	if !a.BoundingBoxIntersects(b) {
		return false, nil
	}
	return geos.Within(a.EWKB(), b.EWKB())
}
