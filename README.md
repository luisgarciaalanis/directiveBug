# directiveBug

Directive is being called after the Resolver beating the purpose of a Role directive:

To repro run:
```
go build
./directiveBug
```
open in browser:
```
http://localhost:8080/ 
 ```

payload to test: 
```
{
  test{
    text
  }
}
```
Bug shows in console:
```
2020/10/07 01:03:25 connect to http://localhost:8080/ for GraphQL playground
Test resolver called!
HasRole called!
```
Expected:
```
2020/10/07 01:03:25 connect to http://localhost:8080/ for GraphQL playground
HasRole called!
Test resolver called!
```
