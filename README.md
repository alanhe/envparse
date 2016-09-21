# envparse

Inspried by the idea of **Store config in the environment** from [The Twelve-Factor App](https://12factor.net/config). The envparse lib aims to make reading environment varaibles easier.

Usage:
``` go
func main() {
    parser := envparse.New()

    parser.Add(&envparse.Param{name: "DB_URL", required: true})
    parser.Parse() // panic if DB_URL is not set

    dbURL := parser.getString("DB_URL") // retive a trimmed value

    // ...
}
```
