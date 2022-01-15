package cuboid

// Non-overlapping Cuboid set

type CuboidSet struct {
	cuboids map[Cuboid]struct{}
}

func NewSet() CuboidSet {
	return CuboidSet{
		cuboids: make(map[Cuboid]struct{}),
	}
}

func (c *CuboidSet) Size() int { return len(c.cuboids) }

func overlapSplitter(referenceCuboid, cuboidToSplit Cuboid) []Cuboid {

	// Check for non-overlap
	if cuboidToSplit.normalizedMaxVertex.X <= referenceCuboid.normalizedMinVertex.X {
		return nil
	}
	if cuboidToSplit.normalizedMinVertex.X >= referenceCuboid.normalizedMaxVertex.X {
		return nil
	}
	if cuboidToSplit.normalizedMaxVertex.Y <= referenceCuboid.normalizedMinVertex.Y {
		return nil
	}
	if cuboidToSplit.normalizedMinVertex.Y >= referenceCuboid.normalizedMaxVertex.Y {
		return nil
	}
	if cuboidToSplit.normalizedMaxVertex.Z <= referenceCuboid.normalizedMinVertex.Z {
		return nil
	}
	if cuboidToSplit.normalizedMinVertex.Z >= referenceCuboid.normalizedMaxVertex.Z {
		return nil
	}

	// Check overlaps x min and split if overlapped
	if cuboidToSplit.normalizedMinVertex.X < referenceCuboid.normalizedMinVertex.X &&
		cuboidToSplit.normalizedMaxVertex.X > referenceCuboid.normalizedMinVertex.X {

		splitCuboids := cuboidToSplit.splitX(referenceCuboid.normalizedMinVertex.X)
		return splitCuboids
	}

	// Check overlaps x max and split if overlapped
	if cuboidToSplit.normalizedMinVertex.X < referenceCuboid.normalizedMaxVertex.X &&
		cuboidToSplit.normalizedMaxVertex.X > referenceCuboid.normalizedMaxVertex.X {

		splitCuboids := cuboidToSplit.splitX(referenceCuboid.normalizedMaxVertex.X)
		return splitCuboids
	}

	// Check overlaps y min and split if overlapped
	if cuboidToSplit.normalizedMinVertex.Y < referenceCuboid.normalizedMinVertex.Y &&
		cuboidToSplit.normalizedMaxVertex.Y > referenceCuboid.normalizedMinVertex.Y {

		splitCuboids := cuboidToSplit.splitY(referenceCuboid.normalizedMinVertex.Y)
		return splitCuboids
	}

	// Check overlaps y max and split if overlapped
	if cuboidToSplit.normalizedMinVertex.Y < referenceCuboid.normalizedMaxVertex.Y &&
		cuboidToSplit.normalizedMaxVertex.Y > referenceCuboid.normalizedMaxVertex.Y {

		splitCuboids := cuboidToSplit.splitY(referenceCuboid.normalizedMaxVertex.Y)
		return splitCuboids
	}

	// Check overlaps z min and split if overlapped
	if cuboidToSplit.normalizedMinVertex.Z < referenceCuboid.normalizedMinVertex.Z &&
		cuboidToSplit.normalizedMaxVertex.Z > referenceCuboid.normalizedMinVertex.Z {

		splitCuboids := cuboidToSplit.splitZ(referenceCuboid.normalizedMinVertex.Z)
		return splitCuboids
	}

	// Check overlaps z max and split if overlapped
	if cuboidToSplit.normalizedMinVertex.Z < referenceCuboid.normalizedMaxVertex.Z &&
		cuboidToSplit.normalizedMaxVertex.Z > referenceCuboid.normalizedMaxVertex.Z {

		splitCuboids := cuboidToSplit.splitZ(referenceCuboid.normalizedMaxVertex.Z)
		return splitCuboids
	}
	return nil
}

func (set *CuboidSet) Add(c Cuboid) {
	// workList contains a set of Cuboids that need to be checked for
	// overlap and added.  If we pop the head and find that it is overlapped
	// with an existing cuboid in the set then we split it and push the
	// sub-cuboids back onto the list to be re-checked.

	workList := make([]Cuboid, 0)
	workList = append(workList, c)

WORKLOOP:
	for len(workList) != 0 {
		newCuboid := workList[0]
		workList = workList[1:]

		for existingCuboid := range set.cuboids {
			if existingCuboid.Contains(newCuboid) {
				// The points contained in newCuboid are already
				// covered so we can drop it.
				continue WORKLOOP
			}

			// Check for overlap, split if overlapped, and add
			// sub-cuboids to work list for re-checking.
			splitCuboids := overlapSplitter(existingCuboid, newCuboid)
			if splitCuboids != nil {
				workList = append(workList, splitCuboids...)
				continue WORKLOOP
			}
		}

		// The cuboid doesn't overlap any other cuboids in the set so add it.
		set.cuboids[newCuboid] = struct{}{}
	}
}

func (set *CuboidSet) Remove(cuboidToRemove Cuboid) {

RESTART:
	for {
		for existingCuboid := range set.cuboids {
			// Remove the existing cuboid if it is completly
			// contained in the remove cuboid
			if cuboidToRemove.Contains(existingCuboid) {
				delete(set.cuboids, existingCuboid)
				continue RESTART // NB.  Continuing the current loop also works here.
			}

			// Call the splitter to see if there is an overlap of
			// the existing cuboid and the cuboid to remove.  If the
			// splitter finds an overlap, it will return split it
			// into 2 cuboids and replace the cuboid with the 2 new
			// sub-cuboids.
			//
			// The splitter will only perform one overlapping split
			// so if there is a multidimensional overlap, we will
			// find it on a subsequent iteration and perform another
			// split.
			//
			// Eventually, the sub-cuboids will be split in a way
			// that they are fully contained or fully outside the
			// delete region and the logic to handle those cases
			// will take care of them.
			//
			// If the splitter doesn't find an overlap then there is
			// nothing to delete and we are done with the existing
			// cuboid.  We could optimize the algorithm to not ever
			// check it again but we don't currently do that.
			//
			splitCuboids := overlapSplitter(cuboidToRemove, existingCuboid)
			if splitCuboids != nil {
				delete(set.cuboids, existingCuboid)

				for _, c := range splitCuboids {
					set.cuboids[c] = struct{}{}
				}

				continue RESTART
			}
		}

		break
	}
}

func (set *CuboidSet) Volume() int {
	volume := 0

	for c := range set.cuboids {
		volume += c.Volume()
	}

	return volume
}
