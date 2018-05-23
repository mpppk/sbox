# sbox
sbox is command line tool for [Scrapbox](https://scrapbox.io)

## Installation

### From source

```
$ go get github.com/mpppk/sbox
```

## Setup

```
$ echo "project: replace_me_by_your_project_name" >> ~/.config/sbox/.sbox.yaml
```

## Commands

### sbox browse [page title]
`sbox browse` open specified title page by default browser.

example1: Open or Create page whose title is "awesome page" by default browser

```Shell
$ sbox browse "awesome page"
``` 

example2: Open or Create page whose title is "awesome page" and add the line "awesome contents"

```Shell
$ sbox browse "awesome page" --contents "awesome contents"
``` 

example3: Open or Create page whose title is "Syntax" of the "help" project

```
$ sbox browse --project help Syntax
```

example4: Create today page

```
$ sbox browse `date "+%Y/%m/%d"` --contents `date "+[%Y/%m]/%d"`
```

### sbox list pages
`sbox list pages` list page titles

```
$ sbox list pages
awesome page title
super awesome page title
hyper awesome page title
```

### sbox show [page title]
`sbox show` show page contents

```
$ sbox show "awesome page"
awesome page
awesome content
awesome contents2
```

## Tips

### sbox with fuzzy finder

example: Choose page by fuzzy finder and open the one.

```
$ sbox list pages | fzf | xargs -0 sbox browse
```

[fzf](https://github.com/junegunn/fzf) is used in the above example, but you can choose your favorite fuzzy finder like [peco](https://github.com/peco/peco), [percel](https://github.com/parcel-bundler/parcel), and more!
