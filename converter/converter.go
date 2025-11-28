package converter

import (
	"fmt"
	"regexp"
	"strings"
)

// Config holds conversion configuration options
type Config struct {
	AddStyling bool // Include CSS styling in output
	Concurrent bool // Use concurrent processing for large files
}

// Tietoevry brand colors
var headingColors = map[int]string{
	1: "#021e57", // Hero Blue
	2: "#021e57", // Dark Navy Blue
	3: "#021e57", // Dark Navy Blue
	4: "#021e57", // Dark Navy Blue
	5: "#021e57", // Dark Navy Blue
	6: "#021e57", // Dark Navy Blue
}

// GetCodeStylingCSS generates MediaWiki CSS for accessible syntax highlighting
func GetCodeStylingCSS() string {
	return `<!-- Accessible Syntax Highlighting (WCAG AA compliant) -->
<style>
/* Code block container */
.mw-highlight {
    background-color: #FAFAFA !important;
    border: 1px solid #CCCCCC !important;
    border-left: 3px solid #0000FF !important;
    padding: 1em !important;
    border-radius: 4px;
    font-family: 'Consolas', 'Monaco', 'Courier New', monospace !important;
    font-size: 0.95em !important;
    line-height: 1.5 !important;
}

/* Syntax highlighting - Industry standard accessible colors */
.mw-highlight .c,   /* Comments */
.mw-highlight .cm,  /* Multi-line comments */
.mw-highlight .c1 { /* Single-line comments */
    color: #A0A1A7 !important;
    font-style: italic !important;
}

.mw-highlight .k,   /* Keywords */
.mw-highlight .kd,  /* Keyword declarations */
.mw-highlight .kn,  /* Keyword namespace */
.mw-highlight .kp,  /* Keyword pseudo */
.mw-highlight .kr,  /* Keyword reserved */
.mw-highlight .kt { /* Keyword type */
    color: #0000FF !important;
    font-weight: 600 !important;
}

.mw-highlight .s,   /* Strings */
.mw-highlight .s1,  /* Single-quoted strings */
.mw-highlight .s2,  /* Double-quoted strings */
.mw-highlight .sb,  /* String backtick */
.mw-highlight .sc { /* String char */
    color: #50A14F !important;
}

.mw-highlight .m,   /* Numbers */
.mw-highlight .mf,  /* Float */
.mw-highlight .mi,  /* Integer */
.mw-highlight .mo,  /* Octal */
.mw-highlight .mh { /* Hex */
    color: #A626A4 !important;
}

.mw-highlight .n,   /* Names/variables */
.mw-highlight .nv { /* Variable */
    color: #383A42 !important;
}

.mw-highlight .nf,  /* Function name */
.mw-highlight .fm { /* Function magic */
    color: #4078F2 !important;
    font-weight: 500 !important;
}

.mw-highlight .o,   /* Operators */
.mw-highlight .ow { /* Operator word */
    color: #383A42 !important;
}

/* XML/JSON specific */
.mw-highlight .nt { /* XML/HTML tags */
    color: #E45649 !important;
    font-weight: 600 !important;
}

.mw-highlight .na { /* XML/HTML attributes */
    color: #986801 !important;
}

.mw-highlight .p { /* Punctuation */
    color: #383A42 !important;
}

/* Line numbers */
.mw-highlight .linenos {
    background-color: #F0F0F0 !important;
    color: #9D9D9F !important;
    padding-right: 1em !important;
    padding-left: 0.5em !important;
    border-right: 1px solid #CCCCCC !important;
    user-select: none !important;
}

/* Boolean/Null/Constants */
.mw-highlight .kc,  /* Constant */
.mw-highlight .bp { /* Builtin pseudo */
    color: #A626A4 !important;
}

/* Inline code styling */
code {
    background-color: #f5ff56 !important;
    color: #021e57 !important;
    padding: 2px 6px !important;
    border-radius: 3px !important;
    font-family: 'Consolas', 'Monaco', 'Courier New', monospace !important;
}

/* Tietoevry Heading Colors - Purple gradient */
.mw-parser-output h1,
h1.firstHeading {
    color: #021e57 !important;  /* Hero Blue */
    border-bottom: 2px solid #021e57 !important;
    font-weight: 600 !important;
}

.mw-parser-output h2,
h2 {
    color: #021e57 !important;  /* Dark Navy Blue */
    border-bottom: 1px solid #021e57 !important;
    font-weight: 600 !important;
}

.mw-parser-output h3,
h3 {
    color: #021e57 !important;  /* Dark Navy Blue */
    font-weight: 600 !important;
}

.mw-parser-output h4,
h4 {
    color: #021e57 !important;  /* Dark Navy Blue */
    font-weight: 600 !important;
}

.mw-parser-output h5,
h5 {
    color: #021e57 !important;  /* Dark Navy Blue */
    font-weight: 600 !important;
}

.mw-parser-output h6,
h6 {
    color: #021e57 !important;  /* Dark Navy Blue */
    font-weight: 600 !important;
}
</style>

`
}

