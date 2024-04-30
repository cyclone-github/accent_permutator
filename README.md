[![Readme Card](https://github-readme-stats.vercel.app/api/pin/?username=cyclone-github&repo=accent_permutator&theme=gruvbox)](https://github.com/cyclone-github/)

# accent_permutator
- accent_permutator is a tool to transform characters to accented characters such as "o" to "ò"
- tool will accept ASCII, UTF-8 and $HEX[] formatted wordlists
- ex input: plaintext "password" or "$HEX[70617373776f7264]"
- ex usage: cat wordlist.txt | ./accent_permutator.bin | hashcat.bin...
- written by cyclone

### version history
- v2023-08-10.2330; initial github release
- v2023-08-16.1430; changed read logic to process large wordlists line by line, preventing memory issues

### Compile from source:
- If you want the latest features, compiling from source is the best option since the release version may run several revisions behind the source code.
- This assumes you have Go and Git installed
  - `git clone https://github.com/cyclone-github/accent_permutator.git`
  - `cd accent_permutator`
  - `go mod init accent_permutator`
  - `go mod tidy`
  - `go build -ldflags="-s -w" .`
- Compile from source code how-to:
  - https://github.com/cyclone-github/scripts/blob/main/intro_to_go.txt
