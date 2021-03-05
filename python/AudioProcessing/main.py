# Python 3.8.5

# Audio Processing Library (pydub) to process the audio
from pydub import AudioSegment
# from pydub.playback import play

# glob library used to look through directories and files
import glob

import threading
from time import sleep

path = r"G:\Users\TheE7Player\Music\Music\Spotify"
songs = []
target = path + "\*.mp3"

print("Adding songs from path:", path)
for song in glob.glob(target):
    songs.append(song)

print("Finished adding songs from path:", path)   

print("Going through", len(songs), "songs...")

revise_songs = []
failed_songs = []

temp = None
tempVal = None
# How far we are doing to determine the validation (1 seconds = 1000ms)

tailLength = 1000 
songCount = 0
songMax = len(songs)

splitSize = int(len(songs) / 100)

chunk = []
adx = 0
idx = -1
for s in songs:
    if idx == -1:
        chunk.append([])
        idx = 0

    chunk[adx].append(s)
    idx+=1

    if idx == 100:
        idx = -1
        adx += 1

def Process(arr):
    global songCount, revise_songs, failed_songs
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

threads = []
print("All threads are now running... Please wait till all complete!")
for i in range(0, splitSize):
    thr = threading.Thread(name=f"ProcessThread{i}", target=Process, args=[chunk[i]], daemon=True)
    threads.append(thr)

for thread in threads:
    print("Thread starting called", thread.name)
    thread.start()
    
while True:
    if len(threads) == 0:
        break
    for thr in threads:       
        if not thr.is_alive():
            threads.remove(thr)
    sleep(1)
    print(f"Processing... {round((songCount / songMax) * 100, 0)}% complete with {len(revise_songs)} songs and {len(failed_songs)} failed processing\t\t", end="\r", flush=True)

# To Fix (?) : Percentage doesn't finish at 100%?

print("finished with", len(revise_songs))

with open("result.txt", "w") as output:
    output.write("\n".join(revise_songs))