// ConvertHeaders converts Markdown headers to MediaWiki format with Tietoevry colors
func ConvertHeaders(text string) string {
	lines := strings.Split(text, "\n")
	result := make([]string, 0, len(lines))

	headerRegex := regexp.MustCompile(`^(#{1,6})\s+(.+)$`)

	for _, line := range lines {
		matches := headerRegex.FindStringSubmatch(line)
		if matches != nil {
			level := len(matches[1])
			content := matches[2]
			color := headingColors[level]
			equals := strings.Repeat("=", level)
			result = append(result, fmt.Sprintf(`%s<span style="color:%s;">%s</span>%s`, equals, color, content, equals))
		} else {
			result = append(result, line)
		}
	}

	return strings.Join(result, "\n")
}

// ConvertBoldItalic converts bold and italic formatting
func ConvertBoldItalic(text string) string {
	// Protect code blocks from processing by temporarily replacing them
	codeBlockRegex := regexp.MustCompile(`(?s)<syntaxhighlight[^>]*>.*?</syntaxhighlight>`)
	codeBlocks := codeBlockRegex.FindAllString(text, -1)

	// Replace code blocks with placeholders (use base64-like format to avoid special chars)
	for i, block := range codeBlocks {
		placeholder := fmt.Sprintf("XYZCODEBLOCKREPLACEMENTXYZ%dXYZ", i)
		text = strings.Replace(text, block, placeholder, 1)
	}

	// Obsidian highlights: ==text== -> <mark style="background-color:#f5ff56">text</mark>
	highlightRegex := regexp.MustCompile(`==([^=\n]+)==`)
	text = highlightRegex.ReplaceAllString(text, `<mark style="background-color:#f5ff56">$1</mark>`)

	// Bold: **text** or __text__ -> '''text'''
	boldRegex1 := regexp.MustCompile(`\*\*(.+?)\*\*`)
	text = boldRegex1.ReplaceAllString(text, `'''$1'''`)
	boldRegex2 := regexp.MustCompile(`__(.+?)__`)
	text = boldRegex2.ReplaceAllString(text, `'''$1'''`)

	// Italic: *text* or _text_ -> ''text''
	// More careful regex to avoid matching bold
	italicRegex1 := regexp.MustCompile(`(?:^|[^\*])\*([^\*\n]+?)\*(?:[^\*]|$)`)
	text = italicRegex1.ReplaceAllString(text, `''$1''`)
	italicRegex2 := regexp.MustCompile(`(?:^|[^_])_([^_\n]+?)_(?:[^_]|$)`)
	text = italicRegex2.ReplaceAllString(text, `''$1''`)

	// Restore code blocks
	for i, block := range codeBlocks {
		placeholder := fmt.Sprintf("XYZCODEBLOCKREPLACEMENTXYZ%dXYZ", i)
		text = strings.Replace(text, placeholder, block, 1)
	}

	return text
}

