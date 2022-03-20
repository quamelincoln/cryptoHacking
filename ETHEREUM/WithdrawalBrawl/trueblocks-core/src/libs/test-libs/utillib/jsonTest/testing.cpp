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
#include "etherlib.h"
#include "options.h"

//-------------------------------------------------------------------------
ostream& operator<<(ostream& os, const CTraceArray& array) {
    cout << "[\n";
    for (size_t i = 0; i < array.size(); i++) {
        array[i].Format(os, "", nullptr);
        if (i < array.size() - 1)
            cout << ",";
        cout << "\n";
    }
    cout << "]\n";
    return os;
}

//-------------------------------------------------------------------------
int main(int argc, const char* argv[]) {
    etherlib_init(quickQuitHandler);

    // Parse command line, allowing for command files
    COptions options;
    if (!options.prepareArguments(argc, argv))
        return 0;

    for (auto command : options.commandLines) {
        if (!options.parseArguments(command))
            return 0;

        if (options.which == 0)
            options.doTestZero();
        else if (options.which == 1)
            options.doTestOne();
        else
            cerr << "Invalid test: " << options.which << endl;
    }

    return 0;
}

//-------------------------------------------------------------------------
void COptions::doTestZero(void) {
    string_q fn = "tests/" + fileName + ".json";
    cout << fn << "\n";

    string contents;
    asciiFileToString(fn, contents);
    if (fileName == "blocks") {
        CBlock block;
        while (block.parseJson3(contents))  // returns 'true' if anyting was parsed
            cout << block << "\n";

    } else if (fileName == "traces") {
        CTrace trace;
        while (trace.parseJson3(contents))
            cout << trace << "\n";

    } else if (fileName == "spaces") {
        CTrace trace;
        while (trace.parseJson3(contents))
            cout << trace << "\n";

    } else if (fileName == "escaped") {
        CTrace trace;
        while (trace.parseJson3(contents))
            cout << trace << "\n";

    } else if (fileName == "big") {
        CRPCResult generic;
        while (generic.parseJson3(contents)) {
            CTrace trace;
            while (trace.parseJson3(generic.result))
                cout << trace << "\n";
        }

    } else {
        GETRUNTIME_CLASS(CTrace)->sortFieldList();
        CTraceArray array;
        CTrace trace;
        while (trace.parseJson3(contents)) {
            array.push_back(trace);
            trace = CTrace();  // reset
        }
        cout << array << "\n";
    }
}

//-------------------------------------------------------------------------
void COptions::doTestOne(void) {
}
