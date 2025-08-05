# TRMNL 2-Bit Demo

A simple web server that serves random 2-bit PNG images for TRMNL displays running firmware 1.6.0+.

## Usage

Returns a random image URL with configurable refresh rate:
- GET `/` - Returns JSON with image URL and refresh rate
- GET `/images/*` - Serves the actual image files

Configure with TRMNL's [Redirect Plugin](https://usetrmnl.com/integrations/redirect).

Web Address: https://2bitdemo.scottlabs.io

Optional query parameter: `?refresh=600` (default: 300 seconds)

## Contributing

Pull requests with new images are welcome! Add PNG files to the `images/` directory.

## Image Attribution

- `images/legendary-palm-tree/` - From https://github.com/GLdashboard
