package main

const (
	author = "@dwisiswant0"
	banner = `
  unch 😗 %s
  --
  Hides message w/ invisible Unicode chars
  by ` + author

	usage = `
  unch [-dln] "MESSAGE" "PLAINTEXT"`
	options = `
  -d, -decode       Decodes the encoded MESSAGE
  -l, -lorem        Generate random PLAINTEXT with lorem ipsum
  -n                Do not output the trailing newline
  -h, -help         Print this helps text`
	examples = `
  unch "s3cr3t" "Hello, world!"
  unch -lorem "s3cr3t"
  unch "${SECRET_MESSAGE}" "${MASKED_MESSAGE}"
  unch -d "${ENCODED_MESSAGE}"
  unch -d < "ENCODED.txt`
)
