<p align="right">
  ⭐ &nbsp;&nbsp;<strong>the project to show your appreciation.</strong> :arrow_upper_right:
</p>

<p align="right">
  <a href="https://pkg.go.dev/github.com/rocketlaunchr/anti-disposable-email"><img src="http://godoc.org/github.com/rocketlaunchr/anti-disposable-email?status.svg" /></a>
  <a href="https://goreportcard.com/report/github.com/rocketlaunchr/anti-disposable-email"><img src="https://goreportcard.com/badge/github.com/rocketlaunchr/anti-disposable-email" /></a>
  <a href="https://gocover.io/github.com/rocketlaunchr/anti-disposable-email"><img src="http://gocover.io/_badge/github.com/rocketlaunchr/anti-disposable-email" /></a>
</p>

<p align="center">
	<img src="https://github.com/rocketlaunchr/anti-disposable-email/raw/master/logo.png" alt="anti-disposable-email" />
</p>

# Anti Disposable Email Address Checker for Go

If you want to detect if new accounts are using disposable email services, then this is your package.



## Installation


```
go get -u github.com/rocketlaunchr/anti-disposable-email
```

```go
import "github.com/rocketlaunchr/anti-disposable-email"
```


## Usage

```go
import "github.com/rocketlaunchr/anti-disposable-email"

ParsedEmail, _ := disposable.ParseEmail("rocketlaunchr.cloud@gmail.com")
````

### Output

```groovy
(disposable.ParsedEmail) {
 Email: (string) (len=21) "rocketlaunchr.cloud@gmail.com",
 Preferred: (string) (len=9) "rocketlaunchr.cloud",
 Normalized: (string) (len=9) "rocketlaunchrcloud",
 Extra: (string) "",
 Disposable: (bool) false,
 Domain: (string) (len=11) "gmail.com",
 LocalPart: (string) (len=9) "rocketlaunchr.cloud"
}

```

If `Disposable` is **true**, then the email address is from a disposable email service.

### Normalized

If you want to block duplicate email addresses from your database, then store as a unique-key the `Normalized` data. See [docs](https://pkg.go.dev/github.com/rocketlaunchr/anti-disposable-email#ParsedEmail).

### Update

This package can auto-update the disposable domain list. It uses the regularly updated list from here: https://github.com/disposable-email-domains/disposable-email-domains.

```go
import "github.com/rocketlaunchr/anti-disposable-email"
import "github.com/rocketlaunchr/anti-disposable-email/update"

update.Update(ctx, &disposable.DisposableList)
```


## Other useful packages

- [awesome-svelte](https://github.com/rocketlaunchr/awesome-svelte) - Resources for killing react
- [dataframe-go](https://github.com/rocketlaunchr/dataframe-go) - Statistics and data manipulation
- [dbq](https://github.com/rocketlaunchr/dbq) - Zero boilerplate database operations for Go
- [electron-alert](https://github.com/rocketlaunchr/electron-alert) - SweetAlert2 for Electron Applications
- [google-search](https://github.com/rocketlaunchr/google-search) - Scrape google search results
- [igo](https://github.com/rocketlaunchr/igo) - A Go transpiler with cool new syntax such as fordefer (defer for for-loops)
- [mysql-go](https://github.com/rocketlaunchr/mysql-go) - Properly cancel slow MySQL queries
- [react](https://github.com/rocketlaunchr/react) - Build front end applications using Go
- [remember-go](https://github.com/rocketlaunchr/remember-go) - Cache slow database queries
- [showerglass](https://github.com/rocketlaunchr/showerglass) - A soothing face filter for social applications.
- [testing-go](https://github.com/rocketlaunchr/testing-go) - Testing framework for unit testing

#

### Legal Information

The license is a modified MIT license. Refer to `LICENSE` file for more details.

**© 2022 PJ Engineering and Business Solutions Pty. Ltd.**