// ConvertLinks converts Markdown links to MediaWiki format
func ConvertLinks(text string) string {
	// External links: [text](url) -> [url text]
	linkRegex := regexp.MustCompile(`\[([^\]]+)\]\((https?://[^\)]+)\)`)
	text = linkRegex.ReplaceAllString(text, `[$2 $1]`)

	return text
}

// ConvertCallouts converts markdown callouts to MediaWiki styled boxes
func ConvertCallouts(text string) string {
	// Info boxes - using Tietoevry cool palette
	infoRegex := regexp.MustCompile(`(?m)>\s*\[!info\]\s*(.*)$`)
	text = infoRegex.ReplaceAllString(text, `{| class="wikitable" style="border-left:4px solid #021e57; background-color:#f7f7fa;"
| <div style="padding:0.5em;">
<strong style="color:#021e57;">ℹ️ Info:</strong> $1
</div>
|}`)

	// Warning boxes - using Tietoevry warm palette
	warningRegex := regexp.MustCompile(`(?m)>\s*\[!warning\]\s*(.*)$`)
	text = warningRegex.ReplaceAllString(text, `{| class="wikitable" style="border-left:4px solid #f5ff56; background-color:#fdfaf8;"
| <div style="padding:0.5em;">
<strong style="color:#f5ff56;">⚠️ Warning:</strong> $1
</div>
|}`)

	// Success boxes
	successRegex := regexp.MustCompile(`(?m)>\s*\[!success\]\s*(.*)$`)
	text = successRegex.ReplaceAllString(text, `{| class="wikitable" style="border-left:4px solid #4e60e7; background-color:#f7f7fa;"
| <div style="padding:0.5em;">
<strong style="color:#4e60e7;">✓ Success:</strong> $1
</div>
|}`)

	// Note boxes
	noteRegex := regexp.MustCompile(`(?m)>\s*\[!note\]\s*(.*)$`)
	text = noteRegex.ReplaceAllString(text, `{| class="wikitable" style="border-left:4px solid #839df9; background-color:#f7f7fa;"
| <div style="padding:0.5em;">
<strong style="color:#071d49;">📝 Note:</strong> $1
</div>
|}`)

	return text
}

// ConvertCode converts code formatting
func ConvertCode(text string) string {
	// Code blocks first (before inline code to avoid conflicts)
	codeBlockRegex := regexp.MustCompile("(?s)```(\\w+)?\\n(.*?)```")
	text = codeBlockRegex.ReplaceAllStringFunc(text, func(match string) string {
		submatch := codeBlockRegex.FindStringSubmatch(match)
		lang := strings.TrimSpace(submatch[1])
		code := strings.TrimSpace(submatch[2])

		// Auto-detect language if not specified
		if lang == "" {
			codeStripped := strings.TrimSpace(code)
			if strings.HasPrefix(codeStripped, "{") || strings.HasPrefix(codeStripped, "[") {
				lang = "json"
			} else if strings.HasPrefix(codeStripped, "<") {
				lang = "xml"
			} else if strings.Contains(strings.ToUpper(code), "SELECT") || strings.Contains(strings.ToUpper(code), "FROM") {
				lang = "sql"
			} else {
				lang = "text"
			}
		}

		return fmt.Sprintf("<syntaxhighlight lang=\"%s\" line>\n%s\n</syntaxhighlight>", lang, code)
	})

	// Inline code: `code` -> <code style="background-color:#f5ff56;color:#021e57;">code</code>
	// Yellow background with Hero Blue text (Tietoevry branding)
	inlineCodeRegex := regexp.MustCompile("`([^`\n]+)`")
	text = inlineCodeRegex.ReplaceAllString(text, `<code style="background-color:#f5ff56;color:#021e57;padding:2px 6px;border-radius:3px;font-family:Consolas,Monaco,monospace;">$1</code>`)

	return text
}

