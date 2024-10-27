package command

import (
	ccwcerror "ccwc/error"
	"strings"
)

type (
	CommandParser interface {
		Parse([]string) (*Command, error)
	}

	commandParserImpl struct {
		validOptions []byte
	}

	Command struct {
		Options         string
		InputFilesPaths []string
	}
)

func NewCommandParser() CommandParser {
	return &commandParserImpl{
		validOptions: []uint8{'l', 'm', 'c', 'w'},
	}
}

func (c *commandParserImpl) Parse(cmdArgs []string) (*Command, error) {
	argsBuilder := strings.Builder{}
	files := make([]string, 0)
	for _, arg := range cmdArgs {
		if arg[0] == '-' {
			if len(arg) < 2 {
				return nil, ccwcerror.WRONG_OPTIONS
			}
			argsBuilder.WriteString(arg[1:])
		} else {
			files = append(files, arg[0:])
		}
	}
	options := argsBuilder.String()
	if options != "" && c.validOptions != nil && len(c.validOptions) > 0 {
		validOpts := make(map[uint8]int)
		for _, vo := range c.validOptions {
			validOpts[vo]++
		}
		for _, o := range []byte(options) {
			_, exist := validOpts[o]
			if !exist {
				return nil, ccwcerror.WRONG_OPTIONS
			}
		}
	}
	return &Command{
		Options:         options,
		InputFilesPaths: files,
	}, nil
}
