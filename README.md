# CCWC

Command line tool to count number of lines, bytes, words, characters

## SYNOPSIS

CCWC [-clmw] [file ...]

## DESCRIPTION

The wc utility displays the number of lines, words, and bytes contained in each input file, or standard input (if no file is specified) to the standard output.
A line is defined as a string of characters delimited by a ⟨newline⟩

The ccwc command/utility tool is to count number of lines, words, bytes, characters in each input file, or stdin (if no file is provided as arguments).
character. Characters beyond the final ⟨newline⟩ character will not be included in the line count.

A line is defined as a string of characters delimited by a ⟨newline⟩.
A word is defined as a string of characters seperated by white space characters.
A line of cumulative counts for all the files is displayed on a separate line after the output for the last file.

The following options are available:

-c The number of bytes in each input file is written to the standard output.

-l The number of lines in each input file is written to the standard output.

-m The number of characters in each input file is written to the standard output.

-w The number of words in each input file is written to the standard output.

When an option is specified, wc only reports the information requested by that option.

If no files are specified, the standard input is used and no file name is displayed. The prompt will accept input until receiving EOF, or [^D] in most
environments.

If wc receives a SIGINFO (see the status argument for stty(1)) signal, the interim data will be written to the standard error output in the same format as the
standard completion message.
