# Hankotsu 🦫🔪📃
## What is Hankotsu?
**__Hankotsu__ are Japanese boning knifes with curve blades for meat and joint separation.** This tool strives to separate relevant entries in very large datasets based on substrings on a particular column.\
Dealing with datasets bigger than memory with limited resources is hard and time consuming, Hankotsu provides an alternative for this: **extracting relevant entries to a new file without doing a full copy to memory.**
## Why Go?
I am developing this tool to build the dataset I will use for another project, so time is key.\
Go has a great development experience and I am also way more familiar with it than similar alternatives. It is a very efficient language that will have a (hopefully) quick and fun development.
## Tutorial
### Installation
#### If you already use Go:
1. Download this project and extract all files on a folder.
2. Go to that folder's path through a terminal.
3. Write ```"go build -o hankotsu main.go"```
4. Write ```sudo mv hankotsu /usr/local/bin/```+ Enter
5. Write ```sudo chmod +x /usr/local/bin/hankotsu```+ Enter
6. Now you can call the __hankotsu__ command anywhere in your system!
#### If you do not use Go:
Pre-build version comming soon!
### Usage
1. Install Hankotsu
2. Create a .json file with this format:
```
{
    "label":"THE_COLUMN_YOU_ARE_LOOKING_FOR",
    "words":["A", "LIST", "OF", "THE", "KEYWORDS", "YOU", "ARE", "LOOKING", "FOR"]
}
```
3. Call __hankotsu__ __path of your original CSV file__ __path of the JSON with the column label and keywords__