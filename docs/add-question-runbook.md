1. Modify Source code
   1. replace `main` with `Solve(src io.Reader) (answer string, err error)`
   2. create defer function to recover from panics
   3. add output to answer string (take care of `+=`)
   4. add init method to make sure os is imported
2. add route for question
3. create testcases
4. run `make test` and `make subimssion`