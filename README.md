# README

mdtree - Parse tree command output to markdown list syntax.

## Usage

Install binary in your $PATH from [releases](https://github.com/fivestar/mdtree/releases).

Use instead `tree`:

```
% mdtree -L 2 -h
* .
    * [ 454]  composer.json
    * [1.1K]  LICENSE
    * [4.0K]  Psr
        * [4.0K]  Log
    * [1.1K]  README.md

* 2 directories, 3 files
```

## Requirements

* `tree`

