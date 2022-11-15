# Buzzy is a NYT Spelling Bee Solver

The NYT Spelling Bee is a fantastic puzzle, but sometimes you just want to win.

To that end I wrote a "solver" for it as an exercise.  This works in two phases:

1. You load a corpus of words, which are used to build an on-disk lookup system.
2. You run the solver to get all possible words for a given puzzle.

Pre-computing the dictionary lets the solver run faster and is, imo, tidier.

```
./solver . f n l o d i g
Pivoting on: f
With letters: n l o d i g
Loading from patterns...
Found 89 words:

  diff
  diffiding
  doff
  doffing
  dolf
  fidding
  ...
```

# How it works

## Precomputing

The bulk of the work is done in the dictionary processor.  For each word in the file, it will create a 4 byte value which is a bitfield.  Each bit in the field is the presence of a particular letter.  Every possible unique combination of letters can be encoded in 4 bytes (or, 26 bits really)

Those four bytes are then base32 encoded (not base64 because MacOS APFS is not case-sensitive by default), and the word is appended to a file matching that value.

## Solving

At lookup time, all the possible unique combinations of the letters in the puzzle are found, the bitfield for each is computed, and then each file is read to get the potential words out for display.

# FAQ

## What is "Spelling Bee"

[Spelling Bee](https://www.nytimes.com/puzzles/spelling-bee) is a New York Times puzzle game, where you are provided with 7 letters.

One of the letters is the "center letter", or what I call the "pivot letter".

The goal is to make as many words as possible using only these letters, always including the center/pivot letter.  Words must be four letters or longer.

## What dictionary should I use?

This is a tough one.  The NYT does not publish a wordlist/dictionary so far as I can find.  As such, you'll probably want to use the largest english wordlist you can, and cope with the false positives.  The NYT excludes words for many reasons, to quote:

> Each Spelling Bee puzzle is curated to focus on relatively common words (with a few tougher ones periodically to keep things challenging). We try to avoid terms that are ultra-specific to any professional field to maintain a level playing field for all of our solvers. 
[- New York Times Games](https://help.nytimes.com/hc/en-us/articles/360029050872-Word-Games-and-Logic-Puzzles)

Personally, I used [dwyl/english-words](https://github.com/dwyl/english-words) when developing this software.  But it is extensive and leads to a _lot_ of false positives.

