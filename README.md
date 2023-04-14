# Custom Log Marshaler
"Don't log any data that is unnecessary or should not be logged in the first place." 

# How?
In most Go logging libs there is a way to override default marshaling function used to "log" to stdout, this package generates that custom function with some super powers like excluding fields(for PII, not required stuff).

If you send less data, ingestion bandwidth is used less so this also leads to lesser costs in longer run.

Example: 
```
type User struct {
	Name    string `json:"name"`
	Email   string `notloggable`
	Address string `json:"address", notloggable`
}

type UserDetailsResponse struct {
	User
	RequestID string   `json:"rid"`
	FromCache bool     `json:"fromCache"`
	Metadata  []string `json:"md"`
}
// MarshalLogObject ...
func (l User) MarshalLogObject(enc zapcore.ObjectEncoder) error {
		enc.AddString("name", l.Name) // not logging things which shouldn't be logged.
		return nil	
}

// MarshalLogObject ...
func (l UserDetailsResponse) MarshalLogObject(enc zapcore.ObjectEncoder) error {
		enc.AddObject("user", l.User)
		enc.AddString("request_id", l.RequestID)
		enc.AddBool("from_cache", l.FromCache)
		enc.AddArray("metadata", l.Metadata)
		return nil	
}

```

See above example in action on [Go Playground](https://go.dev/play/p/cv_u168fm0e?v=goprev). 

# Why? 
Couldn't find any packages which help me generate such logics in bulk and flexible enough to use with different logging libs. 

Closest one is https://github.com/muroon/zmlog and its helpful but few things that it doesn't have:

1. Generate custom functions in the same file as struct, better readability and maintainance. 
2. Extensible to other logging libs.
3. Specify what to log using struct tags, this can help contain PII info. Also save on log ingestion costs.
4. Handle pointer dereferencing.


# Usage

Zerolog

```
go install github.com/solodynamo/custom-log-marshaler
custom-log-marshaler -f "path to go file" -lib zerolog

```

Uber Zap(by default)

```
go install github.com/solodynamo/custom-log-marshaler
custom-log-marshaler -f "path to go file"

```

# PII

When custom struct tag `notloggable` is used for a field, that field is excluded from final logging irrespective of the lib. 

See working examples in fixtures for different libs. Just run tests at root and see generated output.


# Joke

<img width="789" alt="chatgptjoke" src="https://user-images.githubusercontent.com/17698714/220406972-0a42c233-fe71-4f58-b337-f10deb4f171c.png">