package disposable

import (
	"errors"
	"strings"
	"unicode"
)

// ErrInvalidEmail is returned if the email address is invalid.
var ErrInvalidEmail = errors.New("invalid email")

// ParsedEmail returns a parsed email address.
//
// An email address is made up of 3 components: <local-part>@<domain>.
// The local-part is case-sensitive according to the specs, but most
// (if not all) reputable email services will treat it as case-insensitive.
type ParsedEmail struct {

	// Email represents the input email (after white-space has been trimmed).
	Email string

	// Preferred represents the local-part in the way the user seems to prefer it.
	// For example if the local-part is case-insensitive, the user may prefer their
	// email address all upper-case even if it does not matter.
	Preferred string

	// Normalized represents the local-part normalized such that it can be
	// compared for uniqueness.
	//
	// For gmail, since john.smith@gmail.com, johnsmith@gmail.com, and JohnSmith@gmail.com
	// are all equivalent, the normalized local-part is 'johnsmith'.
	Normalized string

	// Extra represents extra information that is domain specific.
	//
	// Example: gmail ignores all characters after the first '+' in the local-part.
	//
	// adam+junk@gmail.com => adam@gmail.com (Extra: junk)
	Extra string

	// Disposable is true if the email address is detected to be from
	// a disposable email service.
	//
	// See: https://github.com/martenson/disposable-email-domains
	Disposable bool

	// Domain represents the component after the '@' character.
	// It is lower-cased since it's not case-sensitive.
	Domain string
}

// ParseEmail parses a given email address. Set caseSensitive to true if you want the local-part
// to be considered case-sensitive. The default is false. Basic email validation is performed but
// it is not comprehensively checked.
//
// See https://github.com/badoux/checkmail for a more robust validation solution.
//
// See also https://davidcel.is/posts/stop-validating-email-addresses-with-regex.
//
func ParseEmail(email string, caseSensitive ...bool) (ParsedEmail, error) {

	// Perform basic validation
	email = strings.TrimSpace(email)

	if email == "" {
		return ParsedEmail{}, ErrInvalidEmail
	}

	if strings.Contains(email, " ") {
		return ParsedEmail{Email: email}, ErrInvalidEmail
	}

	var cs bool
	if len(caseSensitive) > 0 {
		cs = caseSensitive[0]
	}

	splits := strings.Split(email, "@")
	if len(splits) != 2 {
		return ParsedEmail{Email: email}, ErrInvalidEmail
	}

	domain := toLower(splits[1])

	p := ParsedEmail{
		Email:  email,
		Domain: domain,
	}

	// Normalize local part
	p.Normalized, p.Preferred, p.Extra = normalize(splits[0], domain, cs)

	// Check if domain is disposable
	_, p.Disposable = DisposableList[domain]

	return p, nil

}

func normalize(localPart, domain string, caseSensitive bool) (ret string, pref string, sufx string) {
	pref = localPart

	switch domain {
	case "gmail.com":
		// remove suffix from localPart
		splits := strings.SplitN(localPart, "+", 2)
		if len(splits) == 2 {
			localPart, sufx = splits[0], splits[1]
			pref = localPart
		}

		// remove the periods
		localPart = strings.ReplaceAll(localPart, ".", "")
	}

	// lower-case the local part
	if caseSensitive {
		ret = localPart
		return
	}

	ret = toLower(localPart)
	return
}

func toLower(s string) (ret string) {
	for _, r := range s {
		ret += string(unicode.ToLower(r))
	}
	return
}
