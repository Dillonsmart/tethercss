# TetherCSS
## Keeping things simple with CSS

When did CSS get so complicated? TetherCSS is a simple, lightweight CSS framework that helps you build responsive, mobile-first websites quickly and easily. It provides a set of basic styles and components that you can use as a foundation for your projects.

## Installation
It's so easy to get started with TetherCSS. Grab a copy of `tether.css`, then include the CSS file in your HTML:

```html
<link rel="stylesheet" href="path/to/tether.css">
```

## Usage
If you wish to customize TetherCSS, you can either modify the css classes directly, or generate a new `tether.css` file from the provided Go script.

In order to run the script, you will require go version 1.25 or higher.

Once you've made the changes you want, you can generate a new `tether.css` file by executing the build script:

```bash
go run src/generate.go
```