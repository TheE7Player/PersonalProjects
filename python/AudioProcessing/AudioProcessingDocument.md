# Audio Processing Project

> Date created: 4th March 2021

Who does not like listening to their music, after a long day of work or finding that song that is stuck in your head - but you have a scenario...

## Problem

Some music I have tends to either play another song (or random shuffle song) in the same track before it ends - This can be annoying as it suddenly cuts - which makes the experience of listening annoying.

## Solution

Using a program or script available, go through each song and find a way to identify if a song has ended correctly.

## Goals/Requirements:

- Must be simplistic, either through a simple library or API alternative which can deliver the problem to be resolved

- Must be short and not too bloated to achieve the goal (Ideally below 100-300 lines)

- Must make usage of CPU - meaning it must support multi-threading (or at least a thread if viable)

- Scan through a lot of songs and identify them in less than 10 seconds if possible

### Algorithm/Steps

For the project, I do not want to rely on a specific algorithm if possible. I do not require any need of advanced algorithm to identify if a song has ended as it should do - I simply just need a way to identify silence or look at the `rms` or `dbs` to determine it.

---

## Looking at the samples at fault

I sadly need to start to look through the songs to identify *at least* one or two of them to see how I can approach towards a solution. Even if I can find one and attempt to run the script or program to find others will save time and effort.

Let us look at 3 samples...

### Example A