// ConvertLists converts Markdown lists to MediaWiki format
func ConvertLists(text string) string {
	lines := strings.Split(text, "\n")
	result := make([]string, 0, len(lines))

	unorderedRegex := regexp.MustCompile(`^(\s*)[-\*]\s+(.*)$`)
	orderedRegex := regexp.MustCompile(`^(\s*)\d+\.\s+(.*)$`)

	for _, line := range lines {
		// Unordered lists: - item or * item -> * item
		if matches := unorderedRegex.FindStringSubmatch(line); matches != nil {
			indent := len(matches[1])
			content := matches[2]
			level := (indent / 2) + 1
			line = strings.Repeat("*", level) + " " + content
		} else if matches := orderedRegex.FindStringSubmatch(line); matches != nil {
			// Ordered lists: 1. item -> # item
			indent := len(matches[1])
			content := matches[2]
			level := (indent / 2) + 1
			line = strings.Repeat("#", level) + " " + content
		}

		result = append(result, line)
	}

	return strings.Join(result, "\n")
}

// AddHighlights adds highlighting markup for emphasized sections (Tietoevry branding for API endpoints)
func AddHighlights(text string) string {
	// Highlight API endpoints in code tags (e.g., Service/Method patterns)
	codeRegex := regexp.MustCompile(`<code>([^<>]+)</code>`)
	text = codeRegex.ReplaceAllStringFunc(text, func(match string) string {
		submatch := codeRegex.FindStringSubmatch(match)
		endpoint := submatch[1]

		// Check if it looks like an API endpoint (has a slash and CamelCase)
		if strings.Contains(endpoint, "/") && regexp.MustCompile(`[A-Z]`).MatchString(endpoint) {
			return fmt.Sprintf(`<mark style="background-color:#f5ff56"><code>%s</code></mark>`, endpoint)
		}
		return match
	})

	return text
}

// ConvertTables converts Markdown tables to MediaWiki format
func ConvertTables(text string) string {
	lines := strings.Split(text, "\n")
	result := make([]string, 0, len(lines))
	inTable := false

	pipeRegex := regexp.MustCompile(`^\|.*\|$`)
	separatorRegex := regexp.MustCompile(`^\|[\s\-:|]+\|$`)

	i := 0
	for i < len(lines) {
		line := strings.TrimSpace(lines[i])

		// Detect Markdown table (contains pipes)
		if strings.Contains(line, "|") && !inTable {
			if pipeRegex.MatchString(line) {
				// Start MediaWiki table
				result = append(result, `{| class="wikitable"`)
				inTable = true

				// Process header row
				cells := strings.Split(line, "|")
				cells = cells[1 : len(cells)-1] // Remove empty first and last
				result = append(result, "|-")
				for _, cell := range cells {
					result = append(result, "! "+strings.TrimSpace(cell))
				}

				// Skip separator line
				i++
				if i < len(lines) {
					nextLine := strings.TrimSpace(lines[i])
					if strings.Contains(nextLine, "|") && strings.Contains(nextLine, "-") {
						i++
					}
				}
				continue
			}
		} else if inTable && strings.Contains(line, "|") {
			// Check if separator line
			if separatorRegex.MatchString(line) && strings.Contains(line, "-") {
				i++
				continue
			}

			// Process table row
			cells := strings.Split(line, "|")
			if len(cells) > 2 {
				cells = cells[1 : len(cells)-1]

				// Skip if all cells are dashes
				allDashes := true
				for _, cell := range cells {
					if strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(cell, "-", ""), ":", "")) != "" {
						allDashes = false
						break
					}
				}
				if allDashes {
					i++
					continue
				}

				result = append(result, "|-")
				for _, cell := range cells {
					result = append(result, "| "+strings.TrimSpace(cell))
				}
			}
			i++
			continue
		} else if inTable && !strings.Contains(line, "|") {
			// End of table
			result = append(result, "|}")
			inTable = false
		}

		result = append(result, lines[i])
		i++
	}

	if inTable {
		result = append(result, "|}")
	}

	return strings.Join(result, "\n")
}

