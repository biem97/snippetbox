# Customizing how tests run
- Run tests in `web` package
```bash
go test ./cmd/web
```
- Run tests in current project
```bash
go test ./...
```
- Run tests with `-run` flag and pass in a regex
```bash
go test -v run="^TestPing$" ./cmd/web
go test -v run="^TestHumandate/^UTC" ./cmd/web
```
- Run test in full and avoid the cache with `-count=1` flag
```bash
go test -count=1 ./cmd/web
```
- Clear cache
```bash
go clean -testcache
```
- Stop tests ***in the package that had the failure*** with `failfast` flag
```bash
go test -failfast ./cmd/web
```
- Parallel testing
```go
func TestPing(t *testing.T) {
  t.Parallel()

  ...
}
```
  - Tests marked using t.Parallel() will be run in parallel with — and only with — other
    parallel tests.
  - By default, the maximum number of tests that will be run simultaneously is the current
    value of GOMAXPROCS. You can override this by setting a specific value via the -parallel
    flag. For example:
    ```golang
    $ go test -parallel 4 ./...
    ```
  - Not all tests are suitable to be run in parallel. For example, if you have an integration test
    which requires a database table to be in a specific known state, then you wouldn’t want to
    run it in parallel with other tests that manipulate the same database table.
- Race dector
```bash
go test -race ./cmd/web/
```

# Integration test
- Connect to MySQL as the root user and execute the following SQL statements to create a new `test_snippetbox` database and `test_web` user:
```sql
CREATE DATABASE test_snippetbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

```sql
CREATE USER 'test_web'@'localhost';
GRANT CREATE, DROP, ALTER, INDEX, SELECT, INSERT, UPDATE, DELETE ON test_snippetbox.* TO 'test_web'@'localhost';
ALTER USER 'test_web'@'localhost' IDENTIFIED BY 'pass';
```

# Go
- The Go tool ignores any directories called testdata , so these scripts will be
ignored when compiling your application (it also ignores any directories or files which
have names that begin with an _ or . character too).
