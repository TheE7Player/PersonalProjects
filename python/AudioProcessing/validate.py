# Python 3.8.5

# Audio Processing Library (pydub) to process the audio
from pydub import AudioSegment
from pydub.playback import play

# Load in the result.txt file and go through each one

input_value = None
file = 'result.txt'
content = None

# Load the text file into an array
with open(file) as f:
    content = f.readlines()

# Iterate through these array and get rid of the '\n' at the end of each string (using .rstrip)
for i in range(0, len(content)):
    content[i] = content[i].rstrip()

# Now lets go through each one
left_over = []

song_cur = 0
song_max = len(content)
song_length = 3000 # 3s (3000ms)

temp = None
tempVal = None

for song in content:
    song_cur += 1
    print(f"playing song {song_cur} of {song_max} for {song_length / 1000} seconds : Type 'y' if it plays or changes onwards, 'n' if not")
    temp = AudioSegment.from_mp3(song)
    tempVal = temp[-song_length:]
    play(tempVal)
    input_value = input("Did the song change volume or other song started playing? y / n")

    if input_value == "y":
        left_over.append(song)   

print("done")