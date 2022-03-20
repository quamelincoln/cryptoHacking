#pragma once
/*-------------------------------------------------------------------------------------------
 * qblocks - fast, easily-accessible, fully-decentralized data from blockchains
 * copyright (c) 2016, 2021 TrueBlocks, LLC (http://trueblocks.io)
 *
 * This program is free software: you may redistribute it and/or modify it under the terms
 * of the GNU General Public License as published by the Free Software Foundation, either
 * version 3 of the License, or (at your option) any later version. This program is
 * distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even
 * the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
 * General Public License for more details. You should have received a copy of the GNU General
 * Public License along with this program. If not, see http://www.gnu.org/licenses/.
 *-------------------------------------------------------------------------------------------*/
#include "node.h"

namespace qblocks {

extern size_t nTimestamps(void);
extern blknum_t getTimestampBlockAt(blknum_t blk);
extern timestamp_t getTimestampAt(blknum_t bn);
extern bool freshenTimestamps(blknum_t minBlock);
extern bool correctTimestamp(blknum_t blk, timestamp_t ts);
extern bool loadTimestamps(uint32_t** theArray, size_t& cnt);

}  // namespace qblocks
