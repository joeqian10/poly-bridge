//go:build mainnet
// +build mainnet

/*
 * Copyright (C) 2020 The poly network Authors
 * This file is part of The poly network library.
 *
 * The  poly network  is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The  poly network  is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 * You should have received a copy of the GNU Lesser General Public License
 * along with The poly network .  If not, see <http://www.gnu.org/licenses/>.
 */

package basedef

const (
	POLY_CROSSCHAIN_ID = uint64(0)

	ETHEREUM_CROSSCHAIN_ID = uint64(2)

	BSC_CROSSCHAIN_ID = uint64(6)

	NEO3_CROSSCHAIN_ID = uint64(14)

	ONTEVM_CROSSCHAIN_ID = uint64(47)

	ENV = "mainnet"
)

const (
	BSC_NORMAL_GASPRICE   = 5000000000
	ASTAR_NORMAL_GASPRICE = 60000000000
)

var ETH_CHAINS = []uint64{
	ETHEREUM_CROSSCHAIN_ID, BSC_CROSSCHAIN_ID, ONTEVM_CROSSCHAIN_ID,
}
