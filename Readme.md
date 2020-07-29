### Go-envdir - run a program with defined env-vars in files into directory

How to use:

`go-envdir /path/to/env-dir program`

Example: 

dir `env` contains 2 files: `A` has content `1` and `B` has content `2`.

So your program will be runned with an environment `A=1 B=2`
