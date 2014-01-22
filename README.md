Cangjie Parser
==============

This is a command line parser tool that helps studying Cangjie mapping characteristics. The tool will build an sqlite database file of all the Cangjie mappings. Researchers and developers may take advantage of a relational database to analyse and output data needed.

***

Dependencies
------------

- This tool is developed, and has only been tested working, on Linux

- Requires [Go](http://golang.org) version 1.1

- Requires [mattn's go-sqlite library](http://github.com/mattn/go-sqlite3)<br />
  (with internet connection, the build script will install it for you)


Build the Tool and Database
---------------------------

This software comes with a build script (build.sh). To build the commandline tool and the database, just run the build script in the same folder as this file is placed:

    ./build.sh

And a database file named "cangjie.db" will be generated in the "data" folder


***

About the Cangjie Data Used
---------------------------

The Cangjie mapping data files (distributed with this tool) are from [Friends of Cangjie, Malaysia (倉頡之友 馬來西亞)](http://chinesecj.com). Thanks for their efforts. If you appreciate their efforts to organize Cangjie input method, please consider a donation to them.

For continent to process, the data files distributed are processed:

- Cangjie 3 files are converted from BIG5/GB2312 to UTF-8<br />
  (while the original files are distributed also for verification)

- Cangjie 5 files are converted from *dosnoeol* to UNIX *noeol* format


***

Legalities
----------

**A) Cangjie Mapping Data Files**

The Cangjie mapping data files (files under data/cj3 and data/cj5) are created and own by *Friends of Cangjie, Malaysia*.

You are free to use and share the Cangjie mapping data under the following terms:

- **Attribution:** You must give appropriate credit to *Friends of Cangjie, Malaysia*. You may do so in any reasonable manner, but not in any way that suggests the licensor endorses you or your use.

- **Non-Commercial:** You may not use the material for commercial purposes. 

For commercial use, please contact *Friends of Cangjie, Malaysia*.

**B) All the other things**

Other than the files mentioned in (A), all other scripts, source code of cjparser are offered under the terms of the [GNU Lesser General Public License, either version 3 or any later version](http://www.gnu.org/licenses/lgpl.html).

You don't need to sign a copyright assignment or any other kind of silly and tedious legal document, so just send patches and/or pull requests!
