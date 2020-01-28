# Go Run Version

The main package contains it's state in the git tree statically. Meaning, I described statically in the source where in the tree it would be. This might be helpful to `go run` or `go get` to see which version actually was run or pulled. So far this repo hasn't been set up with go.mod intentionally, this can be another branch in the future if needed.

Current overview - remember, the last commit on branch master will always be missing ;)

```git
* ....... some more commits to master which are usually not so relevant
* ec291b5 (HEAD -> master, origin/master) init readme
| * ffdc7de (origin/third, third) ending third
| * 55c13f0 (tag: v2.0.0) third branch with v2
|/  
| * d86a9da (tag: v0.2.1, origin/second, second) full second branch
| * 9cc7822 (tag: v0.2.0) open second branch
|/  
* 10de6d0 (tag: v0.1.0) the first tag
* bcf6244 the initial commit
```

Depending on if you're in a repo with go.mod or not, play around with these and see how it behaves:

- `go run github.com/dionysius/go-run-version`
- `go run github.com/dionysius/go-run-version/v2`
- `go run github.com/dionysius/go-run-version@master`
- `go run github.com/dionysius/go-run-version@v0.1.0`
- `go run github.com/dionysius/go-run-version@second`
- ...
- `go get github.com/dionysius/go-run-version`
- `go get github.com/dionysius/go-run-version/v2`
  - and after that, whats the difference between
  - `go run github.com/dionysius/go-run-version`
  - `go run github.com/dionysius/go-run-version/v2`
- `go get github.com/dionysius/go-run-version@master`
- `go get github.com/dionysius/go-run-version@v0.1.0`
- `go get github.com/dionysius/go-run-version@second`
