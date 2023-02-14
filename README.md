# Custom Log Marshaler
Generates PII containing logics for popular logging libs like zap, zerolog.

# Why? 
Couldn't find any packages which help me generate such logics in bulk and flexible enough to use with different logging libs. 

Closest one is https://github.com/muroon/zmlog and its helpful but few things that it doesn't have:

1. Generate custom functions in the same file as struct, better readability and maintainance. 
2. Extensible to other logging libs.
3. Specify what to log using struct tags, this can help contain PII info. Also save on log ingestion costs.


# Usage

Zerolog

```
go install github.com/solodynamo/pii-marshaler
pii-marshaler -f "path to go file" -lib zerolog

```

Uber Zap(by default)

```
go install github.com/solodynamo/pii-marshaler
pii-marshaler -f "path to go file" 

```

# PII

When custom struct tag `notloggable` is used for a field, that field is excluded from final logging irrespective of the lib. 

```
type User struct {
	Name    string `json:"name"`
	Email   string `notloggable`
	Address string `json:"address", notloggable`
}
```

In above case, only name will be logged!

See working examples in fixtuers.