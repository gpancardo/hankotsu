# Hankotsu 🦫🔪📃
## What is Hankotsu?
**__Hankotsu__ are Japanese boning knifes with curve blades for meat and joint separation.** This tool strives to separate relevant entries in very large datasets based on substrings on a particular column.\
Dealing with datasets bigger than memory with limited resources is hard and time consuming, Hankotsu provides an alternative for this: **extracting relevant entries to a new file without doing a full copy to memory.**
## Why Go?
I am developing this tool to build the dataset I will use for another project, so time is key.\
Go has a great development experience and I am also way more familiar with it than similar alternatives. It is a very efficient language that will have a (hopefully) quick and fun development.
## Tutorial
### Installation
#### Download a ready-made executable (no Go needed)
Go to the [latest release](https://github.com/gpancardo/hankotsu/releases/latest) and download the file for your system:

| System | File |
| --- | --- |
| Windows (most PCs) | `hankotsu-windows-amd64.exe` |
| Windows on ARM | `hankotsu-windows-arm64.exe` |
| macOS, Apple Silicon (M1–M4) | `hankotsu-darwin-arm64` |
| macOS, Intel | `hankotsu-darwin-amd64` |
| Linux (most PCs and servers) | `hankotsu-linux-amd64` |
| Linux on ARM (Raspberry Pi 4/5, ARM servers) | `hankotsu-linux-arm64` |

**Linux / macOS**
```
chmod +x hankotsu-linux-amd64              # use the name of the file you downloaded
sudo mv hankotsu-linux-amd64 /usr/local/bin/hankotsu
hankotsu --version
```
On macOS, if the system says the file is from an unidentified developer, run
`xattr -d com.apple.quarantine /usr/local/bin/hankotsu` once.

**Windows**
Rename the file to `hankotsu.exe`, put it in a folder such as `C:\Tools`, and run it from
PowerShell or the Command Prompt in that folder (or add the folder to your `PATH`).

**Check the download (optional).** Each release includes `SHA256SUMS.txt`:
`sha256sum -c SHA256SUMS.txt --ignore-missing` on Linux, `shasum -a 256 FILE` on macOS,
`Get-FileHash FILE` on Windows.

#### Build it yourself (if you already use Go)
```
git clone https://github.com/gpancardo/hankotsu
cd hankotsu
go build -o hankotsu .
sudo mv hankotsu /usr/local/bin/
```

### Usage
1. Install Hankotsu.
2. Create a .json file with the column to search and the substrings to look for:
```
{
    "label":"THE_COLUMN_YOU_ARE_LOOKING_FOR",
    "words":["A", "LIST", "OF", "THE", "KEYWORDS", "YOU", "ARE", "LOOKING", "FOR"]
}
```
3. Call ```hankotsu  PATH_OF_YOUR_ORIGINAL_CSV_FILE  PATH_OF_YOUR_JSON_FILE```

Hankotsu writes the header plus every row whose column contains at least one of the words to
`READY_<your file>.csv`, in the same folder as the original. It reads the file line by line,
so memory stays flat no matter how big the file is.

#### Example
`people.csv`:
```
id,city,name
1,Mexico City,Ana
2,Monterrey,Luis
3,Mexico City,Sofía
```
`keywords.json`:
```
{ "label": "city", "words": ["Mexico"] }
```
`hankotsu people.csv keywords.json` creates `READY_people.csv`:
```
id,city,name
1,Mexico City,Ana
3,Mexico City,Sofía
```

### Good to know
- Matching is a case-sensitive substring match: `"Mexico"` matches `Mexico City` but not `MEXICO`.
- The file must be comma-separated. Fields that contain commas inside quotes are not supported yet.
- If the column is not in the header, Hankotsu stops and tells you instead of writing an empty file.