// ReverseChangelogOrder reverses the order of changelog version sections so newest appears first
func ReverseChangelogOrder(text string) string {
	// Find the changelog header
	changelogHeaderRegex := regexp.MustCompile(`(?s)(===<span[^>]*>.*?Changelog.*?</span>===)`)
	headerMatch := changelogHeaderRegex.FindStringIndex(text)

	if headerMatch == nil {
		return text
	}

	beforeChangelog := text[:headerMatch[0]]
	changelogHeader := text[headerMatch[0]:headerMatch[1]]

	// Find the next section (H1, H2, or H3) that ends the changelog
	// Matches =<span... or ==<span... or ===<span...
	nextSectionRegex := regexp.MustCompile(`(?m)^={1,3}<span[^>]*>[^<]*</span>={1,3}$`)
	remainingText := text[headerMatch[1]:]

	// Find all version header start indices
	versionHeaderRegex := regexp.MustCompile(`====<span[^>]*>Version[^<]*</span>====`)
	versionMatches := versionHeaderRegex.FindAllStringIndex(remainingText, -1)

	if len(versionMatches) == 0 {
		return text
	}

	// Find the end of the changelog section
	changelogEndIdx := len(remainingText)
	nextSection := nextSectionRegex.FindStringIndex(remainingText)
	if nextSection != nil {
		changelogEndIdx = nextSection[0]
	}

	// Extract version sections
	var versions []string
	for i, match := range versionMatches {
		start := match[0]
		var end int
		if i < len(versionMatches)-1 {
			end = versionMatches[i+1][0]
		} else {
			end = changelogEndIdx
		}

		// Sanity check to ensure we don't go out of bounds or have invalid ranges
		if start < end {
			versions = append(versions, remainingText[start:end])
		}
	}

	// Reverse the versions
	for i, j := 0, len(versions)-1; i < j; i, j = i+1, j-1 {
		versions[i], versions[j] = versions[j], versions[i]
	}

	// Reconstruct
	afterChangelog := remainingText[changelogEndIdx:]
	newChangelog := changelogHeader + "\n\n" + strings.Join(versions, "")

	return beforeChangelog + newChangelog + afterChangelog
}

// PrettifyCheckmarks replaces plain checkmarks with styled/prettier versions
func PrettifyCheckmarks(text string) string {
	// Replace all ✓ with green emoji checkmark ✅
	return strings.ReplaceAll(text, "✓", "✅")
}

// ConvertHorizontalRules converts Markdown horizontal rules to MediaWiki format
func ConvertHorizontalRules(text string) string {
	// Markdown horizontal rules: ---, ***, or ___
	// MediaWiki horizontal rules: ----

	// Match lines with 3 or more dashes, asterisks, or underscores
	hrRegex := regexp.MustCompile(`(?m)^[\s]*[-*_]{3,}[\s]*$`)
	text = hrRegex.ReplaceAllString(text, "----")

	return text
}

// Convert performs the main conversion with optional concurrent processing
func Convert(markdownText string, config Config) string {
	text := markdownText

	// Add CSS styling header if requested
	styleHeader := ""
	if config.AddStyling {
		styleHeader = GetCodeStylingCSS() + "\n\n"
	}

	// Sequential processing (Concurrent mode disabled for stability)
	// Process code blocks FIRST to protect underscores and other special characters
	text = ConvertCode(text)
	text = ConvertBoldItalic(text)
	text = ConvertHeaders(text)
	text = ConvertLinks(text)
	text = ConvertCallouts(text)
	text = ConvertLists(text)
	text = ConvertTables(text)
	text = ConvertHorizontalRules(text)
	text = AddHighlights(text)

	// Post-processing improvements
	text = ReverseChangelogOrder(text)
	text = PrettifyCheckmarks(text)

	return styleHeader + text
}
