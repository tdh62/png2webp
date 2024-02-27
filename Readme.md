# PNG2WEBP

This is a simple script to convert PNG,JPG,GIF files in a directory to WEBP format.

## Usage

```bash
./png2webp <filename|directory> [Quality] [options]
```

## Options

- Quality : Quality of the image. Default is 80.
- -l, -lossless: Convert the image to lossless webp format.
- -d, -dir_walker: Walk through the directory and convert all the images in the directory.
- -o, -overwrite: Overwrite the original image with the webp image.

## Example

Switch all images(png,jpg,gif format) files in the current directory to webp format with 75% quality, overwrite if the webp file already exists:

```bash
./png2webp . 75 -d -o
```

Convert test.png to lossless webp test.webp format:

```bash
./png2webp test.png -l
```

