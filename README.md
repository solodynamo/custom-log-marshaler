# 🕵️ Custom Log Marshaler

"Don't log any data that is unnecessary or should not be logged in the first place." - common sense

We can tag PII struct fields as "notloggable" and this generator will output custom marshal functions for those Golang structs which will prevent sending those fields to stdout!

<p align="center">
  <a href="https://goreportcard.com/report/github.com/solodynamo/custom-log-marshaler">
    <img src="https://goreportcard.com/badge/github.com/solodynamo/custom-log-marshaler" />
  </a>
  <a href="https://github.com/solodynamo/custom-log-marshaler/releases">
    <img src="https://img.shields.io/github/release/solodynamo/custom-log-marshaler.svg" />
  </a>
</p>

## How?

In most Go logging libraries, there is a way to override the default marshaling function used to "log" to stdout. This package generates that custom function with some superpowers like excluding fields (for PII, not required stuff).

If you send less data, ingestion bandwidth is used less so this also leads to lesser costs in the long run.

Example:
```go
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

Couldn't find something like this.

# Installation

MacOS
```bash
brew tap solodynamo/homebrew-tap
brew install custom-log-marshaler
```

Others: 
```bash
go install github.com/solodynamo/custom-log-marshaler
```

# Usage

Zerolog

```bash
custom-log-marshaler -f "path to go file" -lib zerolog
```

Uber Zap(by default)

```bash
custom-log-marshaler -f "path to go file"

```

For bulk operation on multiple go files: 

```bash
curl -sSL https://gist.githubusercontent.com/solodynamo/b23de7cc6576179292871efc9b37e1f1/raw/apply-clm-go.sh | bash -s -- "path1/to/ignore" "path2/to/ignore"

```
Note: Above bulk script handles the installation too but only for MacOS.

# How to exclude PII? 
```go
type User struct {
	Email   string `notloggable`
}
```
When custom struct tag `notloggable` is used for a field, that field is excluded from final logging irrespective of the lib. 

# Contributions
Please do!

# Just a joke
<img width="789" alt="chatgptjoke" src="https://user-images.githubusercontent.com/17698714/220406972-0a42c233-fe71-4f58-b337-f10deb4f171c.png">