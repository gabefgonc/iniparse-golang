# INIParse GoLang

**An INI File format parser written in golang, including a library and a CLI app**

## Building

`go build cmd/iniparse.go`

## Usage

`./iniparse FILENAME [QUERY]`

**QUERY** must be the name of the section you want to query, followed or not by a period and the name of the item to be queried.

**Example**:

`./iniparse example.ini apple`
`./iniparse example.ini pear.color`