![enter image description here](https://github.com/TheE7Player/PersonalProjects/blob/main/python/AudioProcessing/resource/ExampleA.png?raw=true)

> Green: Original Song
> 
> White: Area of the volume when faded
> 
> Red: Area of Silence

From this visual, we can *easily* identify when the song truly ends. It does not take a genius to find that one out. We also know if the song has finished if there is silence then finishes or if the song just simply stops at the end.

From this we have a good idea of how solve it - We just look at the last seconds of the song and determine if its loud or not.

## Example B

![enter image description here](https://github.com/TheE7Player/PersonalProjects/blob/main/python/AudioProcessing/resource/ExampleB.png?raw=true)

> White: max and min amplitude of the original song
> 
> Red: Fade in and out of the current song to another

When I came across this sample, I knew the first solution might not work as well as it did. This showed me that there is a chance a song starts with a fade-in effect... We can somewhat still identify, but it will be much harder to do so.

## Example C

<img title="" src="https://github.com/TheE7Player/PersonalProjects/blob/main/python/AudioProcessing/resource/ExampleC.png?raw=true" alt="enter image description here" data-align="inline">

> White: max and min amplitude of the original song
> 
> Red: Fade in and out of the current song to another

Yet again, we found another sample with this annoying fade in and out effect - We might as well just look at the end after all...

---

## Pseudocode Solution

The idea is simple, grab each song. Store it and flip it around, determine the dbs or amount of data/stream by an array.

> Assume threshold is set at -60dbs

```
SET THRESHOLD TO -60
SET songs TO LIST OF STRING

REPEAT
    SET audio TO GET_NEXT_TRACK()
    SET end TO audio.GET_LAST_SECOND()

    IF end.GET_AVG_DB() > THRESHOLD THEN
        songs.Add(audio)
    END IF   
UNTIL END OF LIST
```

So, in English we:

- Set the threshold (Where its mostly silent)

- Create a list which contains strings of the file’s location

- We loop through each song
  
  - We grab the information of the track
  
  - We then get the last few seconds of the track
  
  - Validate if the average dbs from the end goes past the threshold
  
  - If so, add it to the list and continue onwards

### Attempting to avoid `O(n)` if possible (front-end)

The goal is to make it simple and effective. If we have a linear loop per song, the processing time and power can be the difference between quickness or accuracy. I would like it to be quick, as I want the script or program to identify to songs for me to manually go through...

We cannot *ever* attempt to avoid `O(n)` as we depend on it for enumerating over arrays or lists! We do require it for other things, but I personally do not want to use it to process the audio whenever possible.

---

## The million-dollar question: What language?

Choosing a language for a project is a huge thing to consider... Do we favour: Readability, Efficiency, Robustness, Type Safety?

Do we attempt the project using a compiler which does some neat optimisations?  Or go for a language which uses an interpreter and make it costly in performance side of things? Believe it or not - its `Python`.

### Why Python? Your scared of another language?

`Yes` but `no`, but `yes` but also `no`. Well...

I found this library `pydub` which looks simple and does everything I need it to possibly do for me.

Loading the library up is simple, and with an `virtual environment` - I can install libraries I need to make it work like a piece of cake!

With the way `Python` always works us to have less code in our script compared to other languages which requires extra functionality and setups.

---

## Code breakdown

The main script in `main.py` contains the code to identify the songs, as the `validate.py` is used to go through each file it found in `results.txt` (created after `main.py` is finished)

> Ensure the venv is created before running the code, you can find tutorials about it in YouTube or google it

We first import the library `pydub` which allows us to read or write the audio files in no time!

```python
# Audio Processing Library (pydub) to process the audio
from pydub import AudioSegment
```

For `validate.py`, we require the `simpleaudio` library to play back the songs

```python
from pydub.playback import play
```

## The issue

If we had `.wav` files we are fine - but I must look through `.mp3` files. `pydub` suggests using `ffmpeg` in order to do so - so please install `ffmpeg` for your system to make sure the process works!

After we have `ffmpeg` configured, it is time to get to business!

We have 2 steps:

1. Setup the variables and get the files from the folder

2. Split up the songs into chunks to allow `multi-threading` validation

### Step 1: Setup

We need 3 variables at most for this part:

- `path` - Holds the location to what to look for ( using `raw-string` since the string contains `\`)

- `songs` - List which holds all the songs to validate

- `target` - A string which holds the find condition for `glob` library (`*.mp3`)

```python
path = r"G:\Users\TheE7Player\Music\Music\Spotify"
songs = []
target = path + "\*.mp3"
```

Then we simply perform a `O(n)` operation to fill in the songs into the `songs` list...

```python
for song in glob.glob(target):
    songs.append(song)
```

So, we use the `glob` library to do the searching for us, and simply enumerate over the findings into a list - as easy as that!

### Step 2: Chunk Splitting

For us to make `multi-threading` worth file we need to split the single array into multiple chunks to allow each thread to do 100 independent validations per thread.

Before that - again - we need to setup some variables for cache reasons:

- `revise_songs` - A list which holds the songs that was found to have more than the threshold is set to (possible found)

- `failed_songs` - A list which holds songs that did not manage to process, kinda pointless as we do not do anything with it

- `temp` - A cache variable which holds the song meta data

- `tempVal` - A cache variable alongside `temp` which holds the last few seconds of the same track

- `tailLength` - Length in milliseconds (ms) of how far to trim the song (set to `1000`, meaning `1 second`)

- `songCount` - A counter to count how many songs got processed

- `songMax` - A cached variable which holds the max number of songs, better than just calling `len()` all the time (as the list does not expand at all)

- `splitSize` - A variable which holds how many chunks (in 100) is possible, then rounded to an `int` for a nice even number

- `chunk` - A list of lists which holds a list of 100 songs

- `adx` - Chunk index pointer. Keeps track of which chunk the list is being appended to

- `idx` - Song pointer. Keeps track of the number of songs that has been added to the current chunk pointer

Then again, we do a `O(n)` operation to create the chunks. Not complicated logic here, but hard if you do not know how to lay it out...

```python
# Ignoring other variables at this point - focus on the algorithm

chunk = []
adx = 0
idx = -1

for s in songs:
    if idx == -1:
        chunk.append([])
        idx = 0

    chunk[adx].append(s)
    idx += 1

    if idx == 100:
        idx = -1
        adx += 1
```

## The fun part - The Processing Logic

Here we make good usage of our `cached variables` to prevent the memory for being overload - a casual mistake most programmers tend to make.

We always make usage of `temp` and `tempval` to prevent the `stack` increasing, it is also faster as these variables are already stored!

So, we have a function called `Process` which takes in an argument of `list` called `arr` - how convenient - to allow each thread to have access to the same logic.

Since the variables we want to access is out-with the scope of the script, we need to make usage of `global` in `Python`. This is not an issue with most languages but is with `Python`.

```python
def Process(arr):
    global songCount, revise_songs, failed_songs
    ...
```

> Not sure if there could be a case of a race condition to utilize the variables, you make look into this if you wish. Threading is fun, well kinda.

The iteration loop is simplistic in its understanding:

```python
for song in arr:

        songCount += 1

        # See if its possible to process the file
        try:
            temp = AudioSegment.from_mp3(song)
        except:
            failed_songs.append(song)
            continue

        # This means it has processed the file into a variable called 'temp'
        # the 'temp' file is immutable! It cannot be changed nor modified
        tempVal = temp[-tailLength:]    

        if tempVal.dBFS > -40:
            revise_songs.append(song)
```

For each iteration:

- We increment `songCount` by one

- We call `AudioSegment.from_mp3` function from `pydub` and pass in the song in a `try catch` block.
  
  - We continue if the function continued with issues
  
  - We add the song to the failed list if that operation failed

- Then we trim the song to the very last seconds by storing it into `tempVal`

- Now we validate if the average `dBFS` from the trimmed song `tempVal` is beyond our threshold - in this case `-40db`. If so, we add it onto the list.

## Hold your horses - Threading time!

Now we going to fulfil our goal to make the project utilise threading, making most of the CPU performance and improve speed of the validation.

We import the threading library first in `Python` before we do any operation:

```python
# Put this at the top or near where the threading is being called from
import threading

# To allow the thread to sleep for a few seconds
from time import sleep
```

Then we have a list of threads to keep track of...

```python
threads = []
```

Then we use the cached variable to let the program create the right number of threads to create:

```python
for i in range(0, splitSize):
    thr = threading.Thread(name=f"ProcessThread{i}", target=Process, args=[chunk[i]], daemon=True)
    threads.append(thr)
```

> Since we start at 0, it will say (in my case) 4 threads instead of 5. Remember n - 1 for arrays.

For each thread:

- We give each thread the same prefix (ProcessThread) but different suffix depending on its count (ProcessThread0, ProcessThread1 etc)

- `target` - an argument which tells the thread which function it should run; we want it to run the processing function called `Process` we made either on.

- `args` - another argument that is important as the `Process` function requires a list argument to function.  So, we pass in the correct chunk to each function.

- `daemon` - telling the thread that this is a background operation, it will close itself once its complete, no waiting around for nothing! (Non-blocking thread basically)

After setting up the threads, it is time to enumerate again (`O(n)`) and start the process...

```python
for thread in threads:
    print("Thread starting called", thread.name)
    thread.start()
```

Now in the main thread (script) we simply just loop forever until all threads are finished... we utilise `sleep` to prevent the main loop for being too CPU dependent

```python
while True:
    if len(threads) == 0:
        break
    for thr in threads:       
        if not thr.is_alive():
            threads.remove(thr)
    sleep(1)
    print(f"Processing... {round((songCount / songMax) * 100, 0)}% complete with {len(revise_songs)} songs and {len(failed_songs)} failed processing\t\t", end="\r", flush=True)
```

Our break condition is:

- If the `threads` array does not have any elements in it anymore ( `size` == `0` )

Then we loop through each thread checking if its available by calling its function `.is_alive()`. If it is not alive then we simply remove it from the list.

We then wait 1 second `sleep(1)`, update the text in the console and do the same operation all over again.

> printl utilises the argument `flush` which rewrites the text to the same screen. Using `\t\t` prevents a string that is small to show the last displayed text. Using `\r` for the `end` argument simply just puts the cursor pointer back to the start of the string

Once we done that, we show a message and create a `result.txt` file to look through of the songs that it found was loud near the end of the song.

```python
print("finished with", len(revise_songs))

with open("result.txt", "w") as output:
    output.write("\n".join(revise_songs))
```

In my first completed attempt, it showed that 62 songs were detected - fantastic! It is now time to go over each one manually to see if it found what I wanted it to.

I did have `validate.py` do this for me, but I wanted it to be put through to a program called `Ocenaudio` which shows me the audio wave and *allows* me to re-save the file itself without asking to save it to a new name. You can do this through `pydub` but you have the risk of the song being chopped off.

---

## Validating and Testing

The fun, the boring and the annoying testing!

Now I must go through all the 62 songs to see if it is possible.

The problem is that some songs fade out typically a second near the end - what I was testing for - which makes the process more annoying, but it helped me to target songs faster and not going through or keeping track all the time.

I made a spreadsheet of all the songs and marked if the threshold was accurate to identify the issue - I could *spend* more time perfecting the threshold but not all the songs are mixed to the same level or end in the same manner.

Here are the results of 62 songs:

| Subject   | Amount | Percentage |
| --------- | ------ | ---------- |
| Correct   | 21     | 33.8%      |
| Wrong     | 41     | 66.1%      |
| Precision | 42     | 42%        |

*Not* the best way nor the worst. I did not want to spend all day to find the correct way to do this as the goal was simple and not complex. It is not close to 50 50 in finding the issue but is close enough to be efficient-ish.

But hey, at least I got the solution for it.
