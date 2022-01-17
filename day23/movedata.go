package main

// Map of legal moves from a location.

// We don't include moves from room to room.  They can be formed by a room to
// hall and then a hall to room sequence.
var moveData = map[location]locationInfo{
	roomA1: {
		occupant:  amberAmphipod,
		peerRooms: []location{roomA1, roomA2, roomA3, roomA4},
		moves: []move{
			{
				to:       hall1,
				path:     []location{hall2, hall1},
				distance: 3,
			},
			{
				to:       hall2,
				path:     []location{hall2},
				distance: 2,
			},
			{
				to:       hall3,
				path:     []location{hall3},
				distance: 2,
			},
			{
				to:       hall4,
				path:     []location{hall3, hall4},
				distance: 4,
			},
			{
				to:       hall5,
				path:     []location{hall3, hall4, hall5},
				distance: 6,
			},
			{
				to:       hall6,
				path:     []location{hall3, hall4, hall5, hall6},
				distance: 8,
			},
			{
				to:       hall7,
				path:     []location{hall3, hall4, hall5, hall6, hall7},
				distance: 9,
			},
		},
	},
	roomA2: {
		occupant:  amberAmphipod,
		peerRooms: []location{roomA1, roomA2, roomA3, roomA4},
		moves: []move{
			{
				to:       hall1,
				path:     []location{roomA1, hall2, hall1},
				distance: 4,
			},
			{
				to:       hall2,
				path:     []location{roomA1, hall2},
				distance: 3,
			},
			{
				to:       hall3,
				path:     []location{roomA1, hall3},
				distance: 3,
			},
			{
				to:       hall4,
				path:     []location{roomA1, hall3, hall4},
				distance: 5,
			},
			{
				to:       hall5,
				path:     []location{roomA1, hall3, hall4, hall5},
				distance: 7,
			},
			{
				to:       hall6,
				path:     []location{roomA1, hall3, hall4, hall5, hall6},
				distance: 9,
			},
			{
				to:       hall7,
				path:     []location{roomA1, hall3, hall4, hall5, hall6, hall7},
				distance: 10,
			},
		},
	},
	roomA3: {
		occupant:  amberAmphipod,
		peerRooms: []location{roomA1, roomA2, roomA3, roomA4},
		moves: []move{
			{
				to:       hall1,
				path:     []location{roomA2, roomA1, hall2, hall1},
				distance: 5,
			},
			{
				to:       hall2,
				path:     []location{roomA2, roomA1, hall2},
				distance: 4,
			},
			{
				to:       hall3,
				path:     []location{roomA2, roomA1, hall3},
				distance: 4,
			},
			{
				to:       hall4,
				path:     []location{roomA2, roomA1, hall3, hall4},
				distance: 6,
			},
			{
				to:       hall5,
				path:     []location{roomA2, roomA1, hall3, hall4, hall5},
				distance: 8,
			},
			{
				to:       hall6,
				path:     []location{roomA2, roomA1, hall3, hall4, hall5, hall6},
				distance: 10,
			},
			{
				to:       hall7,
				path:     []location{roomA2, roomA1, hall3, hall4, hall5, hall6, hall7},
				distance: 11,
			},
		},
	},
	roomA4: {
		occupant:  amberAmphipod,
		peerRooms: []location{roomA1, roomA2, roomA3, roomA4},
		moves: []move{
			{
				to:       hall1,
				path:     []location{roomA3, roomA2, roomA1, hall2, hall1},
				distance: 6,
			},
			{
				to:       hall2,
				path:     []location{roomA3, roomA2, roomA1, hall2},
				distance: 5,
			},
			{
				to:       hall3,
				path:     []location{roomA3, roomA2, roomA1, hall3},
				distance: 5,
			},
			{
				to:       hall4,
				path:     []location{roomA3, roomA2, roomA1, hall3, hall4},
				distance: 7,
			},
			{
				to:       hall5,
				path:     []location{roomA3, roomA2, roomA1, hall3, hall4, hall5},
				distance: 9,
			},
			{
				to:       hall6,
				path:     []location{roomA3, roomA2, roomA1, hall3, hall4, hall5, hall6},
				distance: 11,
			},
			{
				to:       hall7,
				path:     []location{roomA3, roomA2, roomA1, hall3, hall4, hall5, hall6, hall7},
				distance: 12,
			},
		},
	},
	roomB1: {
		occupant:  bronzeAmphipod,
		peerRooms: []location{roomB1, roomB2, roomB3, roomB4},
		moves: []move{
			{
				to:       hall1,
				path:     []location{hall3, hall2, hall1},
				distance: 5,
			},
			{
				to:       hall2,
				path:     []location{hall3, hall2},
				distance: 4,
			},
			{
				to:       hall3,
				path:     []location{hall3},
				distance: 2,
			},
			{
				to:       hall4,
				path:     []location{hall4},
				distance: 2,
			},
			{
				to:       hall5,
				path:     []location{hall4, hall5},
				distance: 4,
			},
			{
				to:       hall6,
				path:     []location{hall4, hall5, hall6},
				distance: 6,
			},
			{
				to:       hall7,
				path:     []location{hall4, hall5, hall6, hall7},
				distance: 7,
			},
		},
	},
	roomB2: {
		occupant:  bronzeAmphipod,
		peerRooms: []location{roomB1, roomB2, roomB3, roomB4},
		moves: []move{
			{
				to:       hall1,
				path:     []location{roomB1, hall3, hall2, hall1},
				distance: 6,
			},
			{
				to:       hall2,
				path:     []location{roomB1, hall3, hall2},
				distance: 5,
			},
			{
				to:       hall3,
				path:     []location{roomB1, hall3},
				distance: 3,
			},
			{
				to:       hall4,
				path:     []location{roomB1, hall4},
				distance: 3,
			},
			{
				to:       hall5,
				path:     []location{roomB1, hall4, hall5},
				distance: 5,
			},
			{
				to:       hall6,
				path:     []location{roomB1, hall4, hall5, hall6},
				distance: 7,
			},
			{
				to:       hall7,
				path:     []location{roomB1, hall4, hall5, hall6, hall7},
				distance: 8,
			},
		},
	},
	roomB3: {
		occupant:  bronzeAmphipod,
		peerRooms: []location{roomB1, roomB2, roomB3, roomB4},
		moves: []move{
			{
				to:       hall1,
				path:     []location{roomB2, roomB1, hall3, hall2, hall1},
				distance: 7,
			},
			{
				to:       hall2,
				path:     []location{roomB2, roomB1, hall3, hall2},
				distance: 6,
			},
			{
				to:       hall3,
				path:     []location{roomB2, roomB1, hall3},
				distance: 4,
			},
			{
				to:       hall4,
				path:     []location{roomB2, roomB1, hall4},
				distance: 4,
			},
			{
				to:       hall5,
				path:     []location{roomB2, roomB1, hall4, hall5},
				distance: 6,
			},
			{
				to:       hall6,
				path:     []location{roomB2, roomB1, hall4, hall5, hall6},
				distance: 8,
			},
			{
				to:       hall7,
				path:     []location{roomB2, roomB1, hall4, hall5, hall6, hall7},
				distance: 9,
			},
		},
	},
	roomB4: {
		occupant:  bronzeAmphipod,
		peerRooms: []location{roomB1, roomB2, roomB3, roomB4},
		moves: []move{
			{
				to:       hall1,
				path:     []location{roomB3, roomB2, roomB1, hall3, hall2, hall1},
				distance: 8,
			},
			{
				to:       hall2,
				path:     []location{roomB3, roomB2, roomB1, hall3, hall2},
				distance: 7,
			},
			{
				to:       hall3,
				path:     []location{roomB3, roomB2, roomB1, hall3},
				distance: 5,
			},
			{
				to:       hall4,
				path:     []location{roomB3, roomB2, roomB1, hall4},
				distance: 5,
			},
			{
				to:       hall5,
				path:     []location{roomB3, roomB2, roomB1, hall4, hall5},
				distance: 7,
			},
			{
				to:       hall6,
				path:     []location{roomB3, roomB2, roomB1, hall4, hall5, hall6},
				distance: 9,
			},
			{
				to:       hall7,
				path:     []location{roomB3, roomB2, roomB1, hall4, hall5, hall6, hall7},
				distance: 10,
			},
		},
	},
	roomC1: {
		occupant:  copperAmphipod,
		peerRooms: []location{roomC1, roomC2, roomC3, roomC4},
		moves: []move{
			{
				to:       hall1,
				path:     []location{hall4, hall3, hall2, hall1},
				distance: 7,
			},
			{
				to:       hall2,
				path:     []location{hall4, hall3, hall2},
				distance: 6,
			},
			{
				to:       hall3,
				path:     []location{hall4, hall3},
				distance: 4,
			},
			{
				to:       hall4,
				path:     []location{hall4},
				distance: 2,
			},
			{
				to:       hall5,
				path:     []location{hall5},
				distance: 2,
			},
			{
				to:       hall6,
				path:     []location{hall5, hall6},
				distance: 4,
			},
			{
				to:       hall7,
				path:     []location{hall5, hall6, hall7},
				distance: 5,
			},
		},
	},
	roomC2: {
		occupant:  copperAmphipod,
		peerRooms: []location{roomC1, roomC2, roomC3, roomC4},
		moves: []move{
			{
				to:       hall1,
				path:     []location{roomC1, hall4, hall3, hall2, hall1},
				distance: 8,
			},
			{
				to:       hall2,
				path:     []location{roomC1, hall4, hall3, hall2},
				distance: 7,
			},
			{
				to:       hall3,
				path:     []location{roomC1, hall4, hall3},
				distance: 5,
			},
			{
				to:       hall4,
				path:     []location{roomC1, hall4},
				distance: 3,
			},
			{
				to:       hall5,
				path:     []location{roomC1, hall5},
				distance: 3,
			},
			{
				to:       hall6,
				path:     []location{roomC1, hall5, hall6},
				distance: 5,
			},
			{
				to:       hall7,
				path:     []location{roomC1, hall5, hall6, hall7},
				distance: 6,
			},
		},
	},
	roomC3: {
		occupant:  copperAmphipod,
		peerRooms: []location{roomC1, roomC2, roomC3, roomC4},
		moves: []move{
			{
				to:       hall1,
				path:     []location{roomC2, roomC1, hall4, hall3, hall2, hall1},
				distance: 9,
			},
			{
				to:       hall2,
				path:     []location{roomC2, roomC1, hall4, hall3, hall2},
				distance: 8,
			},
			{
				to:       hall3,
				path:     []location{roomC2, roomC1, hall4, hall3},
				distance: 6,
			},
			{
				to:       hall4,
				path:     []location{roomC2, roomC1, hall4},
				distance: 4,
			},
			{
				to:       hall5,
				path:     []location{roomC2, roomC1, hall5},
				distance: 4,
			},
			{
				to:       hall6,
				path:     []location{roomC2, roomC1, hall5, hall6},
				distance: 6,
			},
			{
				to:       hall7,
				path:     []location{roomC2, roomC1, hall5, hall6, hall7},
				distance: 7,
			},
		},
	},
	roomC4: {
		occupant:  copperAmphipod,
		peerRooms: []location{roomC1, roomC2, roomC3, roomC4},
		moves: []move{
			{
				to:       hall1,
				path:     []location{roomC3, roomC2, roomC1, hall4, hall3, hall2, hall1},
				distance: 10,
			},
			{
				to:       hall2,
				path:     []location{roomC3, roomC2, roomC1, hall4, hall3, hall2},
				distance: 9,
			},
			{
				to:       hall3,
				path:     []location{roomC3, roomC2, roomC1, hall4, hall3},
				distance: 7,
			},
			{
				to:       hall4,
				path:     []location{roomC3, roomC2, roomC1, hall4},
				distance: 5,
			},
			{
				to:       hall5,
				path:     []location{roomC3, roomC2, roomC1, hall5},
				distance: 5,
			},
			{
				to:       hall6,
				path:     []location{roomC3, roomC2, roomC1, hall5, hall6},
				distance: 7,
			},
			{
				to:       hall7,
				path:     []location{roomC3, roomC2, roomC1, hall5, hall6, hall7},
				distance: 8,
			},
		},
	},
	roomD1: {
		occupant:  desertAmphipod,
		peerRooms: []location{roomD1, roomD2, roomD3, roomD4},
		moves: []move{
			{
				to:       hall1,
				path:     []location{hall5, hall4, hall3, hall2, hall1},
				distance: 9,
			},
			{
				to:       hall2,
				path:     []location{hall5, hall4, hall3, hall2},
				distance: 8,
			},
			{
				to:       hall3,
				path:     []location{hall5, hall4, hall3},
				distance: 6,
			},
			{
				to:       hall4,
				path:     []location{hall5, hall4},
				distance: 4,
			},
			{
				to:       hall5,
				path:     []location{hall5},
				distance: 2,
			},
			{
				to:       hall6,
				path:     []location{hall6},
				distance: 2,
			},
			{
				to:       hall7,
				path:     []location{hall6, hall7},
				distance: 3,
			},
		},
	},
	roomD2: {
		occupant:  desertAmphipod,
		peerRooms: []location{roomD1, roomD2, roomD3, roomD4},
		moves: []move{
			{
				to:       hall1,
				path:     []location{roomD1, hall5, hall4, hall3, hall2, hall1},
				distance: 10,
			},
			{
				to:       hall2,
				path:     []location{roomD1, hall5, hall4, hall3, hall2},
				distance: 9,
			},
			{
				to:       hall3,
				path:     []location{roomD1, hall5, hall4, hall3},
				distance: 7,
			},
			{
				to:       hall4,
				path:     []location{roomD1, hall5, hall4},
				distance: 5,
			},
			{
				to:       hall5,
				path:     []location{roomD1, hall5},
				distance: 3,
			},
			{
				to:       hall6,
				path:     []location{roomD1, hall6},
				distance: 3,
			},
			{
				to:       hall7,
				path:     []location{roomD1, hall6, hall7},
				distance: 4,
			},
		},
	},
	roomD3: {
		occupant:  desertAmphipod,
		peerRooms: []location{roomD1, roomD2, roomD3, roomD4},
		moves: []move{
			{
				to:       hall1,
				path:     []location{roomD2, roomD1, hall5, hall4, hall3, hall2, hall1},
				distance: 11,
			},
			{
				to:       hall2,
				path:     []location{roomD2, roomD1, hall5, hall4, hall3, hall2},
				distance: 10,
			},
			{
				to:       hall3,
				path:     []location{roomD2, roomD1, hall5, hall4, hall3},
				distance: 8,
			},
			{
				to:       hall4,
				path:     []location{roomD2, roomD1, hall5, hall4},
				distance: 6,
			},
			{
				to:       hall5,
				path:     []location{roomD2, roomD1, hall5},
				distance: 4,
			},
			{
				to:       hall6,
				path:     []location{roomD2, roomD1, hall6},
				distance: 4,
			},
			{
				to:       hall7,
				path:     []location{roomD2, roomD1, hall6, hall7},
				distance: 5,
			},
		},
	},
	roomD4: {
		occupant:  desertAmphipod,
		peerRooms: []location{roomD1, roomD2, roomD3, roomD4},
		moves: []move{
			{
				to:       hall1,
				path:     []location{roomD3, roomD2, roomD1, hall5, hall4, hall3, hall2, hall1},
				distance: 12,
			},
			{
				to:       hall2,
				path:     []location{roomD3, roomD2, roomD1, hall5, hall4, hall3, hall2},
				distance: 11,
			},
			{
				to:       hall3,
				path:     []location{roomD3, roomD2, roomD1, hall5, hall4, hall3},
				distance: 9,
			},
			{
				to:       hall4,
				path:     []location{roomD3, roomD2, roomD1, hall5, hall4},
				distance: 7,
			},
			{
				to:       hall5,
				path:     []location{roomD3, roomD2, roomD1, hall5},
				distance: 5,
			},
			{
				to:       hall6,
				path:     []location{roomD3, roomD2, roomD1, hall6},
				distance: 5,
			},
			{
				to:       hall7,
				path:     []location{roomD3, roomD2, roomD1, hall6, hall7},
				distance: 6,
			},
		},
	},
	hall1: {
		moves: []move{
			{
				to:       roomA1,
				path:     []location{hall2, roomA1},
				distance: 3,
			},
			{
				to:       roomA2,
				path:     []location{hall2, roomA1, roomA2},
				distance: 4,
			},
			{
				to:       roomA3,
				path:     []location{hall2, roomA1, roomA2, roomA3},
				distance: 5,
			},
			{
				to:       roomA4,
				path:     []location{hall2, roomA1, roomA2, roomA3, roomA4},
				distance: 6,
			},
			{
				to:       roomB1,
				path:     []location{hall2, hall3, roomB1},
				distance: 5,
			},
			{
				to:       roomB2,
				path:     []location{hall2, hall3, roomB1, roomB2},
				distance: 6,
			},
			{
				to:       roomB3,
				path:     []location{hall2, hall3, roomB1, roomB2, roomB3},
				distance: 7,
			},
			{
				to:       roomB4,
				path:     []location{hall2, hall3, roomB1, roomB2, roomB3, roomB4},
				distance: 8,
			},
			{
				to:       roomC1,
				path:     []location{hall2, hall3, hall4, roomC1},
				distance: 7,
			},
			{
				to:       roomC2,
				path:     []location{hall2, hall3, hall4, roomC1, roomC2},
				distance: 8,
			},
			{
				to:       roomC3,
				path:     []location{hall2, hall3, hall4, roomC1, roomC2, roomC3},
				distance: 9,
			},
			{
				to:       roomC4,
				path:     []location{hall2, hall3, hall4, roomC1, roomC2, roomC3, roomC4},
				distance: 10,
			},
			{
				to:       roomD1,
				path:     []location{hall2, hall3, hall4, hall5, roomD1},
				distance: 9,
			},
			{
				to:       roomD2,
				path:     []location{hall2, hall3, hall4, hall5, roomD1, roomD2},
				distance: 10,
			},
			{
				to:       roomD3,
				path:     []location{hall2, hall3, hall4, hall5, roomD1, roomD2, roomD3},
				distance: 11,
			},
			{
				to:       roomD4,
				path:     []location{hall2, hall3, hall4, hall5, roomD1, roomD2, roomD3, roomD4},
				distance: 12,
			},
		},
	},
	hall2: {
		moves: []move{
			{
				to:       roomA1,
				path:     []location{roomA1},
				distance: 2,
			},
			{
				to:       roomA2,
				path:     []location{roomA1, roomA2},
				distance: 3,
			},
			{
				to:       roomA3,
				path:     []location{roomA1, roomA2, roomA3},
				distance: 4,
			},
			{
				to:       roomA4,
				path:     []location{roomA1, roomA2, roomA3, roomA4},
				distance: 5,
			},
			{
				to:       roomB1,
				path:     []location{hall3, roomB1},
				distance: 4,
			},
			{
				to:       roomB2,
				path:     []location{hall3, roomB1, roomB2},
				distance: 5,
			},
			{
				to:       roomB3,
				path:     []location{hall3, roomB1, roomB2, roomB3},
				distance: 6,
			},
			{
				to:       roomB4,
				path:     []location{hall3, roomB1, roomB2, roomB3, roomB4},
				distance: 7,
			},
			{
				to:       roomC1,
				path:     []location{hall3, hall4, roomC1},
				distance: 6,
			},
			{
				to:       roomC2,
				path:     []location{hall3, hall4, roomC1, roomC2},
				distance: 7,
			},
			{
				to:       roomC3,
				path:     []location{hall3, hall4, roomC1, roomC2, roomC3},
				distance: 8,
			},
			{
				to:       roomC4,
				path:     []location{hall3, hall4, roomC1, roomC2, roomC3, roomC4},
				distance: 9,
			},
			{
				to:       roomD1,
				path:     []location{hall3, hall4, hall5, roomD1},
				distance: 8,
			},
			{
				to:       roomD2,
				path:     []location{hall3, hall4, hall5, roomD1, roomD2},
				distance: 9,
			},
			{
				to:       roomD3,
				path:     []location{hall3, hall4, hall5, roomD1, roomD2, roomD3},
				distance: 10,
			},
			{
				to:       roomD4,
				path:     []location{hall3, hall4, hall5, roomD1, roomD2, roomD3, roomD4},
				distance: 11,
			},
		},
	},
	hall3: {
		moves: []move{
			{
				to:       roomA1,
				path:     []location{roomA1},
				distance: 2,
			},
			{
				to:       roomA2,
				path:     []location{roomA1, roomA2},
				distance: 3,
			},
			{
				to:       roomA3,
				path:     []location{roomA1, roomA2, roomA3},
				distance: 4,
			},
			{
				to:       roomA4,
				path:     []location{roomA1, roomA2, roomA3, roomA4},
				distance: 5,
			},
			{
				to:       roomB1,
				path:     []location{roomB1},
				distance: 2,
			},
			{
				to:       roomB2,
				path:     []location{roomB1, roomB2},
				distance: 3,
			},
			{
				to:       roomB3,
				path:     []location{roomB1, roomB2, roomB3},
				distance: 4,
			},
			{
				to:       roomB4,
				path:     []location{roomB1, roomB2, roomB3, roomB4},
				distance: 5,
			},
			{
				to:       roomC1,
				path:     []location{hall4, roomC1},
				distance: 4,
			},
			{
				to:       roomC2,
				path:     []location{hall4, roomC1, roomC2},
				distance: 5,
			},
			{
				to:       roomC3,
				path:     []location{hall4, roomC1, roomC2, roomC3},
				distance: 6,
			},
			{
				to:       roomC4,
				path:     []location{hall4, roomC1, roomC2, roomC3, roomC4},
				distance: 7,
			},
			{
				to:       roomD1,
				path:     []location{hall4, hall5, roomD1},
				distance: 6,
			},
			{
				to:       roomD2,
				path:     []location{hall4, hall5, roomD1, roomD2},
				distance: 7,
			},
			{
				to:       roomD3,
				path:     []location{hall4, hall5, roomD1, roomD2, roomD3},
				distance: 8,
			},
			{
				to:       roomD4,
				path:     []location{hall4, hall5, roomD1, roomD2, roomD3, roomD4},
				distance: 9,
			},
		},
	},
	hall4: {
		moves: []move{
			{
				to:       roomA1,
				path:     []location{hall3, roomA1},
				distance: 4,
			},
			{
				to:       roomA2,
				path:     []location{hall3, roomA1, roomA2},
				distance: 5,
			},
			{
				to:       roomA3,
				path:     []location{hall3, roomA1, roomA2, roomA3},
				distance: 6,
			},
			{
				to:       roomA4,
				path:     []location{hall3, roomA1, roomA2, roomA4},
				distance: 7,
			},
			{
				to:       roomB1,
				path:     []location{roomB1},
				distance: 2,
			},
			{
				to:       roomB2,
				path:     []location{roomB1, roomB2},
				distance: 3,
			},
			{
				to:       roomB3,
				path:     []location{roomB1, roomB2, roomB3},
				distance: 4,
			},
			{
				to:       roomB4,
				path:     []location{roomB1, roomB2, roomB3, roomB4},
				distance: 5,
			},
			{
				to:       roomC1,
				path:     []location{roomC1},
				distance: 2,
			},
			{
				to:       roomC2,
				path:     []location{roomC1, roomC2},
				distance: 3,
			},
			{
				to:       roomC3,
				path:     []location{roomC1, roomC2, roomC3},
				distance: 4,
			},
			{
				to:       roomC4,
				path:     []location{roomC1, roomC2, roomC3, roomC4},
				distance: 5,
			},
			{
				to:       roomD1,
				path:     []location{hall5, roomD1},
				distance: 4,
			},
			{
				to:       roomD2,
				path:     []location{hall5, roomD1, roomD2},
				distance: 5,
			},
			{
				to:       roomD3,
				path:     []location{hall5, roomD1, roomD2, roomD3},
				distance: 6,
			},
			{
				to:       roomD4,
				path:     []location{hall5, roomD1, roomD2, roomD3, roomD4},
				distance: 7,
			},
		},
	},
	hall5: {
		moves: []move{
			{
				to:       roomA1,
				path:     []location{hall4, hall3, roomA1},
				distance: 6,
			},
			{
				to:       roomA2,
				path:     []location{hall4, hall3, roomA1, roomA2},
				distance: 7,
			},
			{
				to:       roomA3,
				path:     []location{hall4, hall3, roomA1, roomA2, roomA3},
				distance: 7,
			},
			{
				to:       roomA4,
				path:     []location{hall4, hall3, roomA1, roomA2, roomA3, roomA4},
				distance: 8,
			},
			{
				to:       roomB1,
				path:     []location{hall4, roomB1},
				distance: 4,
			},
			{
				to:       roomB2,
				path:     []location{hall4, roomB1, roomB2},
				distance: 5,
			},
			{
				to:       roomB3,
				path:     []location{hall4, roomB1, roomB2, roomB3},
				distance: 6,
			},
			{
				to:       roomB4,
				path:     []location{hall4, roomB1, roomB2, roomB3, roomB4},
				distance: 7,
			},
			{
				to:       roomC1,
				path:     []location{roomC1},
				distance: 2,
			},
			{
				to:       roomC2,
				path:     []location{roomC1, roomC2},
				distance: 3,
			},
			{
				to:       roomC3,
				path:     []location{roomC1, roomC2, roomC3},
				distance: 4,
			},
			{
				to:       roomC4,
				path:     []location{roomC1, roomC2, roomC3, roomC4},
				distance: 5,
			},
			{
				to:       roomD1,
				path:     []location{roomD1},
				distance: 2,
			},
			{
				to:       roomD2,
				path:     []location{roomD1, roomD2},
				distance: 3,
			},
			{
				to:       roomD3,
				path:     []location{roomD1, roomD2, roomD3},
				distance: 4,
			},
			{
				to:       roomD4,
				path:     []location{roomD1, roomD2, roomD3, roomD4},
				distance: 5,
			},
		},
	},
	hall6: {
		moves: []move{
			{
				to:       roomA1,
				path:     []location{hall5, hall4, hall3, roomA1},
				distance: 8,
			},
			{
				to:       roomA2,
				path:     []location{hall5, hall4, hall3, roomA1, roomA2},
				distance: 9,
			},
			{
				to:       roomA3,
				path:     []location{hall5, hall4, hall3, roomA1, roomA2, roomA3},
				distance: 10,
			},
			{
				to:       roomA4,
				path:     []location{hall5, hall4, hall3, roomA1, roomA2, roomA3, roomA4},
				distance: 11,
			},
			{
				to:       roomB1,
				path:     []location{hall5, hall4, roomB1},
				distance: 6,
			},
			{
				to:       roomB2,
				path:     []location{hall5, hall4, roomB1, roomB2},
				distance: 7,
			},
			{
				to:       roomB3,
				path:     []location{hall5, hall4, roomB1, roomB2, roomB3},
				distance: 8,
			},
			{
				to:       roomB4,
				path:     []location{hall5, hall4, roomB1, roomB2, roomB3, roomB4},
				distance: 9,
			},
			{
				to:       roomC1,
				path:     []location{hall5, roomC1},
				distance: 4,
			},
			{
				to:       roomC2,
				path:     []location{hall5, roomC1, roomC2},
				distance: 5,
			},
			{
				to:       roomC3,
				path:     []location{hall5, roomC1, roomC2, roomC3},
				distance: 6,
			},
			{
				to:       roomC4,
				path:     []location{hall5, roomC1, roomC2, roomC3, roomC4},
				distance: 7,
			},
			{
				to:       roomD1,
				path:     []location{roomD1},
				distance: 2,
			},
			{
				to:       roomD2,
				path:     []location{roomD1, roomD2},
				distance: 3,
			},
			{
				to:       roomD3,
				path:     []location{roomD1, roomD2, roomD3},
				distance: 4,
			},
			{
				to:       roomD4,
				path:     []location{roomD1, roomD2, roomD3, roomD4},
				distance: 5,
			},
		},
	},
	hall7: {
		moves: []move{
			{
				to:       roomA1,
				path:     []location{hall6, hall5, hall4, hall3, roomA1},
				distance: 9,
			},
			{
				to:       roomA2,
				path:     []location{hall6, hall5, hall4, hall3, roomA1, roomA2},
				distance: 10,
			},
			{
				to:       roomA3,
				path:     []location{hall6, hall5, hall4, hall3, roomA1, roomA2, roomA3},
				distance: 11,
			},
			{
				to:       roomA4,
				path:     []location{hall6, hall5, hall4, hall3, roomA1, roomA2, roomA3, roomA4},
				distance: 12,
			},
			{
				to:       roomB1,
				path:     []location{hall6, hall5, hall4, roomB1},
				distance: 7,
			},
			{
				to:       roomB2,
				path:     []location{hall6, hall5, hall4, roomB1, roomB2},
				distance: 8,
			},
			{
				to:       roomB3,
				path:     []location{hall6, hall5, hall4, roomB1, roomB2, roomB3},
				distance: 9,
			},
			{
				to:       roomB4,
				path:     []location{hall6, hall5, hall4, roomB1, roomB2, roomB3, roomB4},
				distance: 10,
			},
			{
				to:       roomC1,
				path:     []location{hall6, hall5, roomC1},
				distance: 5,
			},
			{
				to:       roomC2,
				path:     []location{hall6, hall5, roomC1, roomC2},
				distance: 6,
			},
			{
				to:       roomC3,
				path:     []location{hall6, hall5, roomC1, roomC2, roomC3},
				distance: 7,
			},
			{
				to:       roomC4,
				path:     []location{hall6, hall5, roomC1, roomC2, roomC3, roomC4},
				distance: 8,
			},
			{
				to:       roomD1,
				path:     []location{hall6, roomD1},
				distance: 3,
			},
			{
				to:       roomD2,
				path:     []location{hall6, roomD1, roomD2},
				distance: 4,
			},
			{
				to:       roomD3,
				path:     []location{hall6, roomD1, roomD2, roomD3},
				distance: 5,
			},
			{
				to:       roomD4,
				path:     []location{hall6, roomD1, roomD2, roomD3, roomD4},
				distance: 6,
			},
		},
	},
}
