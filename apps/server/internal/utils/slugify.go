package utils

import (
	"fmt"
	"regexp"
	"strings"
)

// Slugify converts a string to a URL-friendly slug
func Slugify(text string, opts ...Options) string {
	// Default options
	options := Options{
		Replacement: "-",
		Lower:       true,
		Strict:      false,
		Trim:        true,
	}

	// If options provided, use them
	if len(opts) > 0 {
		options = opts[0]
	}

	// Rest of your slugify logic using options...
	slug := normalizeUnicode(text)

	if options.Lower {
		slug = strings.ToLower(slug)
	}

	if options.Remove != nil {
		slug = options.Remove.ReplaceAllString(slug, "")
	}

	if options.Strict {
		slug = regexp.MustCompile(`[^a-z0-9]+`).ReplaceAllString(slug, options.Replacement)
	} else {
		slug = regexp.MustCompile(`[\s_]+`).ReplaceAllString(slug, options.Replacement)
	}

	if options.Trim {
		slug = strings.Trim(slug, options.Replacement)
	}

	slug = regexp.MustCompile(fmt.Sprintf(`%s+`, regexp.QuoteMeta(options.Replacement))).ReplaceAllString(slug, options.Replacement)

	return slug
}

// type Option func(*Options)

type Options struct {
	Replacement string
	Remove      *regexp.Regexp
	Lower       bool
	Strict      bool
	Locale      string
	Trim        bool
	Separator   string
}

// normalizeUnicode normalizes unicode characters
func normalizeUnicode(s string) string {
	return strings.Map(func(r rune) rune {
		if r >= 0x0080 && r <= 0x00FF {
			switch r {
			case 'à', 'á', 'â', 'ã', 'ä', 'å':
				return 'a'
			case 'è', 'é', 'ê', 'ë':
				return 'e'
			case 'ì', 'í', 'î', 'ï':
				return 'i'
			case 'ò', 'ó', 'ô', 'õ', 'ö':
				return 'o'
			case 'ù', 'ú', 'û', 'ü':
				return 'u'
			case 'ý', 'ÿ':
				return 'y'
			case 'ç':
				return 'c'
			case 'ñ':
				return 'n'
			}
		}
		return r
	}, s)
}
