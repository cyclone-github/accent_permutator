# accent_permutator
- accent_permutator is a tool to transform characters to accented characters such as "o" to "ò"
- tool will accept ASCII, UTF-8 and $HEX[] formatted wordlists
- ex input: plaintext "password" or "$HEX[70617373776f7264]"
- ex usage: cat wordlist.txt | ./accent_permutator.bin | hashcat.bin...
- written by cyclone

### compile hashgen from source
- If you want the latest hashgen features, compiling from source is the best option since the release version may run several revisions behind the source code.
- Compile from source code info:
- https://github.com/cyclone-github/scripts/blob/main/intro_to_go.txt

### version history
- v2023-08-10.2330; initial github release
