# Audio Processing Project

> Date created: 4th March 2021

Who doesn't like listening to their music, after a long day of work or finding that song that is stuck in your head - but you have a scenario...

## Problem

Some music I have tends to either play another song (or random shuffle song) in the same track before it ends - This can be annoying as it suddendly cuts - which makes the experience of listening annoying.

## Solution

Using a program or script available, go through each song and find a way to identify if a song has ended correctly.

## Goals/Requirements:

- Must be simplistic, either through a simple library or API alternative which can develier the problem to be resolved

- Must be short and not too bloated to achieve the goal (Ideally below 100-300 lines)

- Must make usage of CPU - meaning it must support multi-threading (or at least a thread if viable)

- Scan through a lot of songs and identify them in less then 10 seconds if possible

### Algorthim/Steps

For the project, I don't want to relie on a specific algorthim if possible. I don't require any need of advanced algorthim to identify if a song has ended as it should do - I simply just need a way to identify silence or look at the `rms` or `dbs` to dermin it.

---

## Looking at the samples at fault

I sadly need to start to look through the songs to identify *at least* one or two of them to see how I can apporach towards a solution. Even if I can find one and attempt to run the script or program to find others will save time and effort.

Let's look at 3 samples...

### Example A

![enter image description here](https://github.com/TheE7Player/PersonalProjects/blob/main/python/AudioProcessing/resource/ExampleA.png?raw=true)

> Green: Orginal Song
> 
> White: Area of the volume when faded
> 
> Red: Area of Silence



From this visual, we can *easily* identify when the song truly ends. It doesn't take a genuis to find that one out. We also know if the song has finished if their is silence then finishes or if the song just simply stops at the end.

From this we have a good idea of how solve it - We just look at the last seconds of the song and determine if its loud or not.



## Example B

![enter image description here](https://github.com/TheE7Player/PersonalProjects/blob/main/python/AudioProcessing/resource/ExampleB.png?raw=true)

> White: max and min amplitude of the orginal song
> 
> Red: Fade in and out of the current song to another



When I came a cross this sample, I knew the first solution might not work as well as it did. This showed me that there is a chance a song starts with a fade-in effect... We can somewhat still identify, but it will be much harder to do so.



## Example C

<img title="" src="https://github.com/TheE7Player/PersonalProjects/blob/main/python/AudioProcessing/resource/ExampleC.png?raw=true" alt="enter image description here" data-align="inline">

> White: max and min amplitude of the orginal song
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

So in english we:

- Set the threshold (Where its mostly silent)

- Create a list which contains strings of the files location

- We loop through each song
  
  - We grab the information of the track
  
  - We then get the last few seconds of the track
  
  - Validate if the average dbs from the end goes past the threshold
  
  - If so, add it to the list and continue onwards